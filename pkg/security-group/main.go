package securitygroup

import (
	"github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/label"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecurityGroup struct {
	pulumi.ResourceState

	SecurityGroup   ec2.SecurityGroupOutput
	SecurityGroupId pulumi.StringOutput
}

type IngressCIDRArgs struct {
	FromPort   int
	ToPort     int
	Protocol   string
	CidrBlocks pulumi.StringArrayInput
}

type IngressSecurityGroupArgs struct {
	FromPort        int
	ToPort          int
	Protocol        string
	SecurityGroupId pulumi.StringInput
}

type EgressCIDRArgs struct {
	FromPort   int
	ToPort     int
	Protocol   string
	CidrBlocks pulumi.StringArrayInput
}

type EgressSecurityGroupArgs struct {
	FromPort        int
	ToPort          int
	Protocol        string
	SecurityGroupId pulumi.StringInput
}

type SecurityGroupArgs struct {
	VpcId pulumi.StringInput

	IngressCIDRArgs          []IngressCIDRArgs
	IngressSecurityGroupArgs []IngressSecurityGroupArgs

	EgressCIDRArgs          []IngressCIDRArgs
	EgressSecurityGroupArgs []IngressSecurityGroupArgs
}

func NewSecurityGroup(ctx *pulumi.Context, name string, args *SecurityGroupArgs, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {

	tags := label.Tags(ctx, name)

	a := &SecurityGroup{}

	err := ctx.RegisterComponentResource("pkg:index:security-group", name, a, opts...)
	if err != nil {
		return nil, err
	}

	secGroup, err := ec2.NewSecurityGroup(ctx, name, &ec2.SecurityGroupArgs{
		Name:  pulumi.String(name),
		VpcId: args.VpcId,
		Tags:  tags,
	})

	if err != nil {
		return nil, err
	}

	a.SecurityGroupId = secGroup.ID().ToStringOutput()
	a.SecurityGroup = secGroup.ToSecurityGroupOutput()

	for i, v := range args.IngressCIDRArgs {
		resourceName := fmt.Sprintf("%s-ingress-cidr-%d", name, i)
		_, err = ec2.NewSecurityGroupRule(ctx, resourceName, &ec2.SecurityGroupRuleArgs{
			SecurityGroupId: secGroup.ID(),
			CidrBlocks:      v.CidrBlocks,

			Type:     pulumi.String("ingress"),
			Protocol: pulumi.String(v.Protocol),
			FromPort: pulumi.Int(v.FromPort),
			ToPort:   pulumi.Int(v.ToPort),
		})
		if err != nil {
			return nil, err
		}

	}

	for i, v := range args.IngressSecurityGroupArgs {
		resourceName := fmt.Sprintf("%s-ingress-security-group-%d", name, i)
		_, err = ec2.NewSecurityGroupRule(ctx, resourceName, &ec2.SecurityGroupRuleArgs{
			SecurityGroupId:       secGroup.ID(),
			SourceSecurityGroupId: v.SecurityGroupId,

			Type:     pulumi.String("ingress"),
			Protocol: pulumi.String(v.Protocol),
			FromPort: pulumi.Int(v.FromPort),
			ToPort:   pulumi.Int(v.ToPort),
		})
		if err != nil {
			return nil, err
		}

	}

	for i, v := range args.EgressCIDRArgs {
		resourceName := fmt.Sprintf("%s-egress-cidr-%d", name, i)
		_, err = ec2.NewSecurityGroupRule(ctx, resourceName, &ec2.SecurityGroupRuleArgs{
			SecurityGroupId: secGroup.ID(),
			CidrBlocks:      v.CidrBlocks,

			Type:     pulumi.String("egress"),
			Protocol: pulumi.String(v.Protocol),
			FromPort: pulumi.Int(v.FromPort),
			ToPort:   pulumi.Int(v.ToPort),
		})
		if err != nil {
			return nil, err
		}

	}

	for i, v := range args.EgressSecurityGroupArgs {
		resourceName := fmt.Sprintf("%s-egress-security-group-%d", name, i)
		_, err = ec2.NewSecurityGroupRule(ctx, resourceName, &ec2.SecurityGroupRuleArgs{
			SecurityGroupId:       secGroup.ID(),
			SourceSecurityGroupId: v.SecurityGroupId,

			Type:     pulumi.String("egress"),
			Protocol: pulumi.String(v.Protocol),
			FromPort: pulumi.Int(v.FromPort),
			ToPort:   pulumi.Int(v.ToPort),
		})
		if err != nil {
			return nil, err
		}

	}

	return a, nil

}
