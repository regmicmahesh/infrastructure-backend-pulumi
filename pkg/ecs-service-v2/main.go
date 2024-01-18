package ecsservicev2

import (
	"github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/label"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type EcsService struct {
	pulumi.ResourceState
	ServiceConnectEndpoints []pulumi.StringOutput
	TaskExecutionRole       iam.RoleOutput
	TaskDefinition          ecs.TaskDefinitionOutput
}

type PortMappingLoadBalancerConfig struct {
	Enabled        bool
	TargetGroupArn pulumi.StringInput
}

type SSMParameterConfig struct {
	Names pulumi.StringArray
}

type PortMappingConfig struct {
	PortMappingLoadBalancerConfig
	Name          string
	ContainerPort int
	HostPort      int
	Protocol      string

	ServiceConnectDNS string
}

type VolumeConfig struct {
	Name          string
	ContainerPath string
	Size          string
}

type ContainerConfig struct {
	SSMParameterConfig
	Name               string
	Links              []string
	Image              string
	PortMappingConfigs []PortMappingConfig
	VolumeConfigs      []VolumeConfig
	EnvironmentConfig  []Environment
	Command            []string
	HealthCheck        string
	DependsOn          []DependsOn
}

type IAMConfig struct {
	Policies pulumi.StringArray
}

type EcsServiceV2Args struct {
	IAMConfig

	ClusterArn pulumi.StringInput

	SubnetIds pulumi.StringArrayInput

	Cpu    pulumi.StringInput
	Memory pulumi.StringInput

	ContainerConfigs      []ContainerConfig
	ServiceConnectEnabled bool
}

func createTaskExecutionRole(ctx *pulumi.Context, name string, args *EcsServiceV2Args, parent pulumi.Resource) (taskExecutionRole *iam.Role, err error) {

	tags := label.Tags(ctx, name)

	assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
		Statements: []iam.GetPolicyDocumentStatement{
			{
				Effect: pulumi.StringRef("Allow"),
				Principals: []iam.GetPolicyDocumentStatementPrincipal{
					{
						Type: "Service",
						Identifiers: []string{
							"ecs-tasks.amazonaws.com",
						},
					},
				},
				Actions: []string{
					"sts:AssumeRole",
				},
			},
		},
	}, pulumi.Parent(parent))
	if err != nil {
		return nil, err
	}

	taskExecutionRole, err = iam.NewRole(ctx, name, &iam.RoleArgs{
		Path:             pulumi.String("/"),
		Name:             pulumi.String(name),
		AssumeRolePolicy: pulumi.String(assumeRole.Json),
		Tags:             tags,
	}, pulumi.Parent(parent))
	if err != nil {
		return nil, err
	}

	for i, v := range args.IAMConfig.Policies {
		_, err = iam.NewRolePolicyAttachment(ctx, fmt.Sprintf("%s-execution-%d", name, i), &iam.RolePolicyAttachmentArgs{
			Role:      taskExecutionRole.Name,
			PolicyArn: v,
		}, pulumi.Parent(parent))
		if err != nil {
			return nil, err
		}

	}

	_, err = iam.NewRolePolicyAttachment(ctx, fmt.Sprintf("%s-task-execution-role", name), &iam.RolePolicyAttachmentArgs{
		Role:      taskExecutionRole.Name,
		PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"),
	}, pulumi.Parent(parent))
	if err != nil {
		return nil, err
	}

	return
}

