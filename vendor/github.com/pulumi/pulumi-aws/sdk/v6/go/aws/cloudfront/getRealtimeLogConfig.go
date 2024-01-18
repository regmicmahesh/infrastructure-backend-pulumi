// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CloudFront real-time log configuration resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudfront.LookupRealtimeLogConfig(ctx, &cloudfront.LookupRealtimeLogConfigArgs{
//				Name: "example",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRealtimeLogConfig(ctx *pulumi.Context, args *LookupRealtimeLogConfigArgs, opts ...pulumi.InvokeOption) (*LookupRealtimeLogConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRealtimeLogConfigResult
	err := ctx.Invoke("aws:cloudfront/getRealtimeLogConfig:getRealtimeLogConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRealtimeLogConfig.
type LookupRealtimeLogConfigArgs struct {
	// Unique name to identify this real-time log configuration.
	Name string `pulumi:"name"`
}

// A collection of values returned by getRealtimeLogConfig.
type LookupRealtimeLogConfigResult struct {
	// ARN (Amazon Resource Name) of the CloudFront real-time log configuration.
	Arn string `pulumi:"arn"`
	// (Required) Amazon Kinesis data streams where real-time log data is sent.
	Endpoints []GetRealtimeLogConfigEndpoint `pulumi:"endpoints"`
	// (Required) Fields that are included in each real-time log record. See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) for supported values.
	Fields []string `pulumi:"fields"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// (Required) Sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between `1` and `100`, inclusive.
	SamplingRate int `pulumi:"samplingRate"`
}

func LookupRealtimeLogConfigOutput(ctx *pulumi.Context, args LookupRealtimeLogConfigOutputArgs, opts ...pulumi.InvokeOption) LookupRealtimeLogConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRealtimeLogConfigResult, error) {
			args := v.(LookupRealtimeLogConfigArgs)
			r, err := LookupRealtimeLogConfig(ctx, &args, opts...)
			var s LookupRealtimeLogConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRealtimeLogConfigResultOutput)
}

// A collection of arguments for invoking getRealtimeLogConfig.
type LookupRealtimeLogConfigOutputArgs struct {
	// Unique name to identify this real-time log configuration.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupRealtimeLogConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRealtimeLogConfigArgs)(nil)).Elem()
}

// A collection of values returned by getRealtimeLogConfig.
type LookupRealtimeLogConfigResultOutput struct{ *pulumi.OutputState }

func (LookupRealtimeLogConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRealtimeLogConfigResult)(nil)).Elem()
}

func (o LookupRealtimeLogConfigResultOutput) ToLookupRealtimeLogConfigResultOutput() LookupRealtimeLogConfigResultOutput {
	return o
}

func (o LookupRealtimeLogConfigResultOutput) ToLookupRealtimeLogConfigResultOutputWithContext(ctx context.Context) LookupRealtimeLogConfigResultOutput {
	return o
}

// ARN (Amazon Resource Name) of the CloudFront real-time log configuration.
func (o LookupRealtimeLogConfigResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealtimeLogConfigResult) string { return v.Arn }).(pulumi.StringOutput)
}

// (Required) Amazon Kinesis data streams where real-time log data is sent.
func (o LookupRealtimeLogConfigResultOutput) Endpoints() GetRealtimeLogConfigEndpointArrayOutput {
	return o.ApplyT(func(v LookupRealtimeLogConfigResult) []GetRealtimeLogConfigEndpoint { return v.Endpoints }).(GetRealtimeLogConfigEndpointArrayOutput)
}

// (Required) Fields that are included in each real-time log record. See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) for supported values.
func (o LookupRealtimeLogConfigResultOutput) Fields() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRealtimeLogConfigResult) []string { return v.Fields }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRealtimeLogConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealtimeLogConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRealtimeLogConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealtimeLogConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

// (Required) Sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between `1` and `100`, inclusive.
func (o LookupRealtimeLogConfigResultOutput) SamplingRate() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRealtimeLogConfigResult) int { return v.SamplingRate }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRealtimeLogConfigResultOutput{})
}
