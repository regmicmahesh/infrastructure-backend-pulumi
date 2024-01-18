package loadbalancer

import (
	"github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/label"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TargetGroup struct {
	pulumi.ResourceState
	TargetGroup lb.TargetGroupOutput
}

type TargetGroupArgs struct {
	VpcId       pulumi.StringInput
	ListenerArn pulumi.StringInput
	DomainName  pulumi.StringInput
}

func NewTargetGroup(ctx *pulumi.Context, name string, args *TargetGroupArgs, opts ...pulumi.ResourceOption) (*TargetGroup, error) {

	tags := label.Tags(ctx, name)

	t := &TargetGroup{}

	err := ctx.RegisterComponentResource("pkg:index:target-group", name, t, opts...)
	if err != nil {
		return nil, err
	}

	tg, err := lb.NewTargetGroup(ctx, name, &lb.TargetGroupArgs{
		Name:       pulumi.String(name),
		Port:       pulumi.Int(80),
		Protocol:   pulumi.String("HTTP"),
		TargetType: pulumi.String("instance"),
		VpcId:      args.VpcId,
		HealthCheck: &lb.TargetGroupHealthCheckArgs{
			Matcher: pulumi.String("200-499"),
		},
		Tags: tags,
	}, pulumi.Parent(t))
	if err != nil {
		return nil, err
	}

	t.TargetGroup = tg.ToTargetGroupOutput()

	_, err = lb.NewListenerRule(ctx, name, &lb.ListenerRuleArgs{
		ListenerArn: args.ListenerArn,
		Actions: lb.ListenerRuleActionArray{
			&lb.ListenerRuleActionArgs{
				Type:           pulumi.String("forward"),
				TargetGroupArn: tg.Arn,
			},
		},
		Conditions: lb.ListenerRuleConditionArray{
			&lb.ListenerRuleConditionArgs{
				HostHeader: &lb.ListenerRuleConditionHostHeaderArgs{
					Values: pulumi.StringArray{
						args.DomainName,
					},
				},
			},
		},
		Tags: tags,
	}, pulumi.Parent(t))
	if err != nil {
		return nil, err
	}

	return t, nil

}
