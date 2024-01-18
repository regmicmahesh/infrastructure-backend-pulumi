package ssmparameters

import (
	"github.com/regmicmahesh/infrastructure-backend-pulumi/internal/file"
	"github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/label"
	"fmt"
	"slices"
	"strings"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SSMParameters struct {
	pulumi.ResourceState
	CertificateArn  pulumi.StringOutput `pulumi:"CertificateArn"`
	VariableArns    pulumi.StringArray
	SecretArns      pulumi.StringArray
	AccessPolicyArn pulumi.StringOutput `pulumi:"accessPolicyArn"`
}

type SSMParametersArgs struct {
	Variables map[string]pulumi.StringInput
	Secrets   map[string]pulumi.StringInput
}

func NewSSMParameters(ctx *pulumi.Context, name string, args *SSMParametersArgs, opts ...pulumi.ResourceOption) (*SSMParameters, error) {

	tags := label.Tags(ctx, name)

	s := &SSMParameters{}
	err := ctx.RegisterComponentResource("pkg:index:ssm-parameters", name, s, opts...)
	if err != nil {
		return nil, err
	}

	policyTemplate := file.ReadFileOrPanic("../../policies/ssm-parameter.access.json")
	tmpl, err := file.Template.Parse(string(policyTemplate))
	if err != nil {
		panic(err)
	}

	current, err := aws.GetCallerIdentity(ctx, nil, pulumi.Parent(s))

	var sb strings.Builder
	err = tmpl.Execute(&sb, &struct {
		ParameterPrefix string
		AccountID       string
	}{
		ParameterPrefix: name,
		AccountID:       current.AccountId,
	})
	if err != nil {
		return nil, err
	}

	iamAccessPolicy, err := iam.NewPolicy(ctx, name, &iam.PolicyArgs{
		Name:   pulumi.String(name),
		Policy: pulumi.String(sb.String()),
		Tags:   tags,
	}, pulumi.Parent(s))

	if err != nil {
		return nil, err
	}

	s.AccessPolicyArn = iamAccessPolicy.Arn

	variableKeys := make([]string, 0, len(args.Variables))
	for k := range args.Variables {
		variableKeys = append(variableKeys, k)
	}
	slices.Sort(variableKeys)

	for _, k := range variableKeys {

		key := fmt.Sprintf("/%s/%s", name, k)
		param, err := ssm.NewParameter(ctx, key, &ssm.ParameterArgs{
			Type:      pulumi.String("String"),
			Name:      pulumi.String(key),
			Value:     args.Variables[k],
			Tags:      tags,
			//Overwrite: pulumi.Bool(true),
		}, pulumi.Parent(s))
		if err != nil {
			return nil, err
		}
		s.VariableArns = append(s.VariableArns, param.Arn)

	}

	secretKeys := make([]string, 0, len(args.Secrets))
	for k := range args.Secrets {
		secretKeys = append(secretKeys, k)
	}
	slices.Sort(secretKeys)

	for _, k := range secretKeys {

		key := fmt.Sprintf("/%s/%s", name, k)
		param, err := ssm.NewParameter(ctx, key, &ssm.ParameterArgs{
			Type:      pulumi.String("SecureString"),
			Name:      pulumi.String(key),
			Value:     args.Secrets[k],
			//Overwrite: pulumi.Bool(true),
			Tags:      tags,
		}, pulumi.Parent(s))
		if err != nil {
			return nil, err
		}
		s.SecretArns = append(s.SecretArns, param.Arn)

	}

	ctx.RegisterResourceOutputs(s, pulumi.Map{})

	return s, nil

}
