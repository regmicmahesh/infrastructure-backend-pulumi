// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Amazon S3 Express directory bucket resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := s3.NewDirectoryBucket(ctx, "example", &s3.DirectoryBucketArgs{
//				Bucket: pulumi.String("example--usw2-az1--x-s3"),
//				Location: &s3.DirectoryBucketLocationArgs{
//					Name: pulumi.String("usw2-az1"),
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
// Using `pulumi import`, import S3 bucket using `bucket`. For example:
//
// ```sh
//
//	$ pulumi import aws:s3/directoryBucket:DirectoryBucket example example--usw2-az1--x-s3
//
// ```
type DirectoryBucket struct {
	pulumi.CustomResourceState

	// ARN of the bucket.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Name of the bucket. The name must be in the format `[bucketName]--[azid]--x-s3`. Use the `s3.BucketV2` resource to manage general purpose buckets.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Data redundancy. Valid values: `SingleAvailabilityZone`.
	DataRedundancy pulumi.StringOutput `pulumi:"dataRedundancy"`
	// Boolean that indicates all objects should be deleted from the bucket *when the bucket is destroyed* so that the bucket can be destroyed without error. These objects are *not* recoverable. This only deletes objects when the bucket is destroyed, *not* when setting this parameter to `true`. Once this parameter is set to `true`, there must be a successful `pulumi up` run before a destroy is required to update this value in the resource state. Without a successful `pulumi up` after this parameter is set, this flag will have no effect. If setting this field in the same operation that would require replacing the bucket or destroying the bucket, this flag will not work. Additionally when importing a bucket, a successful `pulumi up` is required to set this value in state before it will take effect on a destroy operation.
	ForceDestroy pulumi.BoolOutput `pulumi:"forceDestroy"`
	// Bucket location. See Location below for more details.
	Location DirectoryBucketLocationPtrOutput `pulumi:"location"`
	// Bucket type. Valid values: `Directory`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDirectoryBucket registers a new resource with the given unique name, arguments, and options.
func NewDirectoryBucket(ctx *pulumi.Context,
	name string, args *DirectoryBucketArgs, opts ...pulumi.ResourceOption) (*DirectoryBucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DirectoryBucket
	err := ctx.RegisterResource("aws:s3/directoryBucket:DirectoryBucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDirectoryBucket gets an existing DirectoryBucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDirectoryBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DirectoryBucketState, opts ...pulumi.ResourceOption) (*DirectoryBucket, error) {
	var resource DirectoryBucket
	err := ctx.ReadResource("aws:s3/directoryBucket:DirectoryBucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DirectoryBucket resources.
type directoryBucketState struct {
	// ARN of the bucket.
	Arn *string `pulumi:"arn"`
	// Name of the bucket. The name must be in the format `[bucketName]--[azid]--x-s3`. Use the `s3.BucketV2` resource to manage general purpose buckets.
	Bucket *string `pulumi:"bucket"`
	// Data redundancy. Valid values: `SingleAvailabilityZone`.
	DataRedundancy *string `pulumi:"dataRedundancy"`
	// Boolean that indicates all objects should be deleted from the bucket *when the bucket is destroyed* so that the bucket can be destroyed without error. These objects are *not* recoverable. This only deletes objects when the bucket is destroyed, *not* when setting this parameter to `true`. Once this parameter is set to `true`, there must be a successful `pulumi up` run before a destroy is required to update this value in the resource state. Without a successful `pulumi up` after this parameter is set, this flag will have no effect. If setting this field in the same operation that would require replacing the bucket or destroying the bucket, this flag will not work. Additionally when importing a bucket, a successful `pulumi up` is required to set this value in state before it will take effect on a destroy operation.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// Bucket location. See Location below for more details.
	Location *DirectoryBucketLocation `pulumi:"location"`
	// Bucket type. Valid values: `Directory`.
	Type *string `pulumi:"type"`
}

type DirectoryBucketState struct {
	// ARN of the bucket.
	Arn pulumi.StringPtrInput
	// Name of the bucket. The name must be in the format `[bucketName]--[azid]--x-s3`. Use the `s3.BucketV2` resource to manage general purpose buckets.
	Bucket pulumi.StringPtrInput
	// Data redundancy. Valid values: `SingleAvailabilityZone`.
	DataRedundancy pulumi.StringPtrInput
	// Boolean that indicates all objects should be deleted from the bucket *when the bucket is destroyed* so that the bucket can be destroyed without error. These objects are *not* recoverable. This only deletes objects when the bucket is destroyed, *not* when setting this parameter to `true`. Once this parameter is set to `true`, there must be a successful `pulumi up` run before a destroy is required to update this value in the resource state. Without a successful `pulumi up` after this parameter is set, this flag will have no effect. If setting this field in the same operation that would require replacing the bucket or destroying the bucket, this flag will not work. Additionally when importing a bucket, a successful `pulumi up` is required to set this value in state before it will take effect on a destroy operation.
	ForceDestroy pulumi.BoolPtrInput
	// Bucket location. See Location below for more details.
	Location DirectoryBucketLocationPtrInput
	// Bucket type. Valid values: `Directory`.
	Type pulumi.StringPtrInput
}

func (DirectoryBucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*directoryBucketState)(nil)).Elem()
}

type directoryBucketArgs struct {
	// Name of the bucket. The name must be in the format `[bucketName]--[azid]--x-s3`. Use the `s3.BucketV2` resource to manage general purpose buckets.
	Bucket string `pulumi:"bucket"`
	// Data redundancy. Valid values: `SingleAvailabilityZone`.
	DataRedundancy *string `pulumi:"dataRedundancy"`
	// Boolean that indicates all objects should be deleted from the bucket *when the bucket is destroyed* so that the bucket can be destroyed without error. These objects are *not* recoverable. This only deletes objects when the bucket is destroyed, *not* when setting this parameter to `true`. Once this parameter is set to `true`, there must be a successful `pulumi up` run before a destroy is required to update this value in the resource state. Without a successful `pulumi up` after this parameter is set, this flag will have no effect. If setting this field in the same operation that would require replacing the bucket or destroying the bucket, this flag will not work. Additionally when importing a bucket, a successful `pulumi up` is required to set this value in state before it will take effect on a destroy operation.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// Bucket location. See Location below for more details.
	Location *DirectoryBucketLocation `pulumi:"location"`
	// Bucket type. Valid values: `Directory`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a DirectoryBucket resource.
type DirectoryBucketArgs struct {
	// Name of the bucket. The name must be in the format `[bucketName]--[azid]--x-s3`. Use the `s3.BucketV2` resource to manage general purpose buckets.
	Bucket pulumi.StringInput
	// Data redundancy. Valid values: `SingleAvailabilityZone`.
	DataRedundancy pulumi.StringPtrInput
	// Boolean that indicates all objects should be deleted from the bucket *when the bucket is destroyed* so that the bucket can be destroyed without error. These objects are *not* recoverable. This only deletes objects when the bucket is destroyed, *not* when setting this parameter to `true`. Once this parameter is set to `true`, there must be a successful `pulumi up` run before a destroy is required to update this value in the resource state. Without a successful `pulumi up` after this parameter is set, this flag will have no effect. If setting this field in the same operation that would require replacing the bucket or destroying the bucket, this flag will not work. Additionally when importing a bucket, a successful `pulumi up` is required to set this value in state before it will take effect on a destroy operation.
	ForceDestroy pulumi.BoolPtrInput
	// Bucket location. See Location below for more details.
	Location DirectoryBucketLocationPtrInput
	// Bucket type. Valid values: `Directory`.
	Type pulumi.StringPtrInput
}

func (DirectoryBucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*directoryBucketArgs)(nil)).Elem()
}

type DirectoryBucketInput interface {
	pulumi.Input

	ToDirectoryBucketOutput() DirectoryBucketOutput
	ToDirectoryBucketOutputWithContext(ctx context.Context) DirectoryBucketOutput
}

func (*DirectoryBucket) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectoryBucket)(nil)).Elem()
}

