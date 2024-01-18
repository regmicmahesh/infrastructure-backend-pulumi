// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an EventBridge event archive resource.
//
// > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			orderEventBus, err := cloudwatch.NewEventBus(ctx, "orderEventBus", nil)
//			if err != nil {
//				return err
//			}
//			_, err = cloudwatch.NewEventArchive(ctx, "orderEventArchive", &cloudwatch.EventArchiveArgs{
//				EventSourceArn: orderEventBus.Arn,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Example all optional arguments
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			orderEventBus, err := cloudwatch.NewEventBus(ctx, "orderEventBus", nil)
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"source": []string{
//					"company.team.order",
//				},
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = cloudwatch.NewEventArchive(ctx, "orderEventArchive", &cloudwatch.EventArchiveArgs{
//				Description:    pulumi.String("Archived events from order service"),
//				EventSourceArn: orderEventBus.Arn,
//				RetentionDays:  pulumi.Int(7),
//				EventPattern:   pulumi.String(json0),
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
// Using `pulumi import`, import an EventBridge archive using the `name`. For example:
//
// ```sh
//
//	$ pulumi import aws:cloudwatch/eventArchive:EventArchive imported_event_archive order-archive
//
// ```
type EventArchive struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the event archive.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the new event archive.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the `eventSourceArn`.
	EventPattern pulumi.StringPtrOutput `pulumi:"eventPattern"`
	// Event bus source ARN from where these events should be archived.
	EventSourceArn pulumi.StringOutput `pulumi:"eventSourceArn"`
	// The name of the new event archive. The archive name cannot exceed 48 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
}

// NewEventArchive registers a new resource with the given unique name, arguments, and options.
func NewEventArchive(ctx *pulumi.Context,
	name string, args *EventArchiveArgs, opts ...pulumi.ResourceOption) (*EventArchive, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventSourceArn == nil {
		return nil, errors.New("invalid value for required argument 'EventSourceArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EventArchive
	err := ctx.RegisterResource("aws:cloudwatch/eventArchive:EventArchive", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventArchive gets an existing EventArchive resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventArchive(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventArchiveState, opts ...pulumi.ResourceOption) (*EventArchive, error) {
	var resource EventArchive
	err := ctx.ReadResource("aws:cloudwatch/eventArchive:EventArchive", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventArchive resources.
type eventArchiveState struct {
	// The Amazon Resource Name (ARN) of the event archive.
	Arn *string `pulumi:"arn"`
	// The description of the new event archive.
	Description *string `pulumi:"description"`
	// Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the `eventSourceArn`.
	EventPattern *string `pulumi:"eventPattern"`
	// Event bus source ARN from where these events should be archived.
	EventSourceArn *string `pulumi:"eventSourceArn"`
	// The name of the new event archive. The archive name cannot exceed 48 characters.
	Name *string `pulumi:"name"`
	// The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
	RetentionDays *int `pulumi:"retentionDays"`
}

type EventArchiveState struct {
	// The Amazon Resource Name (ARN) of the event archive.
	Arn pulumi.StringPtrInput
	// The description of the new event archive.
	Description pulumi.StringPtrInput
	// Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the `eventSourceArn`.
	EventPattern pulumi.StringPtrInput
	// Event bus source ARN from where these events should be archived.
	EventSourceArn pulumi.StringPtrInput
	// The name of the new event archive. The archive name cannot exceed 48 characters.
	Name pulumi.StringPtrInput
	// The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
	RetentionDays pulumi.IntPtrInput
}

func (EventArchiveState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventArchiveState)(nil)).Elem()
}

type eventArchiveArgs struct {
	// The description of the new event archive.
	Description *string `pulumi:"description"`
	// Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the `eventSourceArn`.
	EventPattern *string `pulumi:"eventPattern"`
	// Event bus source ARN from where these events should be archived.
	EventSourceArn string `pulumi:"eventSourceArn"`
	// The name of the new event archive. The archive name cannot exceed 48 characters.
	Name *string `pulumi:"name"`
	// The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
	RetentionDays *int `pulumi:"retentionDays"`
}

// The set of arguments for constructing a EventArchive resource.
type EventArchiveArgs struct {
	// The description of the new event archive.
	Description pulumi.StringPtrInput
	// Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the `eventSourceArn`.
	EventPattern pulumi.StringPtrInput
	// Event bus source ARN from where these events should be archived.
	EventSourceArn pulumi.StringInput
	// The name of the new event archive. The archive name cannot exceed 48 characters.
	Name pulumi.StringPtrInput
	// The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
	RetentionDays pulumi.IntPtrInput
}

func (EventArchiveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventArchiveArgs)(nil)).Elem()
}

type EventArchiveInput interface {
	pulumi.Input

	ToEventArchiveOutput() EventArchiveOutput
	ToEventArchiveOutputWithContext(ctx context.Context) EventArchiveOutput
}

func (*EventArchive) ElementType() reflect.Type {
	return reflect.TypeOf((**EventArchive)(nil)).Elem()
}

func (i *EventArchive) ToEventArchiveOutput() EventArchiveOutput {
	return i.ToEventArchiveOutputWithContext(context.Background())
}

func (i *EventArchive) ToEventArchiveOutputWithContext(ctx context.Context) EventArchiveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventArchiveOutput)
}

// EventArchiveArrayInput is an input type that accepts EventArchiveArray and EventArchiveArrayOutput values.
// You can construct a concrete instance of `EventArchiveArrayInput` via:
//
//	EventArchiveArray{ EventArchiveArgs{...} }
type EventArchiveArrayInput interface {
	pulumi.Input

	ToEventArchiveArrayOutput() EventArchiveArrayOutput
	ToEventArchiveArrayOutputWithContext(context.Context) EventArchiveArrayOutput
}

type EventArchiveArray []EventArchiveInput

func (EventArchiveArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventArchive)(nil)).Elem()
}

