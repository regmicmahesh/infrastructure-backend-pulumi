package main

import (
	"fmt"
	ecsservicev2 "github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/ecs-service-v2"
	label "github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/label"
	ssmparameters "github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/ssm-parameters"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

type AppEnvironmentVariables struct {
	ViteBaseAPIURL string `yaml:"viteBaseAPIURL"`
}

type GraphServiceEnvironmentVariables struct {
	Ipfs             string             `yaml:"ipfs"`
	PostgresHost     pulumi.StringInput `yaml:"-"`
	PostgresUser     pulumi.StringInput `yaml:"-"`
	PostgresPassword pulumi.StringInput `yaml:"-"`
	// Not Used
	PostgresDatabase string `yaml:"-"`
}

type ApiServiceEnvironmentVariables struct {
	NodeEnv    string             `yaml:"nodeEnv"`
	Port       string             `yaml:"port"`
	MongodbUrl pulumi.StringInput `yaml:"-"`
}

func main() {

	pulumi.Run(func(ctx *pulumi.Context) error {

		vpcRefName := fmt.Sprintf("organization/network/%s", ctx.Stack())
		vpcRef, err := pulumi.NewStackReference(ctx, vpcRefName, nil)
		if err != nil {
			return err
		}

		computeRefName := fmt.Sprintf("organization/compute/%s", ctx.Stack())
		computeRef, err := pulumi.NewStackReference(ctx, computeRefName, nil)
		if err != nil {
			return err
		}

		ingressRefName := fmt.Sprintf("organization/ingress/%s", ctx.Stack())
		ingressRef, err := pulumi.NewStackReference(ctx, ingressRefName, nil)
		if err != nil {
			return err
		}

		dataRefName := fmt.Sprintf("organization/data/%s", ctx.Stack())
		dataRef, err := pulumi.NewStackReference(ctx, dataRefName, nil)
		if err != nil {
			return err
		}

		graphServiceConf := config.New(ctx, "graph-service")

		graphServiceEnvVars := &GraphServiceEnvironmentVariables{}
		graphServiceConf.RequireObject("env-vars", graphServiceEnvVars)
		fmt.Println(graphServiceEnvVars)

		graphServiceEnvVars.PostgresUser = dataRef.GetStringOutput(pulumi.String("graphDatabaseUsername"))
		graphServiceEnvVars.PostgresPassword = dataRef.GetStringOutput(pulumi.String("graphDatabasePassword"))
		graphServiceEnvVars.PostgresHost = dataRef.GetStringOutput(pulumi.String("graphDatabaseEndpoint"))

		graphNodeParametersName := label.ID(ctx, "graph-node")
		graphNodeParameters, err := ssmparameters.NewSSMParameters(ctx, graphNodeParametersName, &ssmparameters.SSMParametersArgs{

			Secrets: pulumi.StringMap{
				"postgres_host": graphServiceEnvVars.PostgresHost,
				"postgres_user": graphServiceEnvVars.PostgresUser,
				"postgres_pass": graphServiceEnvVars.PostgresPassword,
				"postgres_db":   pulumi.String("graph_v2"),
			},
		})

		if err != nil {
			return err
		}

		var graphNodeParameterNames pulumi.StringArray
		graphNodeParameterNames = append(graphNodeParameterNames, graphNodeParameters.VariableArns...)
		graphNodeParameterNames = append(graphNodeParameterNames, graphNodeParameters.SecretArns...)

		graphServiceResourceName := label.ID(ctx, "graph-service")
		_, err = ecsservicev2.NewEcsService(ctx, graphServiceResourceName, &ecsservicev2.EcsServiceV2Args{
			ClusterArn: computeRef.GetStringOutput(pulumi.String("clusterArn")),
			SubnetIds:  vpcRef.GetOutput(pulumi.String("privateSubnetIds")).AsStringArrayOutput(),
			IAMConfig: ecsservicev2.IAMConfig{
				Policies: pulumi.StringArray{
					graphNodeParameters.AccessPolicyArn,
				},
			},
			Cpu:    pulumi.String(graphServiceConf.Require("cpu")),
			Memory: pulumi.String(graphServiceConf.Require("memory")),
			ContainerConfigs: []ecsservicev2.ContainerConfig{
				{
					Name: "graph",
					SSMParameterConfig: ecsservicev2.SSMParameterConfig{
						Names: graphNodeParameterNames,
					},
					Image: graphServiceConf.Require("image"),
					PortMappingConfigs: []ecsservicev2.PortMappingConfig{
						{
							Name:          "graphql",
							Protocol:      "tcp",
							HostPort:      0,
							ContainerPort: 8000,
							PortMappingLoadBalancerConfig: ecsservicev2.PortMappingLoadBalancerConfig{
								Enabled:        true,
								TargetGroupArn: ingressRef.GetStringOutput(pulumi.String("graphServiceTargetGroupArn")),
							},
						},
						{
							Name:          "jsonrpc",
							Protocol:      "tcp",
							HostPort:      0,
							ContainerPort: 8020,
							PortMappingLoadBalancerConfig: ecsservicev2.PortMappingLoadBalancerConfig{
								Enabled:        true,
								TargetGroupArn: ingressRef.GetStringOutput(pulumi.String("graphRpcServiceTargetGroupArn")),
							},
						},
					},
				},
			},
		})

		apiServiceConf := config.New(ctx, "api-service")
		apiServiceEnvVars := &ApiServiceEnvironmentVariables{}
		apiServiceConf.RequireObject("env-vars", apiServiceEnvVars)

		apiServiceEnvVars.MongodbUrl = dataRef.GetStringOutput(pulumi.String("apiDatabaseUri"))

		apiServiceParametersResourceName := label.ID(ctx, "api-service")
		apiServiceParameters, err := ssmparameters.NewSSMParameters(ctx, apiServiceParametersResourceName, &ssmparameters.SSMParametersArgs{
			Variables: map[string]pulumi.StringInput{
				"MONGODB_URL": apiServiceEnvVars.MongodbUrl,
				"NODE_ENV":    pulumi.String(apiServiceEnvVars.NodeEnv),
				"PORT":        pulumi.String(apiServiceEnvVars.Port),
			},
		})

		if err != nil {
			return err
		}

		var apiServiceParameterNames pulumi.StringArray
		apiServiceParameterNames = append(apiServiceParameterNames, apiServiceParameters.VariableArns...)
		apiServiceParameterNames = append(apiServiceParameterNames, apiServiceParameters.SecretArns...)

		apiServiceResourceName := label.ID(ctx, "api-service")
		_, err = ecsservicev2.NewEcsService(ctx, apiServiceResourceName, &ecsservicev2.EcsServiceV2Args{
			ClusterArn: computeRef.GetStringOutput(pulumi.String("clusterArn")),
			SubnetIds:  vpcRef.GetOutput(pulumi.String("privateSubnetIds")).AsStringArrayOutput(),
			IAMConfig: ecsservicev2.IAMConfig{
				Policies: pulumi.StringArray{
					apiServiceParameters.AccessPolicyArn,
				},
			},
			Cpu:    pulumi.String(apiServiceConf.Require("cpu")),
			Memory: pulumi.String(apiServiceConf.Require("memory")),
			ContainerConfigs: []ecsservicev2.ContainerConfig{
				{
					Name: "backend",
					SSMParameterConfig: ecsservicev2.SSMParameterConfig{
						Names: apiServiceParameterNames,
					},
					Image: apiServiceConf.Require("image"),
					PortMappingConfigs: []ecsservicev2.PortMappingConfig{
						{
							Name:          "backend",
							Protocol:      "tcp",
							HostPort:      0,
							ContainerPort: 3000,
							PortMappingLoadBalancerConfig: ecsservicev2.PortMappingLoadBalancerConfig{
								Enabled:        true,
								TargetGroupArn: ingressRef.GetStringOutput(pulumi.String("apiServiceTargetGroupArn")),
							},
						},
					},
				},
			},
		})

		if err != nil {
			return err
		}
		appConf := config.New(ctx, "app")
		appEnvVars := &AppEnvironmentVariables{}
		appConf.RequireObject("env-vars", appEnvVars)

		appParametersResourceName := label.ID(ctx, "app")
		_, err = ssmparameters.NewSSMParameters(ctx, appParametersResourceName, &ssmparameters.SSMParametersArgs{
			Variables: map[string]pulumi.StringInput{
				"VITE_BASE_API_URL": pulumi.String(appEnvVars.ViteBaseAPIURL),
			},
		})

		if err != nil {
			return err
		}

		return nil

	})
}
