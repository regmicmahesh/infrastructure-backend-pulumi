package main

import (
	label "github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/label"
	mongodatabase "github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/mongo-database"
	postgresdatabase "github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/postgres-database"
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

		graphDatabaseConf := config.New(ctx, "graph-database")

		graphDatabaseResourceName := label.ID(ctx, "graph-database")
		graphDatabase, err := postgresdatabase.NewPostgresDatabase(ctx, graphDatabaseResourceName, &postgresdatabase.PostgresDatabaseArgs{
			VpcSecurityGroupIds: pulumi.StringArray{
				vpcRef.GetStringOutput(pulumi.String("graphDatabaseSecurityGroupId")),
			},
			DatabaseName:     pulumi.String("graph"),
			SubnetIds:        vpcRef.GetOutput(pulumi.String("privateSubnetIds")).AsStringArrayOutput(),
			VpcId:            vpcRef.GetStringOutput(pulumi.String("vpcId")),
			InstanceClass:    pulumi.String(graphDatabaseConf.Require("instanceClass")),
			AllocatedStorage: pulumi.Int(graphDatabaseConf.RequireInt("storage")),
		})

		if err != nil {
			return err
		}

		ctx.Export("graphDatabaseEndpoint", graphDatabase.Instance.Address())
		ctx.Export("graphDatabaseName", graphDatabase.Instance.DbName())
		ctx.Export("graphDatabaseUsername", graphDatabase.Instance.Username())
		ctx.Export("graphDatabasePassword", graphDatabase.Instance.Password())

		apiDatabaseConf := config.New(ctx, "api-database")


		apiDatabaseResourceName := label.ID(ctx, "api-database")
		apiDatabase, err := mongodatabase.NewMongoDatabase(ctx, apiDatabaseResourceName, &mongodatabase.MongoDatabaseArgs{
			DatabaseName: pulumi.String("graph"),
			VpcSecurityGroupIds: pulumi.StringArray{
				vpcRef.GetStringOutput(pulumi.String("apiDatabaseSecurityGroupId")),
			},
			SubnetIds:     vpcRef.GetOutput(pulumi.String("privateSubnetIds")).AsStringArrayOutput(),
			VpcId:         vpcRef.GetStringOutput(pulumi.String("vpcId")),
			InstanceClass: pulumi.String(apiDatabaseConf.Require("instanceClass")),
		})
		if err != nil {
			return err
		}

		endpoint, username, password := apiDatabase.Cluster.Endpoint(), apiDatabase.Cluster.MasterUsername(), apiDatabase.Cluster.MasterPassword().Elem()

		mongoUri := pulumi.All(endpoint, username, password).ApplyT(func(args []any) string {
			endpoint := args[0].(string)
			username := args[1].(string)
			password := args[2].(string)

			return fmt.Sprintf("mongodb://%s:%s@%s:27017/?replicaSet=rs0&readPreference=secondaryPreferred&retryWrites=false", username, password, endpoint)

		}).(pulumi.StringOutput)

		ctx.Export("apiDatabaseUri", mongoUri)

		return nil
	})
}
