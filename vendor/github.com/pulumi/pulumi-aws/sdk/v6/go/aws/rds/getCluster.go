// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about an RDS cluster.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rds.LookupCluster(ctx, &rds.LookupClusterArgs{
//				ClusterIdentifier: "clusterName",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterResult
	err := ctx.Invoke("aws:rds/getCluster:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCluster.
type LookupClusterArgs struct {
	// Cluster identifier of the RDS cluster.
	ClusterIdentifier string `pulumi:"clusterIdentifier"`
	// A map of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getCluster.
type LookupClusterResult struct {
	Arn                              string   `pulumi:"arn"`
	AvailabilityZones                []string `pulumi:"availabilityZones"`
	BacktrackWindow                  int      `pulumi:"backtrackWindow"`
	BackupRetentionPeriod            int      `pulumi:"backupRetentionPeriod"`
	ClusterIdentifier                string   `pulumi:"clusterIdentifier"`
	ClusterMembers                   []string `pulumi:"clusterMembers"`
	ClusterResourceId                string   `pulumi:"clusterResourceId"`
	DatabaseName                     string   `pulumi:"databaseName"`
	DbClusterParameterGroupName      string   `pulumi:"dbClusterParameterGroupName"`
	DbSubnetGroupName                string   `pulumi:"dbSubnetGroupName"`
	DbSystemId                       string   `pulumi:"dbSystemId"`
	EnabledCloudwatchLogsExports     []string `pulumi:"enabledCloudwatchLogsExports"`
	Endpoint                         string   `pulumi:"endpoint"`
	Engine                           string   `pulumi:"engine"`
	EngineMode                       string   `pulumi:"engineMode"`
	EngineVersion                    string   `pulumi:"engineVersion"`
	FinalSnapshotIdentifier          string   `pulumi:"finalSnapshotIdentifier"`
	HostedZoneId                     string   `pulumi:"hostedZoneId"`
	IamDatabaseAuthenticationEnabled bool     `pulumi:"iamDatabaseAuthenticationEnabled"`
	IamRoles                         []string `pulumi:"iamRoles"`
	// The provider-assigned unique ID for this managed resource.
	Id                          string                       `pulumi:"id"`
	KmsKeyId                    string                       `pulumi:"kmsKeyId"`
	MasterUserSecrets           []GetClusterMasterUserSecret `pulumi:"masterUserSecrets"`
	MasterUsername              string                       `pulumi:"masterUsername"`
	NetworkType                 string                       `pulumi:"networkType"`
	Port                        int                          `pulumi:"port"`
	PreferredBackupWindow       string                       `pulumi:"preferredBackupWindow"`
	PreferredMaintenanceWindow  string                       `pulumi:"preferredMaintenanceWindow"`
	ReaderEndpoint              string                       `pulumi:"readerEndpoint"`
	ReplicationSourceIdentifier string                       `pulumi:"replicationSourceIdentifier"`
	StorageEncrypted            bool                         `pulumi:"storageEncrypted"`
	// A map of tags assigned to the resource.
	Tags                map[string]string `pulumi:"tags"`
	VpcSecurityGroupIds []string          `pulumi:"vpcSecurityGroupIds"`
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterResult, error) {
			args := v.(LookupClusterArgs)
			r, err := LookupCluster(ctx, &args, opts...)
			var s LookupClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterResultOutput)
}

// A collection of arguments for invoking getCluster.
type LookupClusterOutputArgs struct {
	// Cluster identifier of the RDS cluster.
	ClusterIdentifier pulumi.StringInput `pulumi:"clusterIdentifier"`
	// A map of tags assigned to the resource.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}

// A collection of values returned by getCluster.
type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Arn }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o LookupClusterResultOutput) BacktrackWindow() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClusterResult) int { return v.BacktrackWindow }).(pulumi.IntOutput)
}

func (o LookupClusterResultOutput) BackupRetentionPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClusterResult) int { return v.BackupRetentionPeriod }).(pulumi.IntOutput)
}

func (o LookupClusterResultOutput) ClusterIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterIdentifier }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) ClusterMembers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.ClusterMembers }).(pulumi.StringArrayOutput)
}

func (o LookupClusterResultOutput) ClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterResourceId }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) DbClusterParameterGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.DbClusterParameterGroupName }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) DbSubnetGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.DbSubnetGroupName }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) DbSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.DbSystemId }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) EnabledCloudwatchLogsExports() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.EnabledCloudwatchLogsExports }).(pulumi.StringArrayOutput)
}

func (o LookupClusterResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Engine }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) EngineMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.EngineMode }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) EngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.EngineVersion }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) FinalSnapshotIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.FinalSnapshotIdentifier }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) HostedZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.HostedZoneId }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) IamDatabaseAuthenticationEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClusterResult) bool { return v.IamDatabaseAuthenticationEnabled }).(pulumi.BoolOutput)
}

func (o LookupClusterResultOutput) IamRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.IamRoles }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) KmsKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.KmsKeyId }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) MasterUserSecrets() GetClusterMasterUserSecretArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []GetClusterMasterUserSecret { return v.MasterUserSecrets }).(GetClusterMasterUserSecretArrayOutput)
}

func (o LookupClusterResultOutput) MasterUsername() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.MasterUsername }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.NetworkType }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClusterResult) int { return v.Port }).(pulumi.IntOutput)
}

func (o LookupClusterResultOutput) PreferredBackupWindow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.PreferredBackupWindow }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) PreferredMaintenanceWindow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.PreferredMaintenanceWindow }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) ReaderEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ReaderEndpoint }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) ReplicationSourceIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ReplicationSourceIdentifier }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) StorageEncrypted() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClusterResult) bool { return v.StorageEncrypted }).(pulumi.BoolOutput)
}

// A map of tags assigned to the resource.
func (o LookupClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupClusterResultOutput) VpcSecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.VpcSecurityGroupIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}