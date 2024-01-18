// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicediscovery

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicediscovery"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := servicediscovery.NewHttpNamespace(ctx, "example", &servicediscovery.HttpNamespaceArgs{
//				Description: pulumi.String("example"),
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
// Using `pulumi import`, import Service Discovery HTTP Namespace using the namespace ID. For example:
//
// ```sh
//
//	$ pulumi import aws:servicediscovery/httpNamespace:HttpNamespace example ns-1234567890
//
// ```
type HttpNamespace struct {
	pulumi.CustomResourceState

	// The ARN that Amazon Route 53 assigns to the namespace when you create it.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description that you specify for the namespace when you create it.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of an HTTP namespace.
	HttpName pulumi.StringOutput `pulumi:"httpName"`
	// The name of the http namespace.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of tags to assign to the namespace. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewHttpNamespace registers a new resource with the given unique name, arguments, and options.
func NewHttpNamespace(ctx *pulumi.Context,
	name string, args *HttpNamespaceArgs, opts ...pulumi.ResourceOption) (*HttpNamespace, error) {
	if args == nil {
		args = &HttpNamespaceArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource HttpNamespace
	err := ctx.RegisterResource("aws:servicediscovery/httpNamespace:HttpNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHttpNamespace gets an existing HttpNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHttpNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HttpNamespaceState, opts ...pulumi.ResourceOption) (*HttpNamespace, error) {
	var resource HttpNamespace
	err := ctx.ReadResource("aws:servicediscovery/httpNamespace:HttpNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HttpNamespace resources.
type httpNamespaceState struct {
	// The ARN that Amazon Route 53 assigns to the namespace when you create it.
	Arn *string `pulumi:"arn"`
	// The description that you specify for the namespace when you create it.
	Description *string `pulumi:"description"`
	// The name of an HTTP namespace.
	HttpName *string `pulumi:"httpName"`
	// The name of the http namespace.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the namespace. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type HttpNamespaceState struct {
	// The ARN that Amazon Route 53 assigns to the namespace when you create it.
	Arn pulumi.StringPtrInput
	// The description that you specify for the namespace when you create it.
	Description pulumi.StringPtrInput
	// The name of an HTTP namespace.
	HttpName pulumi.StringPtrInput
	// The name of the http namespace.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the namespace. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (HttpNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*httpNamespaceState)(nil)).Elem()
}

type httpNamespaceArgs struct {
	// The description that you specify for the namespace when you create it.
	Description *string `pulumi:"description"`
	// The name of the http namespace.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the namespace. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a HttpNamespace resource.
type HttpNamespaceArgs struct {
	// The description that you specify for the namespace when you create it.
	Description pulumi.StringPtrInput
	// The name of the http namespace.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the namespace. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (HttpNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*httpNamespaceArgs)(nil)).Elem()
}

type HttpNamespaceInput interface {
	pulumi.Input

	ToHttpNamespaceOutput() HttpNamespaceOutput
	ToHttpNamespaceOutputWithContext(ctx context.Context) HttpNamespaceOutput
}

func (*HttpNamespace) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpNamespace)(nil)).Elem()
}

func (i *HttpNamespace) ToHttpNamespaceOutput() HttpNamespaceOutput {
	return i.ToHttpNamespaceOutputWithContext(context.Background())
}

func (i *HttpNamespace) ToHttpNamespaceOutputWithContext(ctx context.Context) HttpNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpNamespaceOutput)
}

// HttpNamespaceArrayInput is an input type that accepts HttpNamespaceArray and HttpNamespaceArrayOutput values.
// You can construct a concrete instance of `HttpNamespaceArrayInput` via:
//
//	HttpNamespaceArray{ HttpNamespaceArgs{...} }
type HttpNamespaceArrayInput interface {
	pulumi.Input

	ToHttpNamespaceArrayOutput() HttpNamespaceArrayOutput
	ToHttpNamespaceArrayOutputWithContext(context.Context) HttpNamespaceArrayOutput
}

type HttpNamespaceArray []HttpNamespaceInput

func (HttpNamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HttpNamespace)(nil)).Elem()
}

func (i HttpNamespaceArray) ToHttpNamespaceArrayOutput() HttpNamespaceArrayOutput {
	return i.ToHttpNamespaceArrayOutputWithContext(context.Background())
}

