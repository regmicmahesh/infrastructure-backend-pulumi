package mongodatabase

import (
	"github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/label"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/docdb"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MongoDatabase struct {
	pulumi.ResourceState

	Cluster docdb.ClusterOutput
}

type MongoDatabaseArgs struct {
	DatabaseName        pulumi.StringInput
	SubnetIds           pulumi.StringArrayInput
	VpcId               pulumi.StringInput
	InstanceClass       pulumi.StringInput
	VpcSecurityGroupIds pulumi.StringArray
}

func NewMongoDatabase(ctx *pulumi.Context, name string, args *MongoDatabaseArgs, opts ...pulumi.ResourceOption) (*MongoDatabase, error) {

	tags := label.Tags(ctx, name)

	m := &MongoDatabase{}
	err := ctx.RegisterComponentResource("pkg:index:mongo-database", name, m, opts...)
	if err != nil {
		return nil, err
	}

	dbSubnetGroup, err := docdb.NewSubnetGroup(ctx, name, &docdb.SubnetGroupArgs{
		Name:      pulumi.String(name),
		SubnetIds: args.SubnetIds,
		Tags:      tags,
	}, pulumi.Parent(m))

	if err != nil {
		return nil, err
	}

	dbParameterGroup, err := docdb.NewClusterParameterGroup(ctx, name, &docdb.ClusterParameterGroupArgs{
		Family: pulumi.String("docdb5.0"),
		Parameters: docdb.ClusterParameterGroupParameterArray{
			docdb.ClusterParameterGroupParameterArgs{
				Name:  pulumi.String("tls"),
				Value: pulumi.String("disabled"),
			},
		},
	}, pulumi.Parent(m))

	if err != nil {
		return nil, err
	}

	password, err := random.NewRandomPassword(ctx, fmt.Sprintf("%s-password", name), &random.RandomPasswordArgs{
		Length:  pulumi.Int(16),
		Special: pulumi.Bool(false),
	}, pulumi.Parent(m))

	if err != nil {
		return nil, err
	}

	cluster, err := docdb.NewCluster(ctx, name, &docdb.ClusterArgs{
		DbSubnetGroupName:           dbSubnetGroup.Name,
		DbClusterParameterGroupName: dbParameterGroup.Name,
		ClusterIdentifier:           pulumi.String(name),
		MasterUsername:              args.DatabaseName,
		MasterPassword:              password.Result,
		VpcSecurityGroupIds:         args.VpcSecurityGroupIds,
		SkipFinalSnapshot:           pulumi.Bool(true),
		Tags:                        tags,
	}, pulumi.Parent(m))
	if err != nil {
		return nil, err
	}

	m.Cluster = cluster.ToClusterOutput()

	_, err = docdb.NewClusterInstance(ctx, name, &docdb.ClusterInstanceArgs{
		Identifier:        pulumi.String(name),
		ClusterIdentifier: cluster.ID(),
		InstanceClass:     args.InstanceClass,
		Tags:              tags,
	}, pulumi.Parent(m))
	if err != nil {
		return nil, err
	}

	return m, nil

}