func (i EventArchiveArray) ToEventArchiveArrayOutput() EventArchiveArrayOutput {
	return i.ToEventArchiveArrayOutputWithContext(context.Background())
}

func (i EventArchiveArray) ToEventArchiveArrayOutputWithContext(ctx context.Context) EventArchiveArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventArchiveArrayOutput)
}

// EventArchiveMapInput is an input type that accepts EventArchiveMap and EventArchiveMapOutput values.
// You can construct a concrete instance of `EventArchiveMapInput` via:
//
//	EventArchiveMap{ "key": EventArchiveArgs{...} }
type EventArchiveMapInput interface {
	pulumi.Input

	ToEventArchiveMapOutput() EventArchiveMapOutput
	ToEventArchiveMapOutputWithContext(context.Context) EventArchiveMapOutput
}

type EventArchiveMap map[string]EventArchiveInput

func (EventArchiveMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventArchive)(nil)).Elem()
}

func (i EventArchiveMap) ToEventArchiveMapOutput() EventArchiveMapOutput {
	return i.ToEventArchiveMapOutputWithContext(context.Background())
}

func (i EventArchiveMap) ToEventArchiveMapOutputWithContext(ctx context.Context) EventArchiveMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventArchiveMapOutput)
}

type EventArchiveOutput struct{ *pulumi.OutputState }

func (EventArchiveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventArchive)(nil)).Elem()
}

func (o EventArchiveOutput) ToEventArchiveOutput() EventArchiveOutput {
	return o
}

func (o EventArchiveOutput) ToEventArchiveOutputWithContext(ctx context.Context) EventArchiveOutput {
	return o
}

// The Amazon Resource Name (ARN) of the event archive.
func (o EventArchiveOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *EventArchive) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The description of the new event archive.
func (o EventArchiveOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventArchive) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the `eventSourceArn`.
func (o EventArchiveOutput) EventPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventArchive) pulumi.StringPtrOutput { return v.EventPattern }).(pulumi.StringPtrOutput)
}

// Event bus source ARN from where these events should be archived.
func (o EventArchiveOutput) EventSourceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *EventArchive) pulumi.StringOutput { return v.EventSourceArn }).(pulumi.StringOutput)
}

// The name of the new event archive. The archive name cannot exceed 48 characters.
func (o EventArchiveOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventArchive) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
func (o EventArchiveOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EventArchive) pulumi.IntPtrOutput { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

type EventArchiveArrayOutput struct{ *pulumi.OutputState }

func (EventArchiveArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventArchive)(nil)).Elem()
}

func (o EventArchiveArrayOutput) ToEventArchiveArrayOutput() EventArchiveArrayOutput {
	return o
}

func (o EventArchiveArrayOutput) ToEventArchiveArrayOutputWithContext(ctx context.Context) EventArchiveArrayOutput {
	return o
}

func (o EventArchiveArrayOutput) Index(i pulumi.IntInput) EventArchiveOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EventArchive {
		return vs[0].([]*EventArchive)[vs[1].(int)]
	}).(EventArchiveOutput)
}

type EventArchiveMapOutput struct{ *pulumi.OutputState }

func (EventArchiveMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventArchive)(nil)).Elem()
}

func (o EventArchiveMapOutput) ToEventArchiveMapOutput() EventArchiveMapOutput {
	return o
}

func (o EventArchiveMapOutput) ToEventArchiveMapOutputWithContext(ctx context.Context) EventArchiveMapOutput {
	return o
}

func (o EventArchiveMapOutput) MapIndex(k pulumi.StringInput) EventArchiveOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EventArchive {
		return vs[0].(map[string]*EventArchive)[vs[1].(string)]
	}).(EventArchiveOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventArchiveInput)(nil)).Elem(), &EventArchive{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventArchiveArrayInput)(nil)).Elem(), EventArchiveArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventArchiveMapInput)(nil)).Elem(), EventArchiveMap{})
	pulumi.RegisterOutputType(EventArchiveOutput{})
	pulumi.RegisterOutputType(EventArchiveArrayOutput{})
	pulumi.RegisterOutputType(EventArchiveMapOutput{})
}