func (i *DirectoryBucket) ToDirectoryBucketOutput() DirectoryBucketOutput {
	return i.ToDirectoryBucketOutputWithContext(context.Background())
}

func (i *DirectoryBucket) ToDirectoryBucketOutputWithContext(ctx context.Context) DirectoryBucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryBucketOutput)
}

// DirectoryBucketArrayInput is an input type that accepts DirectoryBucketArray and DirectoryBucketArrayOutput values.
// You can construct a concrete instance of `DirectoryBucketArrayInput` via:
//
//	DirectoryBucketArray{ DirectoryBucketArgs{...} }
type DirectoryBucketArrayInput interface {
	pulumi.Input

	ToDirectoryBucketArrayOutput() DirectoryBucketArrayOutput
	ToDirectoryBucketArrayOutputWithContext(context.Context) DirectoryBucketArrayOutput
}

type DirectoryBucketArray []DirectoryBucketInput

func (DirectoryBucketArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DirectoryBucket)(nil)).Elem()
}

func (i DirectoryBucketArray) ToDirectoryBucketArrayOutput() DirectoryBucketArrayOutput {
	return i.ToDirectoryBucketArrayOutputWithContext(context.Background())
}

func (i DirectoryBucketArray) ToDirectoryBucketArrayOutputWithContext(ctx context.Context) DirectoryBucketArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryBucketArrayOutput)
}

// DirectoryBucketMapInput is an input type that accepts DirectoryBucketMap and DirectoryBucketMapOutput values.
// You can construct a concrete instance of `DirectoryBucketMapInput` via:
//
//	DirectoryBucketMap{ "key": DirectoryBucketArgs{...} }
type DirectoryBucketMapInput interface {
	pulumi.Input

	ToDirectoryBucketMapOutput() DirectoryBucketMapOutput
	ToDirectoryBucketMapOutputWithContext(context.Context) DirectoryBucketMapOutput
}

