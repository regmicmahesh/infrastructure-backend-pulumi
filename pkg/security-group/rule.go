package securitygroup

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IngressCIDRRuleArgs struct {
	IngressCIDRArgs
	VpcId           pulumi.StringInput
	SecurityGroupId pulumi.StringInput
}

type IngressCIDRRule struct {
	pulumi.ResourceState
}

func NewIngressCIDRRule(ctx *pulumi.Context, name string, args *IngressCIDRRuleArgs, opts ...pulumi.ResourceOption) (*IngressCIDRRule, error) {

	a := &IngressCIDRRule{}

	err := ctx.RegisterComponentResource("pkg:index:security-group-ingress-cidr", name, a, opts...)
	if err != nil {
		return nil, err
	}

	v := args.IngressCIDRArgs

	_, err = ec2.NewSecurityGroupRule(ctx, name, &ec2.SecurityGroupRuleArgs{
		SecurityGroupId: args.SecurityGroupId,
		CidrBlocks:      v.CidrBlocks,

		Type:     pulumi.String("ingress"),
		Protocol: pulumi.String(v.Protocol),
		FromPort: pulumi.Int(v.FromPort),
		ToPort:   pulumi.Int(v.ToPort),
	})
	if err != nil {
		return nil, err
	}

	return a, nil

}

type IngressSecurityGroupRuleArgs struct {
	IngressSecurityGroupArgs
	VpcId           pulumi.StringInput
	SecurityGroupId pulumi.StringInput
}

type IngressSecurityGroupRule struct {
	pulumi.ResourceState
}

func NewIngressSecurityGroupRule(ctx *pulumi.Context, name string, args *IngressSecurityGroupRuleArgs, opts ...pulumi.ResourceOption) (*IngressSecurityGroupRule, error) {

	a := &IngressSecurityGroupRule{}

	err := ctx.RegisterComponentResource("pkg:index:security-group-ingress-security-group", name, a, opts...)
	if err != nil {
		return nil, err
	}

	v := args.IngressSecurityGroupArgs

	_, err = ec2.NewSecurityGroupRule(ctx, name, &ec2.SecurityGroupRuleArgs{
		SecurityGroupId:       args.SecurityGroupId,
		SourceSecurityGroupId: v.SecurityGroupId,

		Type:     pulumi.String("ingress"),
		Protocol: pulumi.String(v.Protocol),
		FromPort: pulumi.Int(v.FromPort),
		ToPort:   pulumi.Int(v.ToPort),
	})
	if err != nil {
		return nil, err
	}

	return a, nil

}

type EgressCIDRRuleArgs struct {
	EgressCIDRArgs
	VpcId           pulumi.StringInput
	SecurityGroupId pulumi.StringInput
}

type EgressCIDRRule struct {
	pulumi.ResourceState
}

func NewEgressCIDRRule(ctx *pulumi.Context, name string, args *EgressCIDRRuleArgs, opts ...pulumi.ResourceOption) (*EgressCIDRRule, error) {

	a := &EgressCIDRRule{}

	err := ctx.RegisterComponentResource("pkg:index:security-group-egress-cidr", name, a, opts...)
	if err != nil {
		return nil, err
	}

	v := args.EgressCIDRArgs

	_, err = ec2.NewSecurityGroupRule(ctx, name, &ec2.SecurityGroupRuleArgs{
		SecurityGroupId: args.SecurityGroupId,
		CidrBlocks:      v.CidrBlocks,

		Type:     pulumi.String("egress"),
		Protocol: pulumi.String(v.Protocol),
		FromPort: pulumi.Int(v.FromPort),
		ToPort:   pulumi.Int(v.ToPort),
	})
	if err != nil {
		return nil, err
	}

	return a, nil

}

type EgressSecurityGroupRuleArgs struct {
	EgressSecurityGroupArgs
	VpcId           pulumi.StringInput
	SecurityGroupId pulumi.StringInput
}

type EgressSecurityGroupRule struct {
	pulumi.ResourceState
}

func NewEgressSecurityGroupRule(ctx *pulumi.Context, name string, args *EgressSecurityGroupRuleArgs, opts ...pulumi.ResourceOption) (*EgressSecurityGroupRule, error) {

	a := &EgressSecurityGroupRule{}

	err := ctx.RegisterComponentResource("pkg:index:security-group-egress-security-group", name, a, opts...)
	if err != nil {
		return nil, err
	}

	v := args.EgressSecurityGroupArgs

	_, err = ec2.NewSecurityGroupRule(ctx, name, &ec2.SecurityGroupRuleArgs{
		SecurityGroupId:       args.SecurityGroupId,
		SourceSecurityGroupId: v.SecurityGroupId,

		Type:     pulumi.String("egress"),
		Protocol: pulumi.String(v.Protocol),
		FromPort: pulumi.Int(v.FromPort),
		ToPort:   pulumi.Int(v.ToPort),
	})
	if err != nil {
		return nil, err
	}

	return a, nil

}