func (i HttpNamespaceArray) ToHttpNamespaceArrayOutputWithContext(ctx context.Context) HttpNamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpNamespaceArrayOutput)
}

// HttpNamespaceMapInput is an input type that accepts HttpNamespaceMap and HttpNamespaceMapOutput values.
// You can construct a concrete instance of `HttpNamespaceMapInput` via:
//
//	HttpNamespaceMap{ "key": HttpNamespaceArgs{...} }
type HttpNamespaceMapInput interface {
	pulumi.Input

	ToHttpNamespaceMapOutput() HttpNamespaceMapOutput
	ToHttpNamespaceMapOutputWithContext(context.Context) HttpNamespaceMapOutput
}

type HttpNamespaceMap map[string]HttpNamespaceInput

func (HttpNamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HttpNamespace)(nil)).Elem()
}

func (i HttpNamespaceMap) ToHttpNamespaceMapOutput() HttpNamespaceMapOutput {
	return i.ToHttpNamespaceMapOutputWithContext(context.Background())
}

func (i HttpNamespaceMap) ToHttpNamespaceMapOutputWithContext(ctx context.Context) HttpNamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpNamespaceMapOutput)
}

type HttpNamespaceOutput struct{ *pulumi.OutputState }

func (HttpNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpNamespace)(nil)).Elem()
}

func (o HttpNamespaceOutput) ToHttpNamespaceOutput() HttpNamespaceOutput {
	return o
}

func (o HttpNamespaceOutput) ToHttpNamespaceOutputWithContext(ctx context.Context) HttpNamespaceOutput {
	return o
}

// The ARN that Amazon Route 53 assigns to the namespace when you create it.
func (o HttpNamespaceOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpNamespace) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The description that you specify for the namespace when you create it.
func (o HttpNamespaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpNamespace) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of an HTTP namespace.
func (o HttpNamespaceOutput) HttpName() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpNamespace) pulumi.StringOutput { return v.HttpName }).(pulumi.StringOutput)
}

// The name of the http namespace.
func (o HttpNamespaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpNamespace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A map of tags to assign to the namespace. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o HttpNamespaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HttpNamespace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o HttpNamespaceOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HttpNamespace) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type HttpNamespaceArrayOutput struct{ *pulumi.OutputState }

func (HttpNamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HttpNamespace)(nil)).Elem()
}

func (o HttpNamespaceArrayOutput) ToHttpNamespaceArrayOutput() HttpNamespaceArrayOutput {
	return o
}

func (o HttpNamespaceArrayOutput) ToHttpNamespaceArrayOutputWithContext(ctx context.Context) HttpNamespaceArrayOutput {
	return o
}

func (o HttpNamespaceArrayOutput) Index(i pulumi.IntInput) HttpNamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HttpNamespace {
		return vs[0].([]*HttpNamespace)[vs[1].(int)]
	}).(HttpNamespaceOutput)
}

type HttpNamespaceMapOutput struct{ *pulumi.OutputState }

func (HttpNamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HttpNamespace)(nil)).Elem()
}

func (o HttpNamespaceMapOutput) ToHttpNamespaceMapOutput() HttpNamespaceMapOutput {
	return o
}

func (o HttpNamespaceMapOutput) ToHttpNamespaceMapOutputWithContext(ctx context.Context) HttpNamespaceMapOutput {
	return o
}

func (o HttpNamespaceMapOutput) MapIndex(k pulumi.StringInput) HttpNamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HttpNamespace {
		return vs[0].(map[string]*HttpNamespace)[vs[1].(string)]
	}).(HttpNamespaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HttpNamespaceInput)(nil)).Elem(), &HttpNamespace{})
	pulumi.RegisterInputType(reflect.TypeOf((*HttpNamespaceArrayInput)(nil)).Elem(), HttpNamespaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HttpNamespaceMapInput)(nil)).Elem(), HttpNamespaceMap{})
	pulumi.RegisterOutputType(HttpNamespaceOutput{})
	pulumi.RegisterOutputType(HttpNamespaceArrayOutput{})
	pulumi.RegisterOutputType(HttpNamespaceMapOutput{})
}
