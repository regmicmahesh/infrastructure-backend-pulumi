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

// Provides a CloudWatch Log Data Protection Policy resource.
//
// Read more about protecting sensitive user data in the [User Guide](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", nil)
//			if err != nil {
//				return err
//			}
//			exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
//			if err != nil {
//				return err
//			}
//			_, err = cloudwatch.NewLogDataProtectionPolicy(ctx, "exampleLogDataProtectionPolicy", &cloudwatch.LogDataProtectionPolicyArgs{
//				LogGroupName: exampleLogGroup.Name,
//				PolicyDocument: exampleBucketV2.Bucket.ApplyT(func(bucket string) (pulumi.String, error) {
//					var _zero pulumi.String
//					tmpJSON0, err := json.Marshal(map[string]interface{}{
//						"Name":    "Example",
//						"Version": "2021-06-01",
//						"Statement": []interface{}{
//							map[string]interface{}{
//								"Sid": "Audit",
//								"DataIdentifier": []string{
//									"arn:aws:dataprotection::aws:data-identifier/EmailAddress",
//								},
//								"Operation": map[string]interface{}{
//									"Audit": map[string]interface{}{
//										"FindingsDestination": map[string]interface{}{
//											"S3": map[string]interface{}{
//												"Bucket": bucket,
//											},
//										},
//									},
//								},
//							},
//							map[string]interface{}{
//								"Sid": "Redact",
//								"DataIdentifier": []string{
//									"arn:aws:dataprotection::aws:data-identifier/EmailAddress",
//								},
//								"Operation": map[string]interface{}{
//									"Deidentify": map[string]interface{}{
//										"MaskConfig": nil,
//									},
//								},
//							},
//						},
//					})
//					if err != nil {
//						return _zero, err
//					}
//					json0 := string(tmpJSON0)
//					return pulumi.String(json0), nil
//				}).(pulumi.StringOutput),
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
// Using `pulumi import`, import this resource using the `log_group_name`. For example:
//
// ```sh
//
//	$ pulumi import aws:cloudwatch/logDataProtectionPolicy:LogDataProtectionPolicy example my-log-group
//
// ```
type LogDataProtectionPolicy struct {
	pulumi.CustomResourceState

	// The name of the log group under which the log stream is to be created.
	LogGroupName pulumi.StringOutput `pulumi:"logGroupName"`
	// Specifies the data protection policy in JSON. Read more at [Data protection policy syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data-start.html#mask-sensitive-log-data-policysyntax).
	PolicyDocument pulumi.StringOutput `pulumi:"policyDocument"`
}

// NewLogDataProtectionPolicy registers a new resource with the given unique name, arguments, and options.
func NewLogDataProtectionPolicy(ctx *pulumi.Context,
	name string, args *LogDataProtectionPolicyArgs, opts ...pulumi.ResourceOption) (*LogDataProtectionPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LogGroupName == nil {
		return nil, errors.New("invalid value for required argument 'LogGroupName'")
	}
	if args.PolicyDocument == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDocument'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LogDataProtectionPolicy
	err := ctx.RegisterResource("aws:cloudwatch/logDataProtectionPolicy:LogDataProtectionPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogDataProtectionPolicy gets an existing LogDataProtectionPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogDataProtectionPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogDataProtectionPolicyState, opts ...pulumi.ResourceOption) (*LogDataProtectionPolicy, error) {
	var resource LogDataProtectionPolicy
	err := ctx.ReadResource("aws:cloudwatch/logDataProtectionPolicy:LogDataProtectionPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogDataProtectionPolicy resources.
type logDataProtectionPolicyState struct {
	// The name of the log group under which the log stream is to be created.
	LogGroupName *string `pulumi:"logGroupName"`
	// Specifies the data protection policy in JSON. Read more at [Data protection policy syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data-start.html#mask-sensitive-log-data-policysyntax).
	PolicyDocument *string `pulumi:"policyDocument"`
}

type LogDataProtectionPolicyState struct {
	// The name of the log group under which the log stream is to be created.
	LogGroupName pulumi.StringPtrInput
	// Specifies the data protection policy in JSON. Read more at [Data protection policy syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data-start.html#mask-sensitive-log-data-policysyntax).
	PolicyDocument pulumi.StringPtrInput
}

func (LogDataProtectionPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*logDataProtectionPolicyState)(nil)).Elem()
}

type logDataProtectionPolicyArgs struct {
	// The name of the log group under which the log stream is to be created.
	LogGroupName string `pulumi:"logGroupName"`
	// Specifies the data protection policy in JSON. Read more at [Data protection policy syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data-start.html#mask-sensitive-log-data-policysyntax).
	PolicyDocument string `pulumi:"policyDocument"`
}

// The set of arguments for constructing a LogDataProtectionPolicy resource.
type LogDataProtectionPolicyArgs struct {
	// The name of the log group under which the log stream is to be created.
	LogGroupName pulumi.StringInput
	// Specifies the data protection policy in JSON. Read more at [Data protection policy syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data-start.html#mask-sensitive-log-data-policysyntax).
	PolicyDocument pulumi.StringInput
}

func (LogDataProtectionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logDataProtectionPolicyArgs)(nil)).Elem()
}

