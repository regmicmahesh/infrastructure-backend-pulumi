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

// Provides an EventBridge event API Destination resource.
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
//			_, err := cloudwatch.NewEventApiDestination(ctx, "test", &cloudwatch.EventApiDestinationArgs{
//				Description:                  pulumi.String("An API Destination"),
//				InvocationEndpoint:           pulumi.String("https://api.destination.com/endpoint"),
//				HttpMethod:                   pulumi.String("POST"),
//				InvocationRateLimitPerSecond: pulumi.Int(20),
//				ConnectionArn:                pulumi.Any(aws_cloudwatch_event_connection.Test.Arn),
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
// Using `pulumi import`, import EventBridge API Destinations using the `name`. For example:
//
// ```sh
//
//	$ pulumi import aws:cloudwatch/eventApiDestination:EventApiDestination test api-destination
//
// ```
type EventApiDestination struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the event API Destination.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// ARN of the EventBridge Connection to use for the API Destination.
	ConnectionArn pulumi.StringOutput `pulumi:"connectionArn"`
	// The description of the new API Destination. Maximum of 512 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Select the HTTP method used for the invocation endpoint, such as GET, POST, PUT, etc.
	HttpMethod pulumi.StringOutput `pulumi:"httpMethod"`
	// URL endpoint to invoke as a target. This could be a valid endpoint generated by a partner service. You can include "*" as path parameters wildcards to be set from the Target HttpParameters.
	InvocationEndpoint pulumi.StringOutput `pulumi:"invocationEndpoint"`
	// Enter the maximum number of invocations per second to allow for this destination. Enter a value greater than 0 (default 300).
	InvocationRateLimitPerSecond pulumi.IntPtrOutput `pulumi:"invocationRateLimitPerSecond"`
	// The name of the new API Destination. The name must be unique for your account. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewEventApiDestination registers a new resource with the given unique name, arguments, and options.
func NewEventApiDestination(ctx *pulumi.Context,
	name string, args *EventApiDestinationArgs, opts ...pulumi.ResourceOption) (*EventApiDestination, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionArn == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionArn'")
	}
	if args.HttpMethod == nil {
		return nil, errors.New("invalid value for required argument 'HttpMethod'")
	}
	if args.InvocationEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'InvocationEndpoint'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EventApiDestination
	err := ctx.RegisterResource("aws:cloudwatch/eventApiDestination:EventApiDestination", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventApiDestination gets an existing EventApiDestination resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventApiDestination(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventApiDestinationState, opts ...pulumi.ResourceOption) (*EventApiDestination, error) {
	var resource EventApiDestination
	err := ctx.ReadResource("aws:cloudwatch/eventApiDestination:EventApiDestination", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventApiDestination resources.
type eventApiDestinationState struct {
	// The Amazon Resource Name (ARN) of the event API Destination.
	Arn *string `pulumi:"arn"`
	// ARN of the EventBridge Connection to use for the API Destination.
	ConnectionArn *string `pulumi:"connectionArn"`
	// The description of the new API Destination. Maximum of 512 characters.
	Description *string `pulumi:"description"`
	// Select the HTTP method used for the invocation endpoint, such as GET, POST, PUT, etc.
	HttpMethod *string `pulumi:"httpMethod"`
	// URL endpoint to invoke as a target. This could be a valid endpoint generated by a partner service. You can include "*" as path parameters wildcards to be set from the Target HttpParameters.
	InvocationEndpoint *string `pulumi:"invocationEndpoint"`
	// Enter the maximum number of invocations per second to allow for this destination. Enter a value greater than 0 (default 300).
	InvocationRateLimitPerSecond *int `pulumi:"invocationRateLimitPerSecond"`
	// The name of the new API Destination. The name must be unique for your account. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
	Name *string `pulumi:"name"`
}

type EventApiDestinationState struct {
	// The Amazon Resource Name (ARN) of the event API Destination.
	Arn pulumi.StringPtrInput
	// ARN of the EventBridge Connection to use for the API Destination.
	ConnectionArn pulumi.StringPtrInput
	// The description of the new API Destination. Maximum of 512 characters.
	Description pulumi.StringPtrInput
	// Select the HTTP method used for the invocation endpoint, such as GET, POST, PUT, etc.
	HttpMethod pulumi.StringPtrInput
	// URL endpoint to invoke as a target. This could be a valid endpoint generated by a partner service. You can include "*" as path parameters wildcards to be set from the Target HttpParameters.
	InvocationEndpoint pulumi.StringPtrInput
	// Enter the maximum number of invocations per second to allow for this destination. Enter a value greater than 0 (default 300).
	InvocationRateLimitPerSecond pulumi.IntPtrInput
	// The name of the new API Destination. The name must be unique for your account. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
	Name pulumi.StringPtrInput
}

func (EventApiDestinationState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventApiDestinationState)(nil)).Elem()
}

type eventApiDestinationArgs struct {
	// ARN of the EventBridge Connection to use for the API Destination.
	ConnectionArn string `pulumi:"connectionArn"`
	// The description of the new API Destination. Maximum of 512 characters.
	Description *string `pulumi:"description"`
	// Select the HTTP method used for the invocation endpoint, such as GET, POST, PUT, etc.
	HttpMethod string `pulumi:"httpMethod"`
	// URL endpoint to invoke as a target. This could be a valid endpoint generated by a partner service. You can include "*" as path parameters wildcards to be set from the Target HttpParameters.
	InvocationEndpoint string `pulumi:"invocationEndpoint"`
	// Enter the maximum number of invocations per second to allow for this destination. Enter a value greater than 0 (default 300).
	InvocationRateLimitPerSecond *int `pulumi:"invocationRateLimitPerSecond"`
	// The name of the new API Destination. The name must be unique for your account. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a EventApiDestination resource.
type EventApiDestinationArgs struct {
	// ARN of the EventBridge Connection to use for the API Destination.
	ConnectionArn pulumi.StringInput
	// The description of the new API Destination. Maximum of 512 characters.
	Description pulumi.StringPtrInput
	// Select the HTTP method used for the invocation endpoint, such as GET, POST, PUT, etc.
	HttpMethod pulumi.StringInput
	// URL endpoint to invoke as a target. This could be a valid endpoint generated by a partner service. You can include "*" as path parameters wildcards to be set from the Target HttpParameters.
	InvocationEndpoint pulumi.StringInput
	// Enter the maximum number of invocations per second to allow for this destination. Enter a value greater than 0 (default 300).
	InvocationRateLimitPerSecond pulumi.IntPtrInput
	// The name of the new API Destination. The name must be unique for your account. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
	Name pulumi.StringPtrInput
}

func (EventApiDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventApiDestinationArgs)(nil)).Elem()
}

