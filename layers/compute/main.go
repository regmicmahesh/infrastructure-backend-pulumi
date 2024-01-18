package main

import (
	ecscluster "github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/ecs-cluster"
	label "github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/label"
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {

	pulumi.Run(func(ctx *pulumi.Context) error {

		vpcRefName := fmt.Sprintf("organization/network/%s", ctx.Stack())
		vpcRef, err := pulumi.NewStackReference(ctx, vpcRefName, nil)
		if err != nil {
			return err
		}

		clusterConf := config.New(ctx, "cluster")

		ecsClusterResourceName := label.ID(ctx, "cluster")
    cluster, err := ecscluster.NewEcsCluster(ctx, ecsClusterResourceName, &ecscluster.EcsClusterArgs{
			SubnetIds: vpcRef.GetOutput(pulumi.String("privateSubnetIds")).AsStringArrayOutput(),
			VpcSecurityGroupIds: pulumi.StringArray{
				vpcRef.GetStringOutput(pulumi.String("clusterSecurityGroupId")),
			},
			InstanceClass: pulumi.String(clusterConf.Require("instanceClass")),
			MinSize:       pulumi.Int(1),
			MaxSize:       pulumi.Int(1),
			DesiredSize:   pulumi.Int(1),
		})
		if err != nil {
			return err
		}

		ipfsClusterResourceName := label.ID(ctx, "ipfs-cluster")
		ipfsCluster, err := ecscluster.NewEcsCluster(ctx, ipfsClusterResourceName, &ecscluster.EcsClusterArgs{
			SubnetIds: vpcRef.GetOutput(pulumi.String("privateSubnetIds")).AsStringArrayOutput(),
			VpcSecurityGroupIds: pulumi.StringArray{
				vpcRef.GetStringOutput(pulumi.String("clusterSecurityGroupId")),
			},
			InstanceClass: pulumi.String(clusterConf.Require("instanceClass")),
			MinSize:       pulumi.Int(2),
			MaxSize:       pulumi.Int(2),
			DesiredSize:   pulumi.Int(2),
		})
		if err != nil {
			return err
		}
		ctx.Export("clusterArn", cluster.Arn)
		ctx.Export("ipfsClusterArn", ipfsCluster.Arn)

		return nil

	})
}
