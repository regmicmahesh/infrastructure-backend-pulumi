package s3cloudfrontwebsite

import (
	"fmt"

	"github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/label"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type S3CloudfrontWebsite struct {
	pulumi.ResourceState

	Distribution cloudfront.DistributionOutput
	SourceBucket s3.BucketV2Output
	DnsRecord    route53.RecordOutput
}

type S3CloudfrontWebsiteArgs struct {
	CertificateArn           pulumi.StringInput
	BucketRegionalDomainName pulumi.StringInput
	DomainName               pulumi.StringInput

	BucketName pulumi.StringInput
	BucketArn  pulumi.StringInput
}

func NewS3CloudfrontWebsite(ctx *pulumi.Context, name string, args *S3CloudfrontWebsiteArgs, opts ...pulumi.ResourceOption) (*S3CloudfrontWebsite, error) {
	tags := label.Tags(ctx, name)

	s := &S3CloudfrontWebsite{}
	err := ctx.RegisterComponentResource("pkg:index:s3-cloudfront-website", name, s, opts...)
	if err != nil {
		return nil, err
	}

	appDistributionAccessControl, err := cloudfront.NewOriginAccessControl(ctx, name, &cloudfront.OriginAccessControlArgs{
		Description:                   pulumi.String("Allow access to S3 buckets"),
		OriginAccessControlOriginType: pulumi.String("s3"),
		SigningBehavior:               pulumi.String("always"),
		SigningProtocol:               pulumi.String("sigv4"),
	}, pulumi.Parent(s))

	if err != nil {
		return nil, err
	}

	distribution, err := cloudfront.NewDistribution(ctx, name, &cloudfront.DistributionArgs{

		Origins: cloudfront.DistributionOriginArray{
			&cloudfront.DistributionOriginArgs{
				DomainName:            args.BucketRegionalDomainName,
				OriginId:              args.BucketRegionalDomainName,
				OriginAccessControlId: appDistributionAccessControl.ID(),
			},
		},
		Enabled:           pulumi.Bool(true),
		IsIpv6Enabled:     pulumi.Bool(true),
		Comment:           pulumi.String(name),
		DefaultRootObject: pulumi.String("index.html"),
		Aliases: pulumi.StringArray{
			args.DomainName,
		},

		DefaultCacheBehavior: cloudfront.DistributionDefaultCacheBehaviorArgs{
			CachePolicyId:         pulumi.String("658327ea-f89d-4fab-a63d-7e88639e58f6"), // Managed-CachingOptimized
			OriginRequestPolicyId: pulumi.String("b689b0a8-53d0-40ab-baf2-68738e2966ac"), // Managed-AllViewerExceptHostHeader
			ViewerProtocolPolicy:  pulumi.String("redirect-to-https"),
			AllowedMethods: pulumi.StringArray{
				pulumi.String("DELETE"),
				pulumi.String("GET"),
				pulumi.String("HEAD"),
				pulumi.String("OPTIONS"),
				pulumi.String("PATCH"),
				pulumi.String("POST"),
				pulumi.String("PUT"),
			},
			CachedMethods: pulumi.StringArray{
				pulumi.String("GET"),
				pulumi.String("HEAD"),
			},
			TargetOriginId: args.BucketRegionalDomainName,
		},
		Restrictions: cloudfront.DistributionRestrictionsArgs{
			GeoRestriction: &cloudfront.DistributionRestrictionsGeoRestrictionArgs{
				RestrictionType: pulumi.String("none"),
			},
		},
		CustomErrorResponses: cloudfront.DistributionCustomErrorResponseArray{
			cloudfront.DistributionCustomErrorResponseArgs{
				ErrorCachingMinTtl: pulumi.Int(300),
				ErrorCode:          pulumi.Int(403),
				ResponseCode:       pulumi.Int(200),
				ResponsePagePath:   pulumi.String("/index.html"),
			},
		},
		ViewerCertificate: &cloudfront.DistributionViewerCertificateArgs{
			AcmCertificateArn: args.CertificateArn,
			SslSupportMethod:  pulumi.String("sni-only"),
		},
		Tags: tags,
	}, pulumi.Parent(s))

	if err != nil {
		return nil, err
	}

	s.Distribution = distribution.ToDistributionOutput()

	_, err = s3.NewBucketPolicy(ctx, name, &s3.BucketPolicyArgs{

		Bucket: args.BucketName,
		Policy: pulumix.Apply2(distribution.Arn, args.BucketArn.ToStringOutput(), func(dAarn string, sArn string) string {
			res, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{

				Statements: []iam.GetPolicyDocumentStatement{
					{
						Principals: []iam.GetPolicyDocumentStatementPrincipal{
							{
								Identifiers: []string{"cloudfront.amazonaws.com"},
								Type:        "Service",
							},
						},
						Actions:   []string{"s3:GetObject"},
						Resources: []string{fmt.Sprintf("%s/*", sArn)},
						Conditions: []iam.GetPolicyDocumentStatementCondition{
							{
								Test:     "StringEquals",
								Variable: "AWS:SourceArn",
								Values:   []string{dAarn},
							},
						},
					},
				},
			})

			if err != nil {
				panic(err)
			}

			return res.Json
		}),
	}, pulumi.Parent(s))

	if err != nil {
		return nil, err
	}

	return s, nil

}
