// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type RecordType string

const (
	RecordTypeA     = RecordType("A")
	RecordTypeAAAA  = RecordType("AAAA")
	RecordTypeCNAME = RecordType("CNAME")
	RecordTypeCAA   = RecordType("CAA")
	RecordTypeMX    = RecordType("MX")
	RecordTypeNAPTR = RecordType("NAPTR")
	RecordTypeNS    = RecordType("NS")
	RecordTypePTR   = RecordType("PTR")
	RecordTypeSOA   = RecordType("SOA")
	RecordTypeSPF   = RecordType("SPF")
	RecordTypeSRV   = RecordType("SRV")
	RecordTypeTXT   = RecordType("TXT")
)

func (RecordType) ElementType() reflect.Type {
	return reflect.TypeOf((*RecordType)(nil)).Elem()
}

func (e RecordType) ToRecordTypeOutput() RecordTypeOutput {
	return pulumi.ToOutput(e).(RecordTypeOutput)
}

func (e RecordType) ToRecordTypeOutputWithContext(ctx context.Context) RecordTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RecordTypeOutput)
}

func (e RecordType) ToRecordTypePtrOutput() RecordTypePtrOutput {
	return e.ToRecordTypePtrOutputWithContext(context.Background())
}

func (e RecordType) ToRecordTypePtrOutputWithContext(ctx context.Context) RecordTypePtrOutput {
	return RecordType(e).ToRecordTypeOutputWithContext(ctx).ToRecordTypePtrOutputWithContext(ctx)
}

func (e RecordType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecordType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecordType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RecordType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RecordTypeOutput struct{ *pulumi.OutputState }

func (RecordTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecordType)(nil)).Elem()
}

func (o RecordTypeOutput) ToRecordTypeOutput() RecordTypeOutput {
	return o
}

func (o RecordTypeOutput) ToRecordTypeOutputWithContext(ctx context.Context) RecordTypeOutput {
	return o
}

func (o RecordTypeOutput) ToRecordTypePtrOutput() RecordTypePtrOutput {
	return o.ToRecordTypePtrOutputWithContext(context.Background())
}

func (o RecordTypeOutput) ToRecordTypePtrOutputWithContext(ctx context.Context) RecordTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecordType) *RecordType {
		return &v
	}).(RecordTypePtrOutput)
}

func (o RecordTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RecordTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecordType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RecordTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecordTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecordType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RecordTypePtrOutput struct{ *pulumi.OutputState }

func (RecordTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecordType)(nil)).Elem()
}

func (o RecordTypePtrOutput) ToRecordTypePtrOutput() RecordTypePtrOutput {
	return o
}

func (o RecordTypePtrOutput) ToRecordTypePtrOutputWithContext(ctx context.Context) RecordTypePtrOutput {
	return o
}

func (o RecordTypePtrOutput) Elem() RecordTypeOutput {
	return o.ApplyT(func(v *RecordType) RecordType {
		if v != nil {
			return *v
		}
		var ret RecordType
		return ret
	}).(RecordTypeOutput)
}

func (o RecordTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecordTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RecordType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// RecordTypeInput is an input type that accepts values of the RecordType enum
// A concrete instance of `RecordTypeInput` can be one of the following:
//
//	RecordTypeA
//	RecordTypeAAAA
//	RecordTypeCNAME
//	RecordTypeCAA
//	RecordTypeMX
//	RecordTypeNAPTR
//	RecordTypeNS
//	RecordTypePTR
//	RecordTypeSOA
//	RecordTypeSPF
//	RecordTypeSRV
//	RecordTypeTXT
type RecordTypeInput interface {
	pulumi.Input

	ToRecordTypeOutput() RecordTypeOutput
	ToRecordTypeOutputWithContext(context.Context) RecordTypeOutput
}

var recordTypePtrType = reflect.TypeOf((**RecordType)(nil)).Elem()

type RecordTypePtrInput interface {
	pulumi.Input

	ToRecordTypePtrOutput() RecordTypePtrOutput
	ToRecordTypePtrOutputWithContext(context.Context) RecordTypePtrOutput
}

type recordTypePtr string

func RecordTypePtr(v string) RecordTypePtrInput {
	return (*recordTypePtr)(&v)
}

func (*recordTypePtr) ElementType() reflect.Type {
	return recordTypePtrType
}

func (in *recordTypePtr) ToRecordTypePtrOutput() RecordTypePtrOutput {
	return pulumi.ToOutput(in).(RecordTypePtrOutput)
}

func (in *recordTypePtr) ToRecordTypePtrOutputWithContext(ctx context.Context) RecordTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RecordTypePtrOutput)
}

func (in *recordTypePtr) ToOutput(ctx context.Context) pulumix.Output[*RecordType] {
	return pulumix.Output[*RecordType]{
		OutputState: in.ToRecordTypePtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RecordTypeInput)(nil)).Elem(), RecordType("A"))
	pulumi.RegisterInputType(reflect.TypeOf((*RecordTypePtrInput)(nil)).Elem(), RecordType("A"))
	pulumi.RegisterOutputType(RecordTypeOutput{})
	pulumi.RegisterOutputType(RecordTypePtrOutput{})
}