type DirectoryBucketMap map[string]DirectoryBucketInput

func (DirectoryBucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DirectoryBucket)(nil)).Elem()
}

func (i DirectoryBucketMap) ToDirectoryBucketMapOutput() DirectoryBucketMapOutput {
	return i.ToDirectoryBucketMapOutputWithContext(context.Background())
}

func (i DirectoryBucketMap) ToDirectoryBucketMapOutputWithContext(ctx context.Context) DirectoryBucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryBucketMapOutput)
}

type DirectoryBucketOutput struct{ *pulumi.OutputState }

func (DirectoryBucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectoryBucket)(nil)).Elem()
}

func (o DirectoryBucketOutput) ToDirectoryBucketOutput() DirectoryBucketOutput {
	return o
}

func (o DirectoryBucketOutput) ToDirectoryBucketOutputWithContext(ctx context.Context) DirectoryBucketOutput {
	return o
}

// ARN of the bucket.
func (o DirectoryBucketOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectoryBucket) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Name of the bucket. The name must be in the format `[bucketName]--[azid]--x-s3`. Use the `s3.BucketV2` resource to manage general purpose buckets.
func (o DirectoryBucketOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectoryBucket) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// Data redundancy. Valid values: `SingleAvailabilityZone`.
func (o DirectoryBucketOutput) DataRedundancy() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectoryBucket) pulumi.StringOutput { return v.DataRedundancy }).(pulumi.StringOutput)
}

// Boolean that indicates all objects should be deleted from the bucket *when the bucket is destroyed* so that the bucket can be destroyed without error. These objects are *not* recoverable. This only deletes objects when the bucket is destroyed, *not* when setting this parameter to `true`. Once this parameter is set to `true`, there must be a successful `pulumi up` run before a destroy is required to update this value in the resource state. Without a successful `pulumi up` after this parameter is set, this flag will have no effect. If setting this field in the same operation that would require replacing the bucket or destroying the bucket, this flag will not work. Additionally when importing a bucket, a successful `pulumi up` is required to set this value in state before it will take effect on a destroy operation.
func (o DirectoryBucketOutput) ForceDestroy() pulumi.BoolOutput {
	return o.ApplyT(func(v *DirectoryBucket) pulumi.BoolOutput { return v.ForceDestroy }).(pulumi.BoolOutput)
}

// Bucket location. See Location below for more details.
func (o DirectoryBucketOutput) Location() DirectoryBucketLocationPtrOutput {
	return o.ApplyT(func(v *DirectoryBucket) DirectoryBucketLocationPtrOutput { return v.Location }).(DirectoryBucketLocationPtrOutput)
}

// Bucket type. Valid values: `Directory`.
func (o DirectoryBucketOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectoryBucket) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type DirectoryBucketArrayOutput struct{ *pulumi.OutputState }

func (DirectoryBucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DirectoryBucket)(nil)).Elem()
}

func (o DirectoryBucketArrayOutput) ToDirectoryBucketArrayOutput() DirectoryBucketArrayOutput {
	return o
}

func (o DirectoryBucketArrayOutput) ToDirectoryBucketArrayOutputWithContext(ctx context.Context) DirectoryBucketArrayOutput {
	return o
}

func (o DirectoryBucketArrayOutput) Index(i pulumi.IntInput) DirectoryBucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DirectoryBucket {
		return vs[0].([]*DirectoryBucket)[vs[1].(int)]
	}).(DirectoryBucketOutput)
}

type DirectoryBucketMapOutput struct{ *pulumi.OutputState }

func (DirectoryBucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DirectoryBucket)(nil)).Elem()
}

func (o DirectoryBucketMapOutput) ToDirectoryBucketMapOutput() DirectoryBucketMapOutput {
	return o
}

func (o DirectoryBucketMapOutput) ToDirectoryBucketMapOutputWithContext(ctx context.Context) DirectoryBucketMapOutput {
	return o
}

func (o DirectoryBucketMapOutput) MapIndex(k pulumi.StringInput) DirectoryBucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DirectoryBucket {
		return vs[0].(map[string]*DirectoryBucket)[vs[1].(string)]
	}).(DirectoryBucketOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DirectoryBucketInput)(nil)).Elem(), &DirectoryBucket{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectoryBucketArrayInput)(nil)).Elem(), DirectoryBucketArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectoryBucketMapInput)(nil)).Elem(), DirectoryBucketMap{})
	pulumi.RegisterOutputType(DirectoryBucketOutput{})
	pulumi.RegisterOutputType(DirectoryBucketArrayOutput{})
	pulumi.RegisterOutputType(DirectoryBucketMapOutput{})
}