type EventApiDestinationInput interface {
	pulumi.Input

	ToEventApiDestinationOutput() EventApiDestinationOutput
	ToEventApiDestinationOutputWithContext(ctx context.Context) EventApiDestinationOutput
}

func (*EventApiDestination) ElementType() reflect.Type {
	return reflect.TypeOf((**EventApiDestination)(nil)).Elem()
}

func (i *EventApiDestination) ToEventApiDestinationOutput() EventApiDestinationOutput {
	return i.ToEventApiDestinationOutputWithContext(context.Background())
}

func (i *EventApiDestination) ToEventApiDestinationOutputWithContext(ctx context.Context) EventApiDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventApiDestinationOutput)
}

// EventApiDestinationArrayInput is an input type that accepts EventApiDestinationArray and EventApiDestinationArrayOutput values.
// You can construct a concrete instance of `EventApiDestinationArrayInput` via:
//
//	EventApiDestinationArray{ EventApiDestinationArgs{...} }
type EventApiDestinationArrayInput interface {
	pulumi.Input

	ToEventApiDestinationArrayOutput() EventApiDestinationArrayOutput
	ToEventApiDestinationArrayOutputWithContext(context.Context) EventApiDestinationArrayOutput
}

type EventApiDestinationArray []EventApiDestinationInput

func (EventApiDestinationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventApiDestination)(nil)).Elem()
}

func (i EventApiDestinationArray) ToEventApiDestinationArrayOutput() EventApiDestinationArrayOutput {
	return i.ToEventApiDestinationArrayOutputWithContext(context.Background())
}

func (i EventApiDestinationArray) ToEventApiDestinationArrayOutputWithContext(ctx context.Context) EventApiDestinationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventApiDestinationArrayOutput)
}

// EventApiDestinationMapInput is an input type that accepts EventApiDestinationMap and EventApiDestinationMapOutput values.
// You can construct a concrete instance of `EventApiDestinationMapInput` via:
//
//	EventApiDestinationMap{ "key": EventApiDestinationArgs{...} }
type EventApiDestinationMapInput interface {
	pulumi.Input

	ToEventApiDestinationMapOutput() EventApiDestinationMapOutput
	ToEventApiDestinationMapOutputWithContext(context.Context) EventApiDestinationMapOutput
}