type LogDataProtectionPolicyInput interface {
	pulumi.Input

	ToLogDataProtectionPolicyOutput() LogDataProtectionPolicyOutput
	ToLogDataProtectionPolicyOutputWithContext(ctx context.Context) LogDataProtectionPolicyOutput
}

func (*LogDataProtectionPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**LogDataProtectionPolicy)(nil)).Elem()
}

func (i *LogDataProtectionPolicy) ToLogDataProtectionPolicyOutput() LogDataProtectionPolicyOutput {
	return i.ToLogDataProtectionPolicyOutputWithContext(context.Background())
}

func (i *LogDataProtectionPolicy) ToLogDataProtectionPolicyOutputWithContext(ctx context.Context) LogDataProtectionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogDataProtectionPolicyOutput)
}

// LogDataProtectionPolicyArrayInput is an input type that accepts LogDataProtectionPolicyArray and LogDataProtectionPolicyArrayOutput values.
// You can construct a concrete instance of `LogDataProtectionPolicyArrayInput` via:
//
//	LogDataProtectionPolicyArray{ LogDataProtectionPolicyArgs{...} }
type LogDataProtectionPolicyArrayInput interface {
	pulumi.Input

	ToLogDataProtectionPolicyArrayOutput() LogDataProtectionPolicyArrayOutput
	ToLogDataProtectionPolicyArrayOutputWithContext(context.Context) LogDataProtectionPolicyArrayOutput
}

type LogDataProtectionPolicyArray []LogDataProtectionPolicyInput

func (LogDataProtectionPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogDataProtectionPolicy)(nil)).Elem()
}

func (i LogDataProtectionPolicyArray) ToLogDataProtectionPolicyArrayOutput() LogDataProtectionPolicyArrayOutput {
	return i.ToLogDataProtectionPolicyArrayOutputWithContext(context.Background())
}

func (i LogDataProtectionPolicyArray) ToLogDataProtectionPolicyArrayOutputWithContext(ctx context.Context) LogDataProtectionPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogDataProtectionPolicyArrayOutput)
}

// LogDataProtectionPolicyMapInput is an input type that accepts LogDataProtectionPolicyMap and LogDataProtectionPolicyMapOutput values.
// You can construct a concrete instance of `LogDataProtectionPolicyMapInput` via:
//
//	LogDataProtectionPolicyMap{ "key": LogDataProtectionPolicyArgs{...} }
type LogDataProtectionPolicyMapInput interface {
	pulumi.Input

	ToLogDataProtectionPolicyMapOutput() LogDataProtectionPolicyMapOutput
	ToLogDataProtectionPolicyMapOutputWithContext(context.Context) LogDataProtectionPolicyMapOutput
}

