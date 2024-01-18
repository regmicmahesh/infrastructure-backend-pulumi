package main

import (
	label "github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/label"
	securitygroup "github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/security-group"
	"fmt"

	ec2x "github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func createVpc(ctx *pulumi.Context, name string, config *config.Config) (*ec2x.Vpc, error) {

	cidrBlock := config.Require("cidrBlock")
	azCount := config.RequireInt("azCount")
	natGatewayStrategy := config.Require("natGatewayStrategy")

	vpc, err := ec2x.NewVpc(ctx, name, &ec2x.VpcArgs{
		CidrBlock:                 pulumi.StringRef(cidrBlock),
		NumberOfAvailabilityZones: pulumi.IntRef(azCount),
		NatGateways: &ec2x.NatGatewayConfigurationArgs{
			Strategy: ec2x.NatGatewayStrategy(natGatewayStrategy),
		},
		Tags: label.Tags(ctx, name),
	})

	return vpc, err
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		vpcConf := config.New(ctx, "vpc")

		vpcResourceName := label.ID(ctx, "network")
		vpc, err := createVpc(ctx, vpcResourceName, vpcConf)
		if err != nil {
			return err
		}

		loadBalancerSecurityGroupResourceName := label.ID(ctx, "load-balancer-security-group")
		loadBalancerSecurityGroup, err := securitygroup.NewSecurityGroup(ctx, loadBalancerSecurityGroupResourceName, &securitygroup.SecurityGroupArgs{
			VpcId: vpc.VpcId,
			IngressCIDRArgs: []securitygroup.IngressCIDRArgs{
				{
					FromPort: 443,
					ToPort:   443,
					Protocol: "tcp",
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"), // allow incoming traffic from internet
					},
				},
				{
					FromPort: 80,
					ToPort:   80,
					Protocol: "tcp",
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"), // allow incoming traffic from internet
					},
				},
			},
		})
		if err != nil {
			return err
		}

		internalLoadBalancerSecurityGroupResourceName := label.ID(ctx, "internal-load-balancer-security-group")
		internalLoadBalancerSecurityGroup, err := securitygroup.NewSecurityGroup(ctx, internalLoadBalancerSecurityGroupResourceName, &securitygroup.SecurityGroupArgs{
			VpcId: vpc.VpcId,
			IngressCIDRArgs: []securitygroup.IngressCIDRArgs{
				{
					FromPort: 443,
					ToPort:   443,
					Protocol: "tcp",
					CidrBlocks: pulumi.StringArray{
						vpc.Vpc.CidrBlock(), // allow incoming traffic from vpc
					},
				},
				{
					FromPort: 80,
					ToPort:   80,
					Protocol: "tcp",
					CidrBlocks: pulumi.StringArray{
						vpc.Vpc.CidrBlock(), // allow incoming traffic from vpc
					},
				},
			},
		})
		if err != nil {
			return err
		}

		clusterSecurityGroupResourceName := label.ID(ctx, "cluster-security-group")
		clusterSecurityGroup, err := securitygroup.NewSecurityGroup(ctx, clusterSecurityGroupResourceName, &securitygroup.SecurityGroupArgs{
			VpcId: vpc.VpcId,
			IngressSecurityGroupArgs: []securitygroup.IngressSecurityGroupArgs{
				{
					FromPort:        -1,
					ToPort:          -1,
					Protocol:        "all",
					SecurityGroupId: loadBalancerSecurityGroup.SecurityGroupId, // allow incoming traffic from load balancer
				},
				{
					FromPort:        -1,
					ToPort:          -1,
					Protocol:        "all",
					SecurityGroupId: internalLoadBalancerSecurityGroup.SecurityGroupId, // allow incoming traffic from internal load balancer
				},
			},
			EgressCIDRArgs: []securitygroup.IngressCIDRArgs{
				{
					FromPort: -1,
					ToPort:   -1,
					Protocol: "all",
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"), // allow outgoing traffic to internet
					},
				},
			},
		})
		if err != nil {
			return err
		}
		clusterIngressToClusterResourceName := fmt.Sprintf("%s-ingress-cluster", clusterSecurityGroupResourceName)
		_, err = securitygroup.NewIngressSecurityGroupRule(ctx, clusterIngressToClusterResourceName, &securitygroup.IngressSecurityGroupRuleArgs{
			VpcId:           vpc.VpcId,
			SecurityGroupId: clusterSecurityGroup.SecurityGroupId,
			IngressSecurityGroupArgs: securitygroup.IngressSecurityGroupArgs{
				FromPort:        -1, // cluster has dynamic port mapping
				ToPort:          -1,
				Protocol:        "all",
				SecurityGroupId: clusterSecurityGroup.SecurityGroupId, // allow cluster to cluster communication
			},
		})
		if err != nil {
			return err
		}

		loadBalancerEgressToClusterResourceName := fmt.Sprintf("%s-egress-cluster", loadBalancerSecurityGroupResourceName)
		_, err = securitygroup.NewEgressSecurityGroupRule(ctx, loadBalancerEgressToClusterResourceName, &securitygroup.EgressSecurityGroupRuleArgs{
			VpcId:           vpc.VpcId,
			SecurityGroupId: loadBalancerSecurityGroup.SecurityGroupId,
			EgressSecurityGroupArgs: securitygroup.EgressSecurityGroupArgs{
				FromPort:        -1, // cluster has dynamic port mapping
				ToPort:          -1,
				Protocol:        "all",
				SecurityGroupId: clusterSecurityGroup.SecurityGroupId, // allow load balancer to send requests to the cluster
			},
		})
		if err != nil {
			return err
		}

		internalLoadBalancerEgressToClusterResourceName := fmt.Sprintf("%s-egress-cluster", internalLoadBalancerSecurityGroupResourceName)
		_, err = securitygroup.NewEgressSecurityGroupRule(ctx, internalLoadBalancerEgressToClusterResourceName, &securitygroup.EgressSecurityGroupRuleArgs{
			VpcId:           vpc.VpcId,
			SecurityGroupId: internalLoadBalancerSecurityGroup.SecurityGroupId,
			EgressSecurityGroupArgs: securitygroup.EgressSecurityGroupArgs{
				FromPort:        -1, // cluster has dynamic port mapping
				ToPort:          -1,
				Protocol:        "all",
				SecurityGroupId: clusterSecurityGroup.SecurityGroupId, // allow load balancer to send requests to the cluster
			},
		})
		if err != nil {
			return err
		}

		graphDatabaseSecurityGroupResourceName := label.ID(ctx, "graph-database-security-group")
		graphDatabaseSecurityGroup, err := securitygroup.NewSecurityGroup(ctx, graphDatabaseSecurityGroupResourceName, &securitygroup.SecurityGroupArgs{
			VpcId: vpc.VpcId,
			IngressSecurityGroupArgs: []securitygroup.IngressSecurityGroupArgs{
				{
					FromPort:        5432,
					ToPort:          5432,
					Protocol:        "tcp",
					SecurityGroupId: clusterSecurityGroup.SecurityGroupId, // allow incoming traffic from cluster
				},
			},
		})
		if err != nil {
			return err
		}

		apiDatabaseSecurityGroupResourceName := label.ID(ctx, "api-database-security-group")
		apiDatabaseSecurityGroup, err := securitygroup.NewSecurityGroup(ctx, apiDatabaseSecurityGroupResourceName, &securitygroup.SecurityGroupArgs{
			VpcId: vpc.VpcId,
			IngressSecurityGroupArgs: []securitygroup.IngressSecurityGroupArgs{
				{
					FromPort:        27017,
					ToPort:          27017,
					Protocol:        "tcp",
					SecurityGroupId: clusterSecurityGroup.SecurityGroupId, // allow incoming traffic from cluster
				},
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("vpcId", vpc.VpcId)
		ctx.Export("vpcCidr", vpc.Vpc.CidrBlock())

		ctx.Export("publicSubnetIds", vpc.PublicSubnetIds)
		ctx.Export("privateSubnetIds", vpc.PrivateSubnetIds)

		ctx.Export("apiDatabaseSecurityGroupId", apiDatabaseSecurityGroup.SecurityGroupId)
		ctx.Export("graphDatabaseSecurityGroupId", graphDatabaseSecurityGroup.SecurityGroupId)
		ctx.Export("loadBalancerSecurityGroupId", loadBalancerSecurityGroup.SecurityGroupId)
		ctx.Export("internalLoadBalancerSecurityGroupId", internalLoadBalancerSecurityGroup.SecurityGroupId)
		ctx.Export("clusterSecurityGroupId", clusterSecurityGroup.SecurityGroupId)

		return nil

	})
}
