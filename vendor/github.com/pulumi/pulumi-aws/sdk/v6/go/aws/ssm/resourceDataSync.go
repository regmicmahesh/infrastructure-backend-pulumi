// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a SSM resource data sync.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			hogeBucketV2, err := s3.NewBucketV2(ctx, "hogeBucketV2", nil)
//			if err != nil {
//				return err
//			}
//			hogePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Sid:    pulumi.StringRef("SSMBucketPermissionsCheck"),
//						Effect: pulumi.StringRef("Allow"),
//						Principals: []iam.GetPolicyDocumentStatementPrincipal{
//							{
//								Type: "Service",
//								Identifiers: []string{
//									"ssm.amazonaws.com",
//								},
//							},
//						},
//						Actions: []string{
//							"s3:GetBucketAcl",
//						},
//						Resources: []string{
//							"arn:aws:s3:::tf-test-bucket-1234",
//						},
//					},
//					{
//						Sid:    pulumi.StringRef("SSMBucketDelivery"),
//						Effect: pulumi.StringRef("Allow"),
//						Principals: []iam.GetPolicyDocumentStatementPrincipal{
//							{
//								Type: "Service",
//								Identifiers: []string{
//									"ssm.amazonaws.com",
//								},
//							},
//						},
//						Actions: []string{
//							"s3:PutObject",
//						},
//						Resources: []string{
//							"arn:aws:s3:::tf-test-bucket-1234/*",
//						},
//						Conditions: []iam.GetPolicyDocumentStatementCondition{
//							{
//								Test:     "StringEquals",
//								Variable: "s3:x-amz-acl",
//								Values: []string{
//									"bucket-owner-full-control",
//								},
//							},
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketPolicy(ctx, "hogeBucketPolicy", &s3.BucketPolicyArgs{
//				Bucket: hogeBucketV2.ID(),
//				Policy: *pulumi.String(hogePolicyDocument.Json),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ssm.NewResourceDataSync(ctx, "foo", &ssm.ResourceDataSyncArgs{
//				S3Destination: &ssm.ResourceDataSyncS3DestinationArgs{
//					BucketName: hogeBucketV2.Bucket,
//					Region:     hogeBucketV2.Region,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Using `pulumi import`, import SSM resource data sync using the `name`. For example:
//
// ```sh
//
//	$ pulumi import aws:ssm/resourceDataSync:ResourceDataSync example example-name
//
// ```
type ResourceDataSync struct {
	pulumi.CustomResourceState

	// Name for the configuration.
	Name pulumi.StringOutput `pulumi:"name"`
	// Amazon S3 configuration details for the sync.
	S3Destination ResourceDataSyncS3DestinationOutput `pulumi:"s3Destination"`
}

// NewResourceDataSync registers a new resource with the given unique name, arguments, and options.
func NewResourceDataSync(ctx *pulumi.Context,
	name string, args *ResourceDataSyncArgs, opts ...pulumi.ResourceOption) (*ResourceDataSync, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.S3Destination == nil {
		return nil, errors.New("invalid value for required argument 'S3Destination'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResourceDataSync
	err := ctx.RegisterResource("aws:ssm/resourceDataSync:ResourceDataSync", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceDataSync gets an existing ResourceDataSync resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceDataSync(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceDataSyncState, opts ...pulumi.ResourceOption) (*ResourceDataSync, error) {
	var resource ResourceDataSync
	err := ctx.ReadResource("aws:ssm/resourceDataSync:ResourceDataSync", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceDataSync resources.
type resourceDataSyncState struct {
	// Name for the configuration.
	Name *string `pulumi:"name"`
	// Amazon S3 configuration details for the sync.
	S3Destination *ResourceDataSyncS3Destination `pulumi:"s3Destination"`
}

type ResourceDataSyncState struct {
	// Name for the configuration.
	Name pulumi.StringPtrInput
	// Amazon S3 configuration details for the sync.
	S3Destination ResourceDataSyncS3DestinationPtrInput
}

func (ResourceDataSyncState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceDataSyncState)(nil)).Elem()
}

type resourceDataSyncArgs struct {
	// Name for the configuration.
	Name *string `pulumi:"name"`
	// Amazon S3 configuration details for the sync.
	S3Destination ResourceDataSyncS3Destination `pulumi:"s3Destination"`
}

// The set of arguments for constructing a ResourceDataSync resource.
type ResourceDataSyncArgs struct {
	// Name for the configuration.
	Name pulumi.StringPtrInput
	// Amazon S3 configuration details for the sync.
	S3Destination ResourceDataSyncS3DestinationInput
}

func (ResourceDataSyncArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceDataSyncArgs)(nil)).Elem()
}

type ResourceDataSyncInput interface {
	pulumi.Input

	ToResourceDataSyncOutput() ResourceDataSyncOutput
	ToResourceDataSyncOutputWithContext(ctx context.Context) ResourceDataSyncOutput
}

func (*ResourceDataSync) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceDataSync)(nil)).Elem()
}

func (i *ResourceDataSync) ToResourceDataSyncOutput() ResourceDataSyncOutput {
	return i.ToResourceDataSyncOutputWithContext(context.Background())
}

func (i *ResourceDataSync) ToResourceDataSyncOutputWithContext(ctx context.Context) ResourceDataSyncOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceDataSyncOutput)
}

