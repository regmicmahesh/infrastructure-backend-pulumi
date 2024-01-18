package loadbalancer

import (
	"github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/label"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LoadBalancer struct {
	pulumi.ResourceState

	LoadBalancer     lb.LoadBalancerOutput
	HttpListenerArn  pulumi.StringOutput
	HttpsListenerArn pulumi.StringOutput

	DnsName pulumi.StringOutput
	ZoneId  pulumi.StringOutput
}

type LoadBalancerArgs struct {
	SubnetIds           pulumi.StringArrayInput
	VpcCidr             pulumi.StringInput
	VpcId               pulumi.StringInput
	AcmCertificateArn   pulumi.StringInput
	VpcSecurityGroupIds pulumi.StringArray

	Internal pulumi.Bool
}

func NewLoadBalancer(ctx *pulumi.Context, name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {

	tags := label.Tags(ctx, name)

	l := &LoadBalancer{}
	err := ctx.RegisterComponentResource("pkg:index:load-balancer", name, l, opts...)
	if err != nil {
		return nil, err
	}

	loadbalancer, err := lb.NewLoadBalancer(ctx, name, &lb.LoadBalancerArgs{
		Name:             pulumi.String(name),
		LoadBalancerType: pulumi.String("application"),
		Internal:         args.Internal,
		SecurityGroups:   args.VpcSecurityGroupIds,
		Subnets:          args.SubnetIds,
		Tags:             tags,
	}, pulumi.Parent(l))
	if err != nil {
		return nil, err
	}
	l.LoadBalancer = loadbalancer.ToLoadBalancerOutput()

	l.DnsName = loadbalancer.DnsName
	l.ZoneId = loadbalancer.ZoneId

	httpListener, err := lb.NewListener(ctx, fmt.Sprintf("%s-http", name), &lb.ListenerArgs{
		LoadBalancerArn: loadbalancer.Arn,
		Port:            pulumi.Int(80),
		Protocol:        pulumi.String("HTTP"),
		DefaultActions: lb.ListenerDefaultActionArray{
			&lb.ListenerDefaultActionArgs{
				Type: pulumi.String("redirect"),
				Redirect: &lb.ListenerDefaultActionRedirectArgs{
					Port:       pulumi.String("443"),
					Protocol:   pulumi.String("HTTPS"),
					StatusCode: pulumi.String("HTTP_301"),
				},
			},
		},
		Tags: tags,
	}, pulumi.Parent(l))
	if err != nil {
		return nil, err
	}
	l.HttpListenerArn = httpListener.Arn

	httpsListener, err := lb.NewListener(ctx, fmt.Sprintf("%s-https", name), &lb.ListenerArgs{
		LoadBalancerArn: loadbalancer.Arn,
		Port:            pulumi.Int(443),
		Protocol:        pulumi.String("HTTPS"),
		CertificateArn:  args.AcmCertificateArn,
		DefaultActions: lb.ListenerDefaultActionArray{
			&lb.ListenerDefaultActionArgs{
				Type: pulumi.String("fixed-response"),
				FixedResponse: &lb.ListenerDefaultActionFixedResponseArgs{
					ContentType: pulumi.String("text/plain"),
					MessageBody: pulumi.String("API is up and running."),
					StatusCode:  pulumi.String("200"),
				},
			},
		},
		Tags: tags,
	}, pulumi.Parent(l))
	if err != nil {
		return nil, err
	}
	l.HttpsListenerArn = httpsListener.Arn

	return l, nil

}
