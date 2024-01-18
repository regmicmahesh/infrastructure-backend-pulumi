// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **Note:** `alb.TargetGroup` is known as `lb.TargetGroup`. The functionality is identical.
//
// Provides information about a Load Balancer Target Group.
//
// This data source can prove useful when a module accepts an LB Target Group as an
// input variable and needs to know its attributes. It can also be used to get the ARN of
// an LB Target Group for use in other resources, given LB Target Group name.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			lbTgArn := ""
//			if param := cfg.Get("lbTgArn"); param != "" {
//				lbTgArn = param
//			}
//			lbTgName := ""
//			if param := cfg.Get("lbTgName"); param != "" {
//				lbTgName = param
//			}
//			_, err := lb.LookupTargetGroup(ctx, &lb.LookupTargetGroupArgs{
//				Arn:  pulumi.StringRef(lbTgArn),
//				Name: pulumi.StringRef(lbTgName),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupTargetGroup(ctx *pulumi.Context, args *LookupTargetGroupArgs, opts ...pulumi.InvokeOption) (*LookupTargetGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTargetGroupResult
	err := ctx.Invoke("aws:lb/getTargetGroup:getTargetGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTargetGroup.
type LookupTargetGroupArgs struct {
	// Full ARN of the target group.
	Arn                            *string `pulumi:"arn"`
	LoadBalancingAnomalyMitigation *string `pulumi:"loadBalancingAnomalyMitigation"`
	// Unique name of the target group.
	Name *string `pulumi:"name"`
	// Mapping of tags, each pair of which must exactly match a pair on the desired target group.
	//
	// > **NOTE:** When both `arn` and `name` are specified, `arn` takes precedence. `tags` has the lowest precedence.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getTargetGroup.
type LookupTargetGroupResult struct {
	Arn                   string                    `pulumi:"arn"`
	ArnSuffix             string                    `pulumi:"arnSuffix"`
	ConnectionTermination bool                      `pulumi:"connectionTermination"`
	DeregistrationDelay   string                    `pulumi:"deregistrationDelay"`
	HealthCheck           GetTargetGroupHealthCheck `pulumi:"healthCheck"`
	// The provider-assigned unique ID for this managed resource.
	Id                             string                   `pulumi:"id"`
	LambdaMultiValueHeadersEnabled bool                     `pulumi:"lambdaMultiValueHeadersEnabled"`
	LoadBalancingAlgorithmType     string                   `pulumi:"loadBalancingAlgorithmType"`
	LoadBalancingAnomalyMitigation string                   `pulumi:"loadBalancingAnomalyMitigation"`
	LoadBalancingCrossZoneEnabled  string                   `pulumi:"loadBalancingCrossZoneEnabled"`
	Name                           string                   `pulumi:"name"`
	Port                           int                      `pulumi:"port"`
	PreserveClientIp               string                   `pulumi:"preserveClientIp"`
	Protocol                       string                   `pulumi:"protocol"`
	ProtocolVersion                string                   `pulumi:"protocolVersion"`
	ProxyProtocolV2                bool                     `pulumi:"proxyProtocolV2"`
	SlowStart                      int                      `pulumi:"slowStart"`
	Stickiness                     GetTargetGroupStickiness `pulumi:"stickiness"`
	Tags                           map[string]string        `pulumi:"tags"`
	TargetType                     string                   `pulumi:"targetType"`
	VpcId                          string                   `pulumi:"vpcId"`
}

func LookupTargetGroupOutput(ctx *pulumi.Context, args LookupTargetGroupOutputArgs, opts ...pulumi.InvokeOption) LookupTargetGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTargetGroupResult, error) {
			args := v.(LookupTargetGroupArgs)
			r, err := LookupTargetGroup(ctx, &args, opts...)
			var s LookupTargetGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTargetGroupResultOutput)
}

// A collection of arguments for invoking getTargetGroup.
type LookupTargetGroupOutputArgs struct {
	// Full ARN of the target group.
	Arn                            pulumi.StringPtrInput `pulumi:"arn"`
	LoadBalancingAnomalyMitigation pulumi.StringPtrInput `pulumi:"loadBalancingAnomalyMitigation"`
	// Unique name of the target group.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Mapping of tags, each pair of which must exactly match a pair on the desired target group.
	//
	// > **NOTE:** When both `arn` and `name` are specified, `arn` takes precedence. `tags` has the lowest precedence.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupTargetGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetGroupArgs)(nil)).Elem()
}

// A collection of values returned by getTargetGroup.
type LookupTargetGroupResultOutput struct{ *pulumi.OutputState }

func (LookupTargetGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetGroupResult)(nil)).Elem()
}

func (o LookupTargetGroupResultOutput) ToLookupTargetGroupResultOutput() LookupTargetGroupResultOutput {
	return o
}

func (o LookupTargetGroupResultOutput) ToLookupTargetGroupResultOutputWithContext(ctx context.Context) LookupTargetGroupResultOutput {
	return o
}

func (o LookupTargetGroupResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) string { return v.Arn }).(pulumi.StringOutput)
}

func (o LookupTargetGroupResultOutput) ArnSuffix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) string { return v.ArnSuffix }).(pulumi.StringOutput)
}

func (o LookupTargetGroupResultOutput) ConnectionTermination() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) bool { return v.ConnectionTermination }).(pulumi.BoolOutput)
}

func (o LookupTargetGroupResultOutput) DeregistrationDelay() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) string { return v.DeregistrationDelay }).(pulumi.StringOutput)
}

func (o LookupTargetGroupResultOutput) HealthCheck() GetTargetGroupHealthCheckOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) GetTargetGroupHealthCheck { return v.HealthCheck }).(GetTargetGroupHealthCheckOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupTargetGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTargetGroupResultOutput) LambdaMultiValueHeadersEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) bool { return v.LambdaMultiValueHeadersEnabled }).(pulumi.BoolOutput)
}

func (o LookupTargetGroupResultOutput) LoadBalancingAlgorithmType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) string { return v.LoadBalancingAlgorithmType }).(pulumi.StringOutput)
}

func (o LookupTargetGroupResultOutput) LoadBalancingAnomalyMitigation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) string { return v.LoadBalancingAnomalyMitigation }).(pulumi.StringOutput)
}

func (o LookupTargetGroupResultOutput) LoadBalancingCrossZoneEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) string { return v.LoadBalancingCrossZoneEnabled }).(pulumi.StringOutput)
}

func (o LookupTargetGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTargetGroupResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) int { return v.Port }).(pulumi.IntOutput)
}

func (o LookupTargetGroupResultOutput) PreserveClientIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) string { return v.PreserveClientIp }).(pulumi.StringOutput)
}

func (o LookupTargetGroupResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o LookupTargetGroupResultOutput) ProtocolVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) string { return v.ProtocolVersion }).(pulumi.StringOutput)
}

func (o LookupTargetGroupResultOutput) ProxyProtocolV2() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) bool { return v.ProxyProtocolV2 }).(pulumi.BoolOutput)
}

func (o LookupTargetGroupResultOutput) SlowStart() pulumi.IntOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) int { return v.SlowStart }).(pulumi.IntOutput)
}

func (o LookupTargetGroupResultOutput) Stickiness() GetTargetGroupStickinessOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) GetTargetGroupStickiness { return v.Stickiness }).(GetTargetGroupStickinessOutput)
}

func (o LookupTargetGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupTargetGroupResultOutput) TargetType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) string { return v.TargetType }).(pulumi.StringOutput)
}

func (o LookupTargetGroupResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTargetGroupResultOutput{})
}
