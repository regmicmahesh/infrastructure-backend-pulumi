package postgresdatabase

import (
	"github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/label"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PostgresDatabase struct {
	pulumi.ResourceState

	Instance rds.InstanceOutput

	Password     pulumi.StringOutput `pulumi:"password"`
	Username     pulumi.StringOutput `pulumi:"username"`
	Endpoint     pulumi.StringOutput `pulumi:"endpoint"`
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
}

type PostgresDatabaseArgs struct {
	AllocatedStorage    pulumi.IntInput
	DatabaseName        pulumi.StringInput
	SubnetIds           pulumi.StringArrayInput
	VpcId               pulumi.StringInput
	InstanceClass       pulumi.StringInput
	VpcSecurityGroupIds pulumi.StringArrayInput
}

func NewPostgresDatabase(ctx *pulumi.Context, name string, args *PostgresDatabaseArgs, opts ...pulumi.ResourceOption) (*PostgresDatabase, error) {

	tags := label.Tags(ctx, name)

	p := &PostgresDatabase{}
	err := ctx.RegisterComponentResource("pkg:index:postgres-database", name, p, opts...)
	if err != nil {
		return nil, err
	}

	dbSubnetGroup, err := rds.NewSubnetGroup(ctx, name, &rds.SubnetGroupArgs{
		Name:      pulumi.String(name),
		SubnetIds: args.SubnetIds,
		Tags:      tags,
	}, pulumi.Parent(p))

	if err != nil {
		return nil, err
	}

	dbParameterGroup, err := rds.NewParameterGroup(ctx, name, &rds.ParameterGroupArgs{
		Name:       pulumi.String(name),
		Family:     pulumi.String("postgres16"),
		Parameters: rds.ParameterGroupParameterArray{},
	}, pulumi.Parent(p))

	if err != nil {
		return nil, err
	}

	password, err := random.NewRandomPassword(ctx, fmt.Sprintf("%s-password", name), &random.RandomPasswordArgs{
		Length:  pulumi.Int(16),
		Special: pulumi.Bool(false),
	}, pulumi.Parent(p))

	if err != nil {
		return nil, err
	}

	instance, err := rds.NewInstance(ctx, name, &rds.InstanceArgs{
		AllocatedStorage:                   args.AllocatedStorage,
		Identifier:                         pulumi.String(name),
		DbName:                             args.DatabaseName,
		Engine:                             pulumi.String("postgres"),
		EngineVersion:                      pulumi.String("16.1"),
		InstanceClass:                      args.InstanceClass,
		Password:                           password.Result,
		SkipFinalSnapshot:                  pulumi.Bool(true),
		ApplyImmediately:                   pulumi.BoolPtr(true),
		Username:                           args.DatabaseName,
		DbSubnetGroupName:                  dbSubnetGroup.Name,
		ParameterGroupName:                 dbParameterGroup.Name,
		PerformanceInsightsEnabled:         pulumi.Bool(true),
		PerformanceInsightsRetentionPeriod: pulumi.Int(7),
		VpcSecurityGroupIds:                args.VpcSecurityGroupIds,
		Tags:                               tags,
	}, pulumi.Parent(p))
	if err != nil {
		return nil, err
	}

	p.Password = password.Result
	p.Username = instance.Username
	p.Endpoint = instance.Address
	p.DatabaseName = instance.DbName

	p.Instance = instance.ToInstanceOutput()

	return p, nil

}
