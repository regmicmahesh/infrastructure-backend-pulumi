// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A Region represents any valid Amazon region that may be targeted with deployments.
type Region string

const (
	RegionAFSouth1     = Region("af-south-1")
	RegionAPEast1      = Region("ap-east-1")
	RegionAPNortheast1 = Region("ap-northeast-1")
	RegionAPNortheast2 = Region("ap-northeast-2")
	RegionAPNortheast3 = Region("ap-northeast-3")
	RegionAPSouth1     = Region("ap-south-1")
	RegionAPSouth2     = Region("ap-south-2")
	RegionAPSoutheast1 = Region("ap-southeast-1")
	RegionAPSoutheast2 = Region("ap-southeast-2")
	RegionAPSoutheast3 = Region("ap-southeast-3")
	RegionAPSoutheast4 = Region("ap-southeast-4")
	RegionCACentral    = Region("ca-central-1")
	RegionCAWest1      = Region("ca-west-1")
	RegionCNNorth1     = Region("cn-north-1")
	RegionCNNorthwest1 = Region("cn-northwest-1")
	RegionEUCentral1   = Region("eu-central-1")
	RegionEUCentral2   = Region("eu-central-2")
	RegionEUNorth1     = Region("eu-north-1")
	RegionEUSouth1     = Region("eu-south-1")
	RegionEUSouth2     = Region("eu-south-2")
	RegionEUWest1      = Region("eu-west-1")
	RegionEUWest2      = Region("eu-west-2")
	RegionEUWest3      = Region("eu-west-3")
	RegionMECentral1   = Region("me-central-1")
	RegionMESouth1     = Region("me-south-1")
	RegionSAEast1      = Region("sa-east-1")
	RegionUSGovEast1   = Region("us-gov-east-1")
	RegionUSGovWest1   = Region("us-gov-west-1")
	RegionUSEast1      = Region("us-east-1")
	RegionUSEast2      = Region("us-east-2")
	RegionUSWest1      = Region("us-west-1")
	RegionUSWest2      = Region("us-west-2")
)

func (Region) ElementType() reflect.Type {
	return reflect.TypeOf((*Region)(nil)).Elem()
}

func (e Region) ToRegionOutput() RegionOutput {
	return pulumi.ToOutput(e).(RegionOutput)
}

func (e Region) ToRegionOutputWithContext(ctx context.Context) RegionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RegionOutput)
}

func (e Region) ToRegionPtrOutput() RegionPtrOutput {
	return e.ToRegionPtrOutputWithContext(context.Background())
}

func (e Region) ToRegionPtrOutputWithContext(ctx context.Context) RegionPtrOutput {
	return Region(e).ToRegionOutputWithContext(ctx).ToRegionPtrOutputWithContext(ctx)
}

func (e Region) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Region) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Region) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Region) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RegionOutput struct{ *pulumi.OutputState }

func (RegionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Region)(nil)).Elem()
}

func (o RegionOutput) ToRegionOutput() RegionOutput {
	return o
}

func (o RegionOutput) ToRegionOutputWithContext(ctx context.Context) RegionOutput {
	return o
}

func (o RegionOutput) ToRegionPtrOutput() RegionPtrOutput {
	return o.ToRegionPtrOutputWithContext(context.Background())
}

func (o RegionOutput) ToRegionPtrOutputWithContext(ctx context.Context) RegionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Region) *Region {
		return &v
	}).(RegionPtrOutput)
}

func (o RegionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RegionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Region) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RegionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RegionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Region) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RegionPtrOutput struct{ *pulumi.OutputState }

func (RegionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Region)(nil)).Elem()
}

func (o RegionPtrOutput) ToRegionPtrOutput() RegionPtrOutput {
	return o
}

func (o RegionPtrOutput) ToRegionPtrOutputWithContext(ctx context.Context) RegionPtrOutput {
	return o
}

func (o RegionPtrOutput) Elem() RegionOutput {
	return o.ApplyT(func(v *Region) Region {
		if v != nil {
			return *v
		}
		var ret Region
		return ret
	}).(RegionOutput)
}

func (o RegionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RegionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Region) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// RegionInput is an input type that accepts values of the Region enum
// A concrete instance of `RegionInput` can be one of the following:
//
//	RegionAFSouth1
//	RegionAPEast1
//	RegionAPNortheast1
//	RegionAPNortheast2
//	RegionAPNortheast3
//	RegionAPSouth1
//	RegionAPSouth2
//	RegionAPSoutheast1
//	RegionAPSoutheast2
//	RegionAPSoutheast3
//	RegionAPSoutheast4
//	RegionCACentral
//	RegionCAWest1
//	RegionCNNorth1
//	RegionCNNorthwest1
//	RegionEUCentral1
//	RegionEUCentral2
//	RegionEUNorth1
//	RegionEUSouth1
//	RegionEUSouth2
//	RegionEUWest1
//	RegionEUWest2
//	RegionEUWest3
//	RegionMECentral1
//	RegionMESouth1
//	RegionSAEast1
//	RegionUSGovEast1
//	RegionUSGovWest1
//	RegionUSEast1
//	RegionUSEast2
//	RegionUSWest1
//	RegionUSWest2
type RegionInput interface {
	pulumi.Input

	ToRegionOutput() RegionOutput
	ToRegionOutputWithContext(context.Context) RegionOutput
}

var regionPtrType = reflect.TypeOf((**Region)(nil)).Elem()

type RegionPtrInput interface {
	pulumi.Input

	ToRegionPtrOutput() RegionPtrOutput
	ToRegionPtrOutputWithContext(context.Context) RegionPtrOutput
}

type regionPtr string

func RegionPtr(v string) RegionPtrInput {
	return (*regionPtr)(&v)
}

func (*regionPtr) ElementType() reflect.Type {
	return regionPtrType
}

func (in *regionPtr) ToRegionPtrOutput() RegionPtrOutput {
	return pulumi.ToOutput(in).(RegionPtrOutput)
}

func (in *regionPtr) ToRegionPtrOutputWithContext(ctx context.Context) RegionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RegionPtrOutput)
}

func (in *regionPtr) ToOutput(ctx context.Context) pulumix.Output[*Region] {
	return pulumix.Output[*Region]{
		OutputState: in.ToRegionPtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionInput)(nil)).Elem(), Region("af-south-1"))
	pulumi.RegisterInputType(reflect.TypeOf((*RegionPtrInput)(nil)).Elem(), Region("af-south-1"))
	pulumi.RegisterOutputType(RegionOutput{})
	pulumi.RegisterOutputType(RegionPtrOutput{})
}
