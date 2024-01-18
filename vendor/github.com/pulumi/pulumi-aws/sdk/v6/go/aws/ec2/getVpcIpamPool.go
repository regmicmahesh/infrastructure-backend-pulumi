// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ec2.VpcIpamPool` provides details about an IPAM pool.
//
// This resource can prove useful when an ipam pool was created in another root
// module and you need the pool's id as an input variable. For example, pools
// can be shared via RAM and used to create vpcs with CIDRs from that pool.
//
// ## Example Usage
//
// The following example shows an account that has only 1 pool, perhaps shared
// via RAM, and using that pool id to create a VPC with a CIDR derived from
// AWS IPAM.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testVpcIpamPool, err := ec2.LookupVpcIpamPool(ctx, &ec2.LookupVpcIpamPoolArgs{
//				Filters: []ec2.GetVpcIpamPoolFilter{
//					{
//						Name: "description",
//						Values: []string{
//							"*test*",
//						},
//					},
//					{
//						Name: "address-family",
//						Values: []string{
//							"ipv4",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewVpc(ctx, "testVpc", &ec2.VpcArgs{
//				Ipv4IpamPoolId:    *pulumi.String(testVpcIpamPool.Id),
//				Ipv4NetmaskLength: pulumi.Int(28),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVpcIpamPool(ctx *pulumi.Context, args *LookupVpcIpamPoolArgs, opts ...pulumi.InvokeOption) (*LookupVpcIpamPoolResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcIpamPoolResult
	err := ctx.Invoke("aws:ec2/getVpcIpamPool:getVpcIpamPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcIpamPool.
type LookupVpcIpamPoolArgs struct {
	// Tags that are required to create resources in using this pool.
	AllocationResourceTags map[string]string `pulumi:"allocationResourceTags"`
	// Custom filter block as described below.
	Filters []GetVpcIpamPoolFilter `pulumi:"filters"`
	// ID of the IPAM pool.
	Id *string `pulumi:"id"`
	// ID of the IPAM pool you would like information on.
	IpamPoolId *string `pulumi:"ipamPoolId"`
	// Map of tags to assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getVpcIpamPool.
type LookupVpcIpamPoolResult struct {
	// IP protocol assigned to this pool.
	AddressFamily string `pulumi:"addressFamily"`
	// A default netmask length for allocations added to this pool. If, for example, the CIDR assigned to this pool is `10.0.0.0/8` and you enter 16 here, new allocations will default to `10.0.0.0/16`.
	AllocationDefaultNetmaskLength int `pulumi:"allocationDefaultNetmaskLength"`
	// The maximum netmask length that will be required for CIDR allocations in this pool.
	AllocationMaxNetmaskLength int `pulumi:"allocationMaxNetmaskLength"`
	// The minimum netmask length that will be required for CIDR allocations in this pool.
	AllocationMinNetmaskLength int `pulumi:"allocationMinNetmaskLength"`
	// Tags that are required to create resources in using this pool.
	AllocationResourceTags map[string]string `pulumi:"allocationResourceTags"`
	// ARN of the pool
	Arn string `pulumi:"arn"`
	// If enabled, IPAM will continuously look for resources within the CIDR range of this pool and automatically import them as allocations into your IPAM.
	AutoImport bool `pulumi:"autoImport"`
	// Limits which service in AWS that the pool can be used in. `ec2` for example, allows users to use space for Elastic IP addresses and VPCs.
	AwsService string `pulumi:"awsService"`
	// Description for the IPAM pool.
	Description string                 `pulumi:"description"`
	Filters     []GetVpcIpamPoolFilter `pulumi:"filters"`
	// ID of the IPAM pool.
	Id         *string `pulumi:"id"`
	IpamPoolId *string `pulumi:"ipamPoolId"`
	// ID of the scope the pool belongs to.
	IpamScopeId   string `pulumi:"ipamScopeId"`
	IpamScopeType string `pulumi:"ipamScopeType"`
	// Locale is the Region where your pool is available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region.
	Locale    string `pulumi:"locale"`
	PoolDepth int    `pulumi:"poolDepth"`
	// Defines whether or not IPv6 pool space is publicly advertisable over the internet.
	PubliclyAdvertisable bool `pulumi:"publiclyAdvertisable"`
	// ID of the source IPAM pool.
	SourceIpamPoolId string `pulumi:"sourceIpamPoolId"`
	State            string `pulumi:"state"`
	// Map of tags to assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}

func LookupVpcIpamPoolOutput(ctx *pulumi.Context, args LookupVpcIpamPoolOutputArgs, opts ...pulumi.InvokeOption) LookupVpcIpamPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcIpamPoolResult, error) {
			args := v.(LookupVpcIpamPoolArgs)
			r, err := LookupVpcIpamPool(ctx, &args, opts...)
			var s LookupVpcIpamPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpcIpamPoolResultOutput)
}

// A collection of arguments for invoking getVpcIpamPool.
type LookupVpcIpamPoolOutputArgs struct {
	// Tags that are required to create resources in using this pool.
	AllocationResourceTags pulumi.StringMapInput `pulumi:"allocationResourceTags"`
	// Custom filter block as described below.
	Filters GetVpcIpamPoolFilterArrayInput `pulumi:"filters"`
	// ID of the IPAM pool.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// ID of the IPAM pool you would like information on.
	IpamPoolId pulumi.StringPtrInput `pulumi:"ipamPoolId"`
	// Map of tags to assigned to the resource.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupVpcIpamPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcIpamPoolArgs)(nil)).Elem()
}

// A collection of values returned by getVpcIpamPool.
type LookupVpcIpamPoolResultOutput struct{ *pulumi.OutputState }

func (LookupVpcIpamPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcIpamPoolResult)(nil)).Elem()
}