type LogDataProtectionPolicyMap map[string]LogDataProtectionPolicyInput

func (LogDataProtectionPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogDataProtectionPolicy)(nil)).Elem()
}

func (i LogDataProtectionPolicyMap) ToLogDataProtectionPolicyMapOutput() LogDataProtectionPolicyMapOutput {
	return i.ToLogDataProtectionPolicyMapOutputWithContext(context.Background())
}

func (i LogDataProtectionPolicyMap) ToLogDataProtectionPolicyMapOutputWithContext(ctx context.Context) LogDataProtectionPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogDataProtectionPolicyMapOutput)
}

type LogDataProtectionPolicyOutput struct{ *pulumi.OutputState }

func (LogDataProtectionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogDataProtectionPolicy)(nil)).Elem()
}

func (o LogDataProtectionPolicyOutput) ToLogDataProtectionPolicyOutput() LogDataProtectionPolicyOutput {
	return o
}

func (o LogDataProtectionPolicyOutput) ToLogDataProtectionPolicyOutputWithContext(ctx context.Context) LogDataProtectionPolicyOutput {
	return o
}

// The name of the log group under which the log stream is to be created.
func (o LogDataProtectionPolicyOutput) LogGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDataProtectionPolicy) pulumi.StringOutput { return v.LogGroupName }).(pulumi.StringOutput)
}

// Specifies the data protection policy in JSON. Read more at [Data protection policy syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data-start.html#mask-sensitive-log-data-policysyntax).
func (o LogDataProtectionPolicyOutput) PolicyDocument() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDataProtectionPolicy) pulumi.StringOutput { return v.PolicyDocument }).(pulumi.StringOutput)
}

type LogDataProtectionPolicyArrayOutput struct{ *pulumi.OutputState }

func (LogDataProtectionPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogDataProtectionPolicy)(nil)).Elem()
}

func (o LogDataProtectionPolicyArrayOutput) ToLogDataProtectionPolicyArrayOutput() LogDataProtectionPolicyArrayOutput {
	return o
}

func (o LogDataProtectionPolicyArrayOutput) ToLogDataProtectionPolicyArrayOutputWithContext(ctx context.Context) LogDataProtectionPolicyArrayOutput {
	return o
}

func (o LogDataProtectionPolicyArrayOutput) Index(i pulumi.IntInput) LogDataProtectionPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogDataProtectionPolicy {
		return vs[0].([]*LogDataProtectionPolicy)[vs[1].(int)]
	}).(LogDataProtectionPolicyOutput)
}

type LogDataProtectionPolicyMapOutput struct{ *pulumi.OutputState }

func (LogDataProtectionPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogDataProtectionPolicy)(nil)).Elem()
}

func (o LogDataProtectionPolicyMapOutput) ToLogDataProtectionPolicyMapOutput() LogDataProtectionPolicyMapOutput {
	return o
}

func (o LogDataProtectionPolicyMapOutput) ToLogDataProtectionPolicyMapOutputWithContext(ctx context.Context) LogDataProtectionPolicyMapOutput {
	return o
}

func (o LogDataProtectionPolicyMapOutput) MapIndex(k pulumi.StringInput) LogDataProtectionPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogDataProtectionPolicy {
		return vs[0].(map[string]*LogDataProtectionPolicy)[vs[1].(string)]
	}).(LogDataProtectionPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogDataProtectionPolicyInput)(nil)).Elem(), &LogDataProtectionPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogDataProtectionPolicyArrayInput)(nil)).Elem(), LogDataProtectionPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogDataProtectionPolicyMapInput)(nil)).Elem(), LogDataProtectionPolicyMap{})
	pulumi.RegisterOutputType(LogDataProtectionPolicyOutput{})
	pulumi.RegisterOutputType(LogDataProtectionPolicyArrayOutput{})
	pulumi.RegisterOutputType(LogDataProtectionPolicyMapOutput{})
}