func NewEcsService(ctx *pulumi.Context, name string, args *EcsServiceV2Args, opts ...pulumi.ResourceOption) (*EcsService, error) {

	tags := label.Tags(ctx, name)

	e := &EcsService{}
	err := ctx.RegisterComponentResource("pkg:index:ecs-service", name, e, opts...)
	if err != nil {
		return nil, err
	}

	taskExecutionRole, err := createTaskExecutionRole(ctx, name, args, e)

	e.TaskExecutionRole = taskExecutionRole.ToRoleOutput()

	if err != nil {
		return nil, err
	}

	var serviceLbConfigs ecs.ServiceLoadBalancerArray

	logGroupName := fmt.Sprintf("/ecs/%s", name)
	logGroup, err := cloudwatch.NewLogGroup(ctx, name, &cloudwatch.LogGroupArgs{
		Name:            pulumi.String(logGroupName),
		RetentionInDays: pulumi.Int(3),
		Tags:            tags,
	}, pulumi.Parent(e))
	if err != nil {
		return nil, err
	}

	var svcVolumeArray ecs.TaskDefinitionVolumeArray
	for i, containerConfigs := range args.ContainerConfigs {
		for j, volumeConfigs := range containerConfigs.VolumeConfigs {
			volumeConfigs.Name = fmt.Sprintf("%s-%s-%d", name, containerConfigs.Name, i)
			// wtf golang??
			args.ContainerConfigs[i].VolumeConfigs[j].Name = fmt.Sprintf("%s-%s-%d", name, containerConfigs.Name, i)
			svcVolumeArray = append(svcVolumeArray, &ecs.TaskDefinitionVolumeArgs{
				Name: pulumi.String(volumeConfigs.Name),
				DockerVolumeConfiguration: &ecs.TaskDefinitionVolumeDockerVolumeConfigurationArgs{
					Scope:         pulumi.String("shared"),
					Autoprovision: pulumi.Bool(true),
					Driver:        pulumi.String("rexray/ebs"),
					DriverOpts: pulumi.StringMap{
						"Size": pulumi.String(volumeConfigs.Size),
					},
				},
			})
		}
	}

	var containerDefns pulumix.Array[*ContainerDefinition]

	serviceConnectConfigs := ecs.ServiceServiceConnectConfigurationServiceArray{}
	for _, config := range args.ContainerConfigs {
		c := &ContainerDefinition{}
		c.Name = config.Name

		c.Links = config.Links
		c.DependsOn = config.DependsOn

		c.Command = config.Command

		if config.HealthCheck != "" {
			c.HealthCheck = &HealthCheck{
				Command: []string{"CMD-SHELL", config.HealthCheck},
			}
		}

		c.Environment = config.EnvironmentConfig
		c.Image = config.Image
		c.Essential = true

		c.LogConfiguration.LogDriver = "awslogs"

		c.LogConfiguration.Options.AwslogsRegion = "us-east-1"
		c.LogConfiguration.Options.AwslogsStreamPrefix = "ecs"

		containerPorts := PortMappings{}
		for _, portConfig := range config.PortMappingConfigs {

			containerPort := PortMapping{}
			containerPort.Name = portConfig.Name

			containerPort.HostPort = portConfig.HostPort
			containerPort.ContainerPort = portConfig.ContainerPort
			containerPort.Protocol = portConfig.Protocol

			if portConfig.ServiceConnectDNS != "" {
				serviceConnectConfig := &ecs.ServiceServiceConnectConfigurationServiceArgs{
					ClientAlias: &ecs.ServiceServiceConnectConfigurationServiceClientAliasArray{
						ecs.ServiceServiceConnectConfigurationServiceClientAliasArgs{
							DnsName: pulumi.String(portConfig.ServiceConnectDNS),
							Port:    pulumi.Int(portConfig.ContainerPort),
						},
					},
					DiscoveryName: pulumi.String(portConfig.ServiceConnectDNS),
					PortName:      pulumi.String(portConfig.Name),
				}
				serviceConnectConfigs = append(serviceConnectConfigs, serviceConnectConfig)
			}

			containerPorts = append(containerPorts, containerPort)

			if portConfig.PortMappingLoadBalancerConfig.Enabled {

				serviceConf := &ecs.ServiceLoadBalancerArgs{
					ContainerName:  pulumi.String(config.Name),
					ContainerPort:  pulumi.Int(portConfig.ContainerPort),
					TargetGroupArn: portConfig.PortMappingLoadBalancerConfig.TargetGroupArn,
				}
				if err != nil {
					return nil, err
				}
				serviceLbConfigs = append(serviceLbConfigs, serviceConf)
			}

		}

		c.PortMappings = containerPorts

		mountPoints := MountPoints{}

		for _, v := range config.VolumeConfigs {
			m := MountPoint{}
			m.ContainerPath = v.ContainerPath
			m.SourceVolume = v.Name
			mountPoints = append(mountPoints, m)
		}
		c.MountPoints = mountPoints

		cdfn := pulumix.Apply2(logGroup.Name, config.SSMParameterConfig.Names, func(logGroupName string, paramArns []string) *ContainerDefinition {
			var secrets Secrets
			for _, arn := range paramArns {

				k := strings.Split(arn, "/")
				name := k[len(k)-1]
				p := Secret{
					Name:      name,
					ValueFrom: arn,
				}
				secrets = append(secrets, p)
			}
			c.Secrets = secrets

			c.LogConfiguration.Options.AwslogsGroup = logGroupName

			return c

		})

		containerDefns = append(containerDefns, cdfn)

	}

	cdfnsString := pulumix.Apply(containerDefns, func(cd []*ContainerDefinition) string {
		str, err := json.Marshal(cd)
		if err != nil {
			panic(err)
		}
		return string(str)
	})

	taskDefn, err := ecs.NewTaskDefinition(ctx, name, &ecs.TaskDefinitionArgs{
		Family:               pulumi.String(name),
		ContainerDefinitions: pulumix.Cast[pulumi.StringOutput](cdfnsString),
		RequiresCompatibilities: pulumi.StringArray{
			pulumi.String("EC2"),
		},
		Cpu:              args.Cpu,
		Memory:           args.Memory,
		NetworkMode:      pulumi.String("bridge"),
		ExecutionRoleArn: taskExecutionRole.Arn,
		Volumes:          svcVolumeArray,
		Tags:             tags,
	}, pulumi.Parent(e), pulumi.IgnoreChanges([]string{"ContainerDefinitions"}))
	if err != nil {
		return nil, err
	}
	e.TaskDefinition = taskDefn.ToTaskDefinitionOutput()

	_, err = ecs.NewService(ctx, name, &ecs.ServiceArgs{
		Name:    pulumi.String(name),
		Cluster: args.ClusterArn,
		ServiceConnectConfiguration: &ecs.ServiceServiceConnectConfigurationArgs{
			Enabled:  pulumi.Bool(args.ServiceConnectEnabled),
			Services: serviceConnectConfigs,
		},
		TaskDefinition: taskDefn.Arn,
		DesiredCount:   pulumi.Int(1),
		LoadBalancers:  serviceLbConfigs,
		Tags:           tags,
	}, pulumi.Parent(e))

	ctx.RegisterResourceOutputs(e, pulumi.Map{})

	return e, nil

}