func (o LookupVpcIpamPoolResultOutput) ToLookupVpcIpamPoolResultOutput() LookupVpcIpamPoolResultOutput {
	return o
}

func (o LookupVpcIpamPoolResultOutput) ToLookupVpcIpamPoolResultOutputWithContext(ctx context.Context) LookupVpcIpamPoolResultOutput {
	return o
}

// IP protocol assigned to this pool.
func (o LookupVpcIpamPoolResultOutput) AddressFamily() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) string { return v.AddressFamily }).(pulumi.StringOutput)
}

// A default netmask length for allocations added to this pool. If, for example, the CIDR assigned to this pool is `10.0.0.0/8` and you enter 16 here, new allocations will default to `10.0.0.0/16`.
func (o LookupVpcIpamPoolResultOutput) AllocationDefaultNetmaskLength() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) int { return v.AllocationDefaultNetmaskLength }).(pulumi.IntOutput)
}

// The maximum netmask length that will be required for CIDR allocations in this pool.
func (o LookupVpcIpamPoolResultOutput) AllocationMaxNetmaskLength() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) int { return v.AllocationMaxNetmaskLength }).(pulumi.IntOutput)
}

// The minimum netmask length that will be required for CIDR allocations in this pool.
func (o LookupVpcIpamPoolResultOutput) AllocationMinNetmaskLength() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) int { return v.AllocationMinNetmaskLength }).(pulumi.IntOutput)
}

// Tags that are required to create resources in using this pool.
func (o LookupVpcIpamPoolResultOutput) AllocationResourceTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) map[string]string { return v.AllocationResourceTags }).(pulumi.StringMapOutput)
}

// ARN of the pool
func (o LookupVpcIpamPoolResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) string { return v.Arn }).(pulumi.StringOutput)
}

// If enabled, IPAM will continuously look for resources within the CIDR range of this pool and automatically import them as allocations into your IPAM.
func (o LookupVpcIpamPoolResultOutput) AutoImport() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) bool { return v.AutoImport }).(pulumi.BoolOutput)
}

// Limits which service in AWS that the pool can be used in. `ec2` for example, allows users to use space for Elastic IP addresses and VPCs.
func (o LookupVpcIpamPoolResultOutput) AwsService() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) string { return v.AwsService }).(pulumi.StringOutput)
}

// Description for the IPAM pool.
func (o LookupVpcIpamPoolResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupVpcIpamPoolResultOutput) Filters() GetVpcIpamPoolFilterArrayOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) []GetVpcIpamPoolFilter { return v.Filters }).(GetVpcIpamPoolFilterArrayOutput)
}

// ID of the IPAM pool.
func (o LookupVpcIpamPoolResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVpcIpamPoolResultOutput) IpamPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) *string { return v.IpamPoolId }).(pulumi.StringPtrOutput)
}

// ID of the scope the pool belongs to.
func (o LookupVpcIpamPoolResultOutput) IpamScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) string { return v.IpamScopeId }).(pulumi.StringOutput)
}

func (o LookupVpcIpamPoolResultOutput) IpamScopeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) string { return v.IpamScopeType }).(pulumi.StringOutput)
}

// Locale is the Region where your pool is available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region.
func (o LookupVpcIpamPoolResultOutput) Locale() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) string { return v.Locale }).(pulumi.StringOutput)
}

func (o LookupVpcIpamPoolResultOutput) PoolDepth() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) int { return v.PoolDepth }).(pulumi.IntOutput)
}

// Defines whether or not IPv6 pool space is publicly advertisable over the internet.
func (o LookupVpcIpamPoolResultOutput) PubliclyAdvertisable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) bool { return v.PubliclyAdvertisable }).(pulumi.BoolOutput)
}

// ID of the source IPAM pool.
func (o LookupVpcIpamPoolResultOutput) SourceIpamPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) string { return v.SourceIpamPoolId }).(pulumi.StringOutput)
}

func (o LookupVpcIpamPoolResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) string { return v.State }).(pulumi.StringOutput)
}

// Map of tags to assigned to the resource.
func (o LookupVpcIpamPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpcIpamPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcIpamPoolResultOutput{})
}
