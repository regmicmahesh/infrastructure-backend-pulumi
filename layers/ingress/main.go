package main

import (
	label "github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/label"
	loadbalancer "github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/load-balancer"
	s3cloudfrontwebsite "github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/s3-cloudfront-website"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {

	pulumi.Run(func(ctx *pulumi.Context) error {

		domainConf := config.New(ctx, "domain")

		networkRefName := fmt.Sprintf("organization/network/%s", ctx.Stack())
		networkRef, err := pulumi.NewStackReference(ctx, networkRefName, nil)
		if err != nil {
			return err
		}

		certsRefName := fmt.Sprintf("organization/certs/%s", ctx.Stack())
		certsRef, err := pulumi.NewStackReference(ctx, certsRefName, nil)
		if err != nil {
			return err
		}

		assetsRefName := fmt.Sprintf("organization/assets/%s", ctx.Stack())
		assetsRef, err := pulumi.NewStackReference(ctx, assetsRefName, nil)
		if err != nil {
			return err
		}

		loadBalancerResourceName := label.ID(ctx, "load-balancer")
		lb, err := loadbalancer.NewLoadBalancer(ctx, loadBalancerResourceName, &loadbalancer.LoadBalancerArgs{
			Internal:          pulumi.Bool(false),
			SubnetIds:         networkRef.GetOutput(pulumi.String("publicSubnetIds")).AsStringArrayOutput(),
			VpcCidr:           networkRef.GetStringOutput(pulumi.String("vpcCidr")),
			VpcId:             networkRef.GetStringOutput(pulumi.String("vpcId")),
			AcmCertificateArn: certsRef.GetStringOutput(pulumi.String("wildCardCertificateArn")),
			VpcSecurityGroupIds: pulumi.StringArray{
				networkRef.GetStringOutput(pulumi.String("loadBalancerSecurityGroupId")),
			},
		})
		if err != nil {
			return err
		}

		graphServiceTargetGroupResourceName := label.ID(ctx, "graph-service")
		graphServiceTargetGroup, err := loadbalancer.NewTargetGroup(ctx, graphServiceTargetGroupResourceName, &loadbalancer.TargetGroupArgs{
			VpcId:       networkRef.GetStringOutput(pulumi.String("vpcId")),
			ListenerArn: lb.HttpsListenerArn,
			DomainName:  pulumi.String(domainConf.Require("graph")),
		})
		apiServiceTargetGroupResourceName := label.ID(ctx, "api-service")
		apiServiceTargetGroup, err := loadbalancer.NewTargetGroup(ctx, apiServiceTargetGroupResourceName, &loadbalancer.TargetGroupArgs{
			VpcId:       networkRef.GetStringOutput(pulumi.String("vpcId")),
			ListenerArn: lb.HttpsListenerArn,
			DomainName:  pulumi.String(domainConf.Require("api")),
		})

		internalLoadBalancerResourceName := label.ID(ctx, "internal-load-balancer")
		ilb, err := loadbalancer.NewLoadBalancer(ctx, internalLoadBalancerResourceName, &loadbalancer.LoadBalancerArgs{
			Internal:          pulumi.Bool(true),
			SubnetIds:         networkRef.GetOutput(pulumi.String("privateSubnetIds")).AsStringArrayOutput(),
			VpcCidr:           networkRef.GetStringOutput(pulumi.String("vpcCidr")),
			VpcId:             networkRef.GetStringOutput(pulumi.String("vpcId")),
			AcmCertificateArn: certsRef.GetStringOutput(pulumi.String("wildCardCertificateArn")),
			VpcSecurityGroupIds: pulumi.StringArray{
				networkRef.GetStringOutput(pulumi.String("internalLoadBalancerSecurityGroupId")),
			},
		})

		graphRpcServiceTargetGroupResourceName := label.ID(ctx, "graph-rpc-service")
		graphRpcServiceTargetGroup, err := loadbalancer.NewTargetGroup(ctx, graphRpcServiceTargetGroupResourceName, &loadbalancer.TargetGroupArgs{
			VpcId:       networkRef.GetStringOutput(pulumi.String("vpcId")),
			ListenerArn: ilb.HttpsListenerArn,
			DomainName:  pulumi.String(domainConf.Require("graph-rpc")),
		})

		ipfsServiceTargetGroupResourceName := label.ID(ctx, "ipfs-service")
		ipfsServiceTargetGroup, err := loadbalancer.NewTargetGroup(ctx, ipfsServiceTargetGroupResourceName, &loadbalancer.TargetGroupArgs{
			VpcId:       networkRef.GetStringOutput(pulumi.String("vpcId")),
			ListenerArn: ilb.HttpsListenerArn,
			DomainName:  pulumi.String(domainConf.Require("ipfs")),
		})

		appWebsiteResourceName := label.ID(ctx, "app-website")
		appWebsite, err := s3cloudfrontwebsite.NewS3CloudfrontWebsite(ctx, appWebsiteResourceName, &s3cloudfrontwebsite.S3CloudfrontWebsiteArgs{
			DomainName:               pulumi.String(domainConf.Get("app")),
			BucketRegionalDomainName: assetsRef.GetStringOutput(pulumi.String("appBucketRegionalDomainName")),
			BucketName:               assetsRef.GetStringOutput(pulumi.String("appBucketName")),
			BucketArn:                assetsRef.GetStringOutput(pulumi.String("appBucketArn")),
			CertificateArn:           certsRef.GetStringOutput(pulumi.String("wildCardCertificateArn")),
		})

		landingWebsiteResourceName := label.ID(ctx, "landing-website")
		landingWebsite, err := s3cloudfrontwebsite.NewS3CloudfrontWebsite(ctx, landingWebsiteResourceName, &s3cloudfrontwebsite.S3CloudfrontWebsiteArgs{
			DomainName:               pulumi.String(domainConf.Get("landing")),
			BucketRegionalDomainName: assetsRef.GetStringOutput(pulumi.String("landingBucketRegionalDomainName")),
			BucketName:               assetsRef.GetStringOutput(pulumi.String("landingBucketName")),
			BucketArn:                assetsRef.GetStringOutput(pulumi.String("landingBucketArn")),
			CertificateArn:           certsRef.GetStringOutput(pulumi.String("rootCertificateArn")),
		})

		rootDomain := domainConf.Require("root")
		rootZone, err := route53.LookupZone(ctx, &route53.LookupZoneArgs{
			Name:        &rootDomain,
			PrivateZone: pulumi.BoolRef(false),
		})

		if err != nil {
			panic(err)
		}

		_, err = route53.NewRecord(ctx, appWebsiteResourceName, &route53.RecordArgs{
			Type:   pulumi.String("A"),
			ZoneId: pulumi.String(rootZone.ZoneId),
			Aliases: route53.RecordAliasArray{
				&route53.RecordAliasArgs{
					Name:                 appWebsite.Distribution.DomainName(),
					ZoneId:               appWebsite.Distribution.HostedZoneId(),
					EvaluateTargetHealth: pulumi.Bool(false),
				},
			},
			Name: pulumi.String(domainConf.Require("app")),
		})

		if err != nil {
			return nil
		}

		_, err = route53.NewRecord(ctx, landingWebsiteResourceName, &route53.RecordArgs{
			Type:   pulumi.String("A"),
			ZoneId: pulumi.String(rootZone.ZoneId),
			Aliases: route53.RecordAliasArray{
				&route53.RecordAliasArgs{
					Name:                 landingWebsite.Distribution.DomainName(),
					ZoneId:               landingWebsite.Distribution.HostedZoneId(),
					EvaluateTargetHealth: pulumi.Bool(false),
				},
			},
			Name: pulumi.String(domainConf.Require("landing")),
		})

		if err != nil {
			return nil
		}

		_, err = route53.NewRecord(ctx, graphServiceTargetGroupResourceName, &route53.RecordArgs{
			Type:   pulumi.String("CNAME"),
			ZoneId: pulumi.String(rootZone.ZoneId),
			Ttl:    pulumi.Int(300),
			Records: pulumi.StringArray{
				lb.DnsName,
			},
			Name: pulumi.String(domainConf.Require("graph")),
		})

		_, err = route53.NewRecord(ctx, apiServiceTargetGroupResourceName, &route53.RecordArgs{
			Type:   pulumi.String("CNAME"),
			ZoneId: pulumi.String(rootZone.ZoneId),
			Ttl:    pulumi.Int(300),
			Records: pulumi.StringArray{
				lb.DnsName,
			},
			Name: pulumi.String(domainConf.Require("api")),
		})

		_, err = route53.NewRecord(ctx, ipfsServiceTargetGroupResourceName, &route53.RecordArgs{
			Type:   pulumi.String("CNAME"),
			ZoneId: pulumi.String(rootZone.ZoneId),
			Ttl:    pulumi.Int(300),
			Records: pulumi.StringArray{
				ilb.DnsName,
			},
			Name: pulumi.String(domainConf.Require("ipfs")),
		})

		_, err = route53.NewRecord(ctx, graphRpcServiceTargetGroupResourceName, &route53.RecordArgs{
			Type:   pulumi.String("CNAME"),
			ZoneId: pulumi.String(rootZone.ZoneId),
			Ttl:    pulumi.Int(300),
			Records: pulumi.StringArray{
				ilb.DnsName,
			},
			Name: pulumi.String(domainConf.Require("graph-rpc")),
		})

		if err != nil {
			return nil
		}

		ctx.Export("loadBalancerArn", lb.LoadBalancer.Arn())
		ctx.Export("internalLoadBalancerArn", ilb.LoadBalancer.Arn())

		ctx.Export("graphServiceTargetGroupArn", graphServiceTargetGroup.TargetGroup.Arn())
		ctx.Export("apiServiceTargetGroupArn", apiServiceTargetGroup.TargetGroup.Arn())
		ctx.Export("ipfsServiceTargetGroupArn", ipfsServiceTargetGroup.TargetGroup.Arn())
		ctx.Export("graphRpcServiceTargetGroupArn", graphRpcServiceTargetGroup.TargetGroup.Arn())

		if err != nil {
			return err
		}

		return nil
	})
}
