package ecscluster

import (
	"github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/label"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicediscovery"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EcsCluster struct {
	pulumi.ResourceState
	Arn pulumi.StringOutput `pulumi:"arn"`
}

type EcsClusterArgs struct {
	SubnetIds           pulumi.StringArrayInput
	VpcSecurityGroupIds pulumi.StringArrayInput

	InstanceClass pulumi.StringInput

	MinSize     pulumi.IntInput
	MaxSize     pulumi.IntInput
	DesiredSize pulumi.IntInput
}

func NewEcsCluster(ctx *pulumi.Context, name string, args *EcsClusterArgs, opts ...pulumi.ResourceOption) (*EcsCluster, error) {

	tags := label.Tags(ctx, name)

	asgName := label.ID(ctx, fmt.Sprintf("%s-autoscaling-group", name))
	asgTags := label.Tags(ctx, fmt.Sprintf("%s-autoscaling-group", name))

	e := &EcsCluster{}
	err := ctx.RegisterComponentResource("pkg:index:ecs-cluster", name, e, opts...)
	if err != nil {
		return nil, err
	}

	httpNamespace, err := servicediscovery.NewHttpNamespace(ctx, name, &servicediscovery.HttpNamespaceArgs{
		Name:        pulumi.String(name),
		Description: pulumi.String(name),
		Tags:        tags,
	}, pulumi.Parent(e))

	assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
		Statements: []iam.GetPolicyDocumentStatement{
			{
				Effect: pulumi.StringRef("Allow"),
				Principals: []iam.GetPolicyDocumentStatementPrincipal{
					{
						Type: "Service",
						Identifiers: []string{
							"ec2.amazonaws.com",
						},
					},
				},
				Actions: []string{
					"sts:AssumeRole",
				},
			},
		},
	}, pulumi.Parent(e))
	if err != nil {
		return nil, err
	}

	policy, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
		Statements: []iam.GetPolicyDocumentStatement{
			{
				Effect: pulumi.StringRef("Allow"),
				Actions: []string{
					"ec2:AttachVolume",
					"ec2:CreateVolume",
					"ec2:CreateSnapshot",
					"ec2:CreateTags",
					"ec2:DeleteVolume",
					"ec2:DeleteSnapshot",
					"ec2:DescribeAvailabilityZones",
					"ec2:DescribeInstances",
					"ec2:DescribeVolumes",
					"ec2:DescribeVolumeAttribute",
					"ec2:DescribeVolumeStatus",
					"ec2:DescribeSnapshots",
					"ec2:CopySnapshot",
					"ec2:DescribeSnapshotAttribute",
					"ec2:DetachVolume",
					"ec2:ModifySnapshotAttribute",
					"ec2:ModifyVolumeAttribute",
					"ec2:DescribeTag",
				},
				Resources: []string{
					"*",
				},
			},
		},
	})

	instanceRole, err := iam.NewRole(ctx, asgName, &iam.RoleArgs{
		Path:             pulumi.String("/"),
		Name:             pulumi.String(asgName),
		AssumeRolePolicy: pulumi.String(assumeRole.Json),
		InlinePolicies: iam.RoleInlinePolicyArray{
			iam.RoleInlinePolicyArgs{
				Policy: pulumi.String(policy.Json),
				Name:   pulumi.String("rexray-ebs-permissions"),
			},
		},
		Tags: asgTags,
	}, pulumi.Parent(e))
	if err != nil {
		return nil, err
	}

	_, err = iam.NewRolePolicyAttachment(ctx, fmt.Sprintf("%s-ecs-role", asgName), &iam.RolePolicyAttachmentArgs{
		Role:      instanceRole.Name,
		PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"),
	}, pulumi.Parent(e))
	if err != nil {
		return nil, err
	}

	_, err = iam.NewRolePolicyAttachment(ctx, fmt.Sprintf("%s-ssm", asgName), &iam.RolePolicyAttachmentArgs{
		Role:      instanceRole.Name,
		PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"),
	}, pulumi.Parent(e))
	if err != nil {
		return nil, err
	}

	instanceProfile, err := iam.NewInstanceProfile(ctx, asgName, &iam.InstanceProfileArgs{
		Name: pulumi.String(asgName),
		Role: instanceRole.Name,
		Tags: asgTags,
	}, pulumi.Parent(e))

	if err != nil {
		return nil, err
	}

	ssmLookupResponse, err := ssm.LookupParameter(ctx, &ssm.LookupParameterArgs{
		Name: "/aws/service/ecs/optimized-ami/amazon-linux-2/recommended",
	}, pulumi.Parent(e))
	if err != nil {
		return nil, err
	}

	ecsAmi := &struct {
		ImageID string `json:"image_id"`
	}{}
	err = json.Unmarshal([]byte(ssmLookupResponse.Value), ecsAmi)
	if err != nil {
		return nil, err
	}

	userData := fmt.Sprintf(`#!/bin/bash
docker plugin install rexray/ebs --grant-all-permissions
echo ECS_CLUSTER=%s  >> /etc/ecs/ecs.config
`, name)

	userDataEncoded := base64.StdEncoding.EncodeToString([]byte(userData))

	launchTemplate, err := ec2.NewLaunchTemplate(ctx, asgName, &ec2.LaunchTemplateArgs{
		Name:                pulumi.String(asgName),
		ImageId:             pulumi.String(ecsAmi.ImageID),
		InstanceType:        args.InstanceClass,
		VpcSecurityGroupIds: args.VpcSecurityGroupIds,
		TagSpecifications: &ec2.LaunchTemplateTagSpecificationArray{
			&ec2.LaunchTemplateTagSpecificationArgs{
				ResourceType: pulumi.String("instance"),
				Tags:         asgTags,
			},
		},
		IamInstanceProfile: &ec2.LaunchTemplateIamInstanceProfileArgs{
			Arn: instanceProfile.Arn,
		},
		UserData: pulumi.String(userDataEncoded),
		Tags:     asgTags,
	}, pulumi.Parent(e), pulumi.IgnoreChanges([]string{"Tags"}))
	if err != nil {
		return nil, err
	}

	autoscalingGroup, err := autoscaling.NewGroup(ctx, asgName, &autoscaling.GroupArgs{
		Name:            pulumi.String(name),
		DesiredCapacity: args.DesiredSize,
		MaxSize:         args.MaxSize,
		MinSize:         args.MinSize,
		LaunchTemplate: &autoscaling.GroupLaunchTemplateArgs{
			Id:      launchTemplate.ID(),
			Version: pulumi.String("$Latest"),
		},
		VpcZoneIdentifiers: args.SubnetIds,
		Tags: &autoscaling.GroupTagArray{
			autoscaling.GroupTagArgs{
				Key:               pulumi.String("AmazonECSManaged"),
				PropagateAtLaunch: pulumi.Bool(true),
				Value:             pulumi.String(""),
			},
		},
	}, pulumi.Parent(e), pulumi.IgnoreChanges([]string{"Tags"}))
	if err != nil {
		return nil, err
	}

	capacityProvider, err := ecs.NewCapacityProvider(ctx, name, &ecs.CapacityProviderArgs{
		Tags: tags,
		AutoScalingGroupProvider: &ecs.CapacityProviderAutoScalingGroupProviderArgs{
			AutoScalingGroupArn: autoscalingGroup.Arn,
			ManagedScaling: &ecs.CapacityProviderAutoScalingGroupProviderManagedScalingArgs{
				MaximumScalingStepSize: pulumi.Int(1000),
				MinimumScalingStepSize: pulumi.Int(1),
				Status:                 pulumi.String("ENABLED"),
				TargetCapacity:         pulumi.Int(100),
			},
		},
	}, pulumi.Parent(e), pulumi.DeleteBeforeReplace(true))
	if err != nil {
		return nil, err
	}

	ecsCluster, err := ecs.NewCluster(ctx, name, &ecs.ClusterArgs{
		Name: pulumi.String(name),
		Settings: ecs.ClusterSettingArray{
			&ecs.ClusterSettingArgs{
				Name:  pulumi.String("containerInsights"),
				Value: pulumi.String("enabled"),
			},
		},
		ServiceConnectDefaults: &ecs.ClusterServiceConnectDefaultsArgs{
			Namespace: httpNamespace.Arn,
		},
		Tags: tags,
	}, pulumi.Parent(e))

	if err != nil {
		return nil, err
	}

	_, err = ecs.NewClusterCapacityProviders(ctx, name, &ecs.ClusterCapacityProvidersArgs{
		ClusterName: ecsCluster.Name,
		CapacityProviders: pulumi.StringArray{
			capacityProvider.Name,
		},
	}, pulumi.Parent(e))
	if err != nil {
		return nil, err
	}

	e.Arn = ecsCluster.Arn

	ctx.RegisterResourceOutputs(e, pulumi.Map{
		"arn": ecsCluster.Arn,
	})

	return e, nil

}