type EventApiDestinationMap map[string]EventApiDestinationInput

func (EventApiDestinationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventApiDestination)(nil)).Elem()
}

func (i EventApiDestinationMap) ToEventApiDestinationMapOutput() EventApiDestinationMapOutput {
	return i.ToEventApiDestinationMapOutputWithContext(context.Background())
}

func (i EventApiDestinationMap) ToEventApiDestinationMapOutputWithContext(ctx context.Context) EventApiDestinationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventApiDestinationMapOutput)
}

type EventApiDestinationOutput struct{ *pulumi.OutputState }

func (EventApiDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventApiDestination)(nil)).Elem()
}

func (o EventApiDestinationOutput) ToEventApiDestinationOutput() EventApiDestinationOutput {
	return o
}

func (o EventApiDestinationOutput) ToEventApiDestinationOutputWithContext(ctx context.Context) EventApiDestinationOutput {
	return o
}

// The Amazon Resource Name (ARN) of the event API Destination.
func (o EventApiDestinationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *EventApiDestination) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// ARN of the EventBridge Connection to use for the API Destination.
func (o EventApiDestinationOutput) ConnectionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *EventApiDestination) pulumi.StringOutput { return v.ConnectionArn }).(pulumi.StringOutput)
}

// The description of the new API Destination. Maximum of 512 characters.
func (o EventApiDestinationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventApiDestination) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Select the HTTP method used for the invocation endpoint, such as GET, POST, PUT, etc.
func (o EventApiDestinationOutput) HttpMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *EventApiDestination) pulumi.StringOutput { return v.HttpMethod }).(pulumi.StringOutput)
}

// URL endpoint to invoke as a target. This could be a valid endpoint generated by a partner service. You can include "*" as path parameters wildcards to be set from the Target HttpParameters.
func (o EventApiDestinationOutput) InvocationEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *EventApiDestination) pulumi.StringOutput { return v.InvocationEndpoint }).(pulumi.StringOutput)
}

// Enter the maximum number of invocations per second to allow for this destination. Enter a value greater than 0 (default 300).
func (o EventApiDestinationOutput) InvocationRateLimitPerSecond() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EventApiDestination) pulumi.IntPtrOutput { return v.InvocationRateLimitPerSecond }).(pulumi.IntPtrOutput)
}

// The name of the new API Destination. The name must be unique for your account. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
func (o EventApiDestinationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventApiDestination) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type EventApiDestinationArrayOutput struct{ *pulumi.OutputState }

func (EventApiDestinationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventApiDestination)(nil)).Elem()
}

func (o EventApiDestinationArrayOutput) ToEventApiDestinationArrayOutput() EventApiDestinationArrayOutput {
	return o
}

func (o EventApiDestinationArrayOutput) ToEventApiDestinationArrayOutputWithContext(ctx context.Context) EventApiDestinationArrayOutput {
	return o
}

func (o EventApiDestinationArrayOutput) Index(i pulumi.IntInput) EventApiDestinationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EventApiDestination {
		return vs[0].([]*EventApiDestination)[vs[1].(int)]
	}).(EventApiDestinationOutput)
}

type EventApiDestinationMapOutput struct{ *pulumi.OutputState }

func (EventApiDestinationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventApiDestination)(nil)).Elem()
}

func (o EventApiDestinationMapOutput) ToEventApiDestinationMapOutput() EventApiDestinationMapOutput {
	return o
}

func (o EventApiDestinationMapOutput) ToEventApiDestinationMapOutputWithContext(ctx context.Context) EventApiDestinationMapOutput {
	return o
}

func (o EventApiDestinationMapOutput) MapIndex(k pulumi.StringInput) EventApiDestinationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EventApiDestination {
		return vs[0].(map[string]*EventApiDestination)[vs[1].(string)]
	}).(EventApiDestinationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventApiDestinationInput)(nil)).Elem(), &EventApiDestination{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventApiDestinationArrayInput)(nil)).Elem(), EventApiDestinationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventApiDestinationMapInput)(nil)).Elem(), EventApiDestinationMap{})
	pulumi.RegisterOutputType(EventApiDestinationOutput{})
	pulumi.RegisterOutputType(EventApiDestinationArrayOutput{})
	pulumi.RegisterOutputType(EventApiDestinationMapOutput{})
}