// ResourceDataSyncArrayInput is an input type that accepts ResourceDataSyncArray and ResourceDataSyncArrayOutput values.
// You can construct a concrete instance of `ResourceDataSyncArrayInput` via:
//
//	ResourceDataSyncArray{ ResourceDataSyncArgs{...} }
type ResourceDataSyncArrayInput interface {
	pulumi.Input

	ToResourceDataSyncArrayOutput() ResourceDataSyncArrayOutput
	ToResourceDataSyncArrayOutputWithContext(context.Context) ResourceDataSyncArrayOutput
}

type ResourceDataSyncArray []ResourceDataSyncInput

func (ResourceDataSyncArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceDataSync)(nil)).Elem()
}

func (i ResourceDataSyncArray) ToResourceDataSyncArrayOutput() ResourceDataSyncArrayOutput {
	return i.ToResourceDataSyncArrayOutputWithContext(context.Background())
}

func (i ResourceDataSyncArray) ToResourceDataSyncArrayOutputWithContext(ctx context.Context) ResourceDataSyncArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceDataSyncArrayOutput)
}

// ResourceDataSyncMapInput is an input type that accepts ResourceDataSyncMap and ResourceDataSyncMapOutput values.
// You can construct a concrete instance of `ResourceDataSyncMapInput` via:
//
//	ResourceDataSyncMap{ "key": ResourceDataSyncArgs{...} }
type ResourceDataSyncMapInput interface {
	pulumi.Input

	ToResourceDataSyncMapOutput() ResourceDataSyncMapOutput
	ToResourceDataSyncMapOutputWithContext(context.Context) ResourceDataSyncMapOutput
}

type ResourceDataSyncMap map[string]ResourceDataSyncInput

func (ResourceDataSyncMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceDataSync)(nil)).Elem()
}

func (i ResourceDataSyncMap) ToResourceDataSyncMapOutput() ResourceDataSyncMapOutput {
	return i.ToResourceDataSyncMapOutputWithContext(context.Background())
}

func (i ResourceDataSyncMap) ToResourceDataSyncMapOutputWithContext(ctx context.Context) ResourceDataSyncMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceDataSyncMapOutput)
}

type ResourceDataSyncOutput struct{ *pulumi.OutputState }

func (ResourceDataSyncOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceDataSync)(nil)).Elem()
}

func (o ResourceDataSyncOutput) ToResourceDataSyncOutput() ResourceDataSyncOutput {
	return o
}

func (o ResourceDataSyncOutput) ToResourceDataSyncOutputWithContext(ctx context.Context) ResourceDataSyncOutput {
	return o
}

// Name for the configuration.
func (o ResourceDataSyncOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceDataSync) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Amazon S3 configuration details for the sync.
func (o ResourceDataSyncOutput) S3Destination() ResourceDataSyncS3DestinationOutput {
	return o.ApplyT(func(v *ResourceDataSync) ResourceDataSyncS3DestinationOutput { return v.S3Destination }).(ResourceDataSyncS3DestinationOutput)
}

type ResourceDataSyncArrayOutput struct{ *pulumi.OutputState }

func (ResourceDataSyncArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceDataSync)(nil)).Elem()
}

func (o ResourceDataSyncArrayOutput) ToResourceDataSyncArrayOutput() ResourceDataSyncArrayOutput {
	return o
}

func (o ResourceDataSyncArrayOutput) ToResourceDataSyncArrayOutputWithContext(ctx context.Context) ResourceDataSyncArrayOutput {
	return o
}

func (o ResourceDataSyncArrayOutput) Index(i pulumi.IntInput) ResourceDataSyncOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceDataSync {
		return vs[0].([]*ResourceDataSync)[vs[1].(int)]
	}).(ResourceDataSyncOutput)
}

type ResourceDataSyncMapOutput struct{ *pulumi.OutputState }

func (ResourceDataSyncMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceDataSync)(nil)).Elem()
}

func (o ResourceDataSyncMapOutput) ToResourceDataSyncMapOutput() ResourceDataSyncMapOutput {
	return o
}

func (o ResourceDataSyncMapOutput) ToResourceDataSyncMapOutputWithContext(ctx context.Context) ResourceDataSyncMapOutput {
	return o
}

func (o ResourceDataSyncMapOutput) MapIndex(k pulumi.StringInput) ResourceDataSyncOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceDataSync {
		return vs[0].(map[string]*ResourceDataSync)[vs[1].(string)]
	}).(ResourceDataSyncOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceDataSyncInput)(nil)).Elem(), &ResourceDataSync{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceDataSyncArrayInput)(nil)).Elem(), ResourceDataSyncArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceDataSyncMapInput)(nil)).Elem(), ResourceDataSyncMap{})
	pulumi.RegisterOutputType(ResourceDataSyncOutput{})
	pulumi.RegisterOutputType(ResourceDataSyncArrayOutput{})
	pulumi.RegisterOutputType(ResourceDataSyncMapOutput{})
}
