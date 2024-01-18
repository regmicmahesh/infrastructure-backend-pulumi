// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a CloudFront cache policy.
//
// ## Example Usage
// ### Basic Usage
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
//			_, err := cloudfront.LookupResponseHeadersPolicy(ctx, &cloudfront.LookupResponseHeadersPolicyArgs{
//				Name: pulumi.StringRef("example-policy"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### AWS-Managed Policies
//
// AWS managed response header policy names are prefixed with `Managed-`:
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
//			_, err := cloudfront.LookupResponseHeadersPolicy(ctx, &cloudfront.LookupResponseHeadersPolicyArgs{
//				Name: pulumi.StringRef("Managed-SimpleCORS"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupResponseHeadersPolicy(ctx *pulumi.Context, args *LookupResponseHeadersPolicyArgs, opts ...pulumi.InvokeOption) (*LookupResponseHeadersPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupResponseHeadersPolicyResult
	err := ctx.Invoke("aws:cloudfront/getResponseHeadersPolicy:getResponseHeadersPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResponseHeadersPolicy.
type LookupResponseHeadersPolicyArgs struct {
	// Identifier for the response headers policy.
	Id *string `pulumi:"id"`
	// Unique name to identify the response headers policy.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getResponseHeadersPolicy.
type LookupResponseHeadersPolicyResult struct {
	// Comment to describe the response headers policy. The comment cannot be longer than 128 characters.
	Comment string `pulumi:"comment"`
	// Configuration for a set of HTTP response headers that are used for Cross-Origin Resource Sharing (CORS). See Cors Config for more information.
	CorsConfigs []GetResponseHeadersPolicyCorsConfig `pulumi:"corsConfigs"`
	// Object that contains an attribute `items` that contains a list of Custom Headers. See Custom Header for more information.
	CustomHeadersConfigs []GetResponseHeadersPolicyCustomHeadersConfig `pulumi:"customHeadersConfigs"`
	// Current version of the response headers policy.
	Etag string `pulumi:"etag"`
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Object that contains an attribute `items` that contains a list of Remove Headers. See Remove Header for more information.
	RemoveHeadersConfigs []GetResponseHeadersPolicyRemoveHeadersConfig `pulumi:"removeHeadersConfigs"`
	// A configuration for a set of security-related HTTP response headers. See Security Headers Config for more information.
	SecurityHeadersConfigs []GetResponseHeadersPolicySecurityHeadersConfig `pulumi:"securityHeadersConfigs"`
	// (Optional) Configuration for enabling the Server-Timing header in HTTP responses sent from CloudFront. See Server Timing Headers Config for more information.
	ServerTimingHeadersConfigs []GetResponseHeadersPolicyServerTimingHeadersConfig `pulumi:"serverTimingHeadersConfigs"`
}

func LookupResponseHeadersPolicyOutput(ctx *pulumi.Context, args LookupResponseHeadersPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupResponseHeadersPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResponseHeadersPolicyResult, error) {
			args := v.(LookupResponseHeadersPolicyArgs)
			r, err := LookupResponseHeadersPolicy(ctx, &args, opts...)
			var s LookupResponseHeadersPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResponseHeadersPolicyResultOutput)
}

// A collection of arguments for invoking getResponseHeadersPolicy.
type LookupResponseHeadersPolicyOutputArgs struct {
	// Identifier for the response headers policy.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Unique name to identify the response headers policy.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupResponseHeadersPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResponseHeadersPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getResponseHeadersPolicy.
type LookupResponseHeadersPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupResponseHeadersPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResponseHeadersPolicyResult)(nil)).Elem()
}

func (o LookupResponseHeadersPolicyResultOutput) ToLookupResponseHeadersPolicyResultOutput() LookupResponseHeadersPolicyResultOutput {
	return o
}

func (o LookupResponseHeadersPolicyResultOutput) ToLookupResponseHeadersPolicyResultOutputWithContext(ctx context.Context) LookupResponseHeadersPolicyResultOutput {
	return o
}

// Comment to describe the response headers policy. The comment cannot be longer than 128 characters.
func (o LookupResponseHeadersPolicyResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResponseHeadersPolicyResult) string { return v.Comment }).(pulumi.StringOutput)
}

// Configuration for a set of HTTP response headers that are used for Cross-Origin Resource Sharing (CORS). See Cors Config for more information.
func (o LookupResponseHeadersPolicyResultOutput) CorsConfigs() GetResponseHeadersPolicyCorsConfigArrayOutput {
	return o.ApplyT(func(v LookupResponseHeadersPolicyResult) []GetResponseHeadersPolicyCorsConfig { return v.CorsConfigs }).(GetResponseHeadersPolicyCorsConfigArrayOutput)
}

// Object that contains an attribute `items` that contains a list of Custom Headers. See Custom Header for more information.
func (o LookupResponseHeadersPolicyResultOutput) CustomHeadersConfigs() GetResponseHeadersPolicyCustomHeadersConfigArrayOutput {
	return o.ApplyT(func(v LookupResponseHeadersPolicyResult) []GetResponseHeadersPolicyCustomHeadersConfig {
		return v.CustomHeadersConfigs
	}).(GetResponseHeadersPolicyCustomHeadersConfigArrayOutput)
}

// Current version of the response headers policy.
func (o LookupResponseHeadersPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResponseHeadersPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupResponseHeadersPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResponseHeadersPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupResponseHeadersPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResponseHeadersPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Object that contains an attribute `items` that contains a list of Remove Headers. See Remove Header for more information.
func (o LookupResponseHeadersPolicyResultOutput) RemoveHeadersConfigs() GetResponseHeadersPolicyRemoveHeadersConfigArrayOutput {
	return o.ApplyT(func(v LookupResponseHeadersPolicyResult) []GetResponseHeadersPolicyRemoveHeadersConfig {
		return v.RemoveHeadersConfigs
	}).(GetResponseHeadersPolicyRemoveHeadersConfigArrayOutput)
}

// A configuration for a set of security-related HTTP response headers. See Security Headers Config for more information.
func (o LookupResponseHeadersPolicyResultOutput) SecurityHeadersConfigs() GetResponseHeadersPolicySecurityHeadersConfigArrayOutput {
	return o.ApplyT(func(v LookupResponseHeadersPolicyResult) []GetResponseHeadersPolicySecurityHeadersConfig {
		return v.SecurityHeadersConfigs
	}).(GetResponseHeadersPolicySecurityHeadersConfigArrayOutput)
}

// (Optional) Configuration for enabling the Server-Timing header in HTTP responses sent from CloudFront. See Server Timing Headers Config for more information.
func (o LookupResponseHeadersPolicyResultOutput) ServerTimingHeadersConfigs() GetResponseHeadersPolicyServerTimingHeadersConfigArrayOutput {
	return o.ApplyT(func(v LookupResponseHeadersPolicyResult) []GetResponseHeadersPolicyServerTimingHeadersConfig {
		return v.ServerTimingHeadersConfigs
	}).(GetResponseHeadersPolicyServerTimingHeadersConfigArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResponseHeadersPolicyResultOutput{})
}
