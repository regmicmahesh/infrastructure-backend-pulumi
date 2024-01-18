// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about a specific VPC NAT Gateway.
//
// ## Example Usage
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
//			_, err := ec2.LookupNatGateway(ctx, &ec2.LookupNatGatewayArgs{
//				SubnetId: pulumi.StringRef(aws_subnet.Public.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With tags
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
//			_, err := ec2.LookupNatGateway(ctx, &ec2.LookupNatGatewayArgs{
//				SubnetId: pulumi.StringRef(aws_subnet.Public.Id),
//				Tags: map[string]interface{}{
//					"Name": "gw NAT",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNatGateway(ctx *pulumi.Context, args *LookupNatGatewayArgs, opts ...pulumi.InvokeOption) (*LookupNatGatewayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNatGatewayResult
	err := ctx.Invoke("aws:ec2/getNatGateway:getNatGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNatGateway.
type LookupNatGatewayArgs struct {
	// Custom filter block as described below.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Filters []GetNatGatewayFilter `pulumi:"filters"`
	// ID of the specific NAT Gateway to retrieve.
	Id *string `pulumi:"id"`
	// State of the NAT Gateway (pending | failed | available | deleting | deleted ).
	State *string `pulumi:"state"`
	// ID of subnet that the NAT Gateway resides in.
	SubnetId *string `pulumi:"subnetId"`
	// Map of tags, each pair of which must exactly match
	// a pair on the desired NAT Gateway.
	Tags map[string]string `pulumi:"tags"`
	// ID of the VPC that the NAT Gateway resides in.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getNatGateway.
type LookupNatGatewayResult struct {
	// ID of the EIP allocated to the selected NAT Gateway.
	AllocationId string `pulumi:"allocationId"`
	// The association ID of the Elastic IP address that's associated with the NAT Gateway. Only available when `connectivityType` is `public`.
	AssociationId string `pulumi:"associationId"`
	// Connectivity type of the NAT Gateway.
	ConnectivityType string                `pulumi:"connectivityType"`
	Filters          []GetNatGatewayFilter `pulumi:"filters"`
	Id               string                `pulumi:"id"`
	// The ID of the ENI allocated to the selected NAT Gateway.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
	// Private IP address of the selected NAT Gateway.
	PrivateIp string `pulumi:"privateIp"`
	// Public IP (EIP) address of the selected NAT Gateway.
	PublicIp string `pulumi:"publicIp"`
	// Secondary allocation EIP IDs for the selected NAT Gateway.
	SecondaryAllocationIds []string `pulumi:"secondaryAllocationIds"`
	// The number of secondary private IPv4 addresses assigned to the selected NAT Gateway.
	SecondaryPrivateIpAddressCount int `pulumi:"secondaryPrivateIpAddressCount"`
	// Secondary private IPv4 addresses assigned to the selected NAT Gateway.
	SecondaryPrivateIpAddresses []string          `pulumi:"secondaryPrivateIpAddresses"`
	State                       string            `pulumi:"state"`
	SubnetId                    string            `pulumi:"subnetId"`
	Tags                        map[string]string `pulumi:"tags"`
	VpcId                       string            `pulumi:"vpcId"`
}

func LookupNatGatewayOutput(ctx *pulumi.Context, args LookupNatGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupNatGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNatGatewayResult, error) {
			args := v.(LookupNatGatewayArgs)
			r, err := LookupNatGateway(ctx, &args, opts...)
			var s LookupNatGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNatGatewayResultOutput)
}

// A collection of arguments for invoking getNatGateway.
type LookupNatGatewayOutputArgs struct {
	// Custom filter block as described below.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Filters GetNatGatewayFilterArrayInput `pulumi:"filters"`
	// ID of the specific NAT Gateway to retrieve.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// State of the NAT Gateway (pending | failed | available | deleting | deleted ).
	State pulumi.StringPtrInput `pulumi:"state"`
	// ID of subnet that the NAT Gateway resides in.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
	// Map of tags, each pair of which must exactly match
	// a pair on the desired NAT Gateway.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// ID of the VPC that the NAT Gateway resides in.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (LookupNatGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNatGatewayArgs)(nil)).Elem()
}

// A collection of values returned by getNatGateway.
type LookupNatGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupNatGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNatGatewayResult)(nil)).Elem()
}

func (o LookupNatGatewayResultOutput) ToLookupNatGatewayResultOutput() LookupNatGatewayResultOutput {
	return o
}

func (o LookupNatGatewayResultOutput) ToLookupNatGatewayResultOutputWithContext(ctx context.Context) LookupNatGatewayResultOutput {
	return o
}

// ID of the EIP allocated to the selected NAT Gateway.
func (o LookupNatGatewayResultOutput) AllocationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.AllocationId }).(pulumi.StringOutput)
}

// The association ID of the Elastic IP address that's associated with the NAT Gateway. Only available when `connectivityType` is `public`.
func (o LookupNatGatewayResultOutput) AssociationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.AssociationId }).(pulumi.StringOutput)
}

// Connectivity type of the NAT Gateway.
func (o LookupNatGatewayResultOutput) ConnectivityType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.ConnectivityType }).(pulumi.StringOutput)
}

func (o LookupNatGatewayResultOutput) Filters() GetNatGatewayFilterArrayOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) []GetNatGatewayFilter { return v.Filters }).(GetNatGatewayFilterArrayOutput)
}

func (o LookupNatGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ID of the ENI allocated to the selected NAT Gateway.
func (o LookupNatGatewayResultOutput) NetworkInterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.NetworkInterfaceId }).(pulumi.StringOutput)
}

// Private IP address of the selected NAT Gateway.
func (o LookupNatGatewayResultOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.PrivateIp }).(pulumi.StringOutput)
}

// Public IP (EIP) address of the selected NAT Gateway.
func (o LookupNatGatewayResultOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.PublicIp }).(pulumi.StringOutput)
}

// Secondary allocation EIP IDs for the selected NAT Gateway.
func (o LookupNatGatewayResultOutput) SecondaryAllocationIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) []string { return v.SecondaryAllocationIds }).(pulumi.StringArrayOutput)
}

// The number of secondary private IPv4 addresses assigned to the selected NAT Gateway.
func (o LookupNatGatewayResultOutput) SecondaryPrivateIpAddressCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) int { return v.SecondaryPrivateIpAddressCount }).(pulumi.IntOutput)
}

// Secondary private IPv4 addresses assigned to the selected NAT Gateway.
func (o LookupNatGatewayResultOutput) SecondaryPrivateIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) []string { return v.SecondaryPrivateIpAddresses }).(pulumi.StringArrayOutput)
}

func (o LookupNatGatewayResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupNatGatewayResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o LookupNatGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNatGatewayResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNatGatewayResultOutput{})
}
