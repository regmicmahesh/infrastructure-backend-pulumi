package main

import (
	ecrrepository "github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/ecr-repository"
	label "github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/label"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {

	pulumi.Run(func(ctx *pulumi.Context) error {

		ecrRepositoryResourceName := label.ID(ctx, "repository")
		_, err := ecrrepository.NewEcrRepository(ctx, ecrRepositoryResourceName, &ecrrepository.EcrRepositoryArgs{})
		if err != nil {
			return err
		}

		appBucketResourceName := label.ID(ctx, "app-website")
		appBucket, err := s3.NewBucketV2(ctx, appBucketResourceName, &s3.BucketV2Args{
			BucketPrefix: pulumi.StringPtr(appBucketResourceName),
			Tags:         label.Tags(ctx, appBucketResourceName),
		})

		if err != nil {
			return nil
		}

		_, err = s3.NewBucketObject(ctx, appBucketResourceName, &s3.BucketObjectArgs{
			Key:         pulumi.String("index.html"),
			Bucket:      appBucket.ID(),
			Source:      pulumi.NewStringAsset("<h1>Hello World</h1>"),
			ContentType: pulumi.String("text/html"),
			Tags:        label.Tags(ctx, appBucketResourceName),
		}, pulumi.IgnoreChanges([]string{"Source", "Tags"}))

		if err != nil {
			return nil
		}
		landingBucketResourceName := label.ID(ctx, "landing-website")
		landingBucket, err := s3.NewBucketV2(ctx, landingBucketResourceName, &s3.BucketV2Args{
			BucketPrefix: pulumi.StringPtr(landingBucketResourceName),
			Tags:         label.Tags(ctx, landingBucketResourceName),
		})

		if err != nil {
			return nil
		}

		_, err = s3.NewBucketObject(ctx, landingBucketResourceName, &s3.BucketObjectArgs{
			Key:         pulumi.String("index.html"),
			Bucket:      landingBucket.ID(),
			Source:      pulumi.NewStringAsset("<h1>Hello World</h1>"),
			ContentType: pulumi.String("text/html"),
			Tags:        label.Tags(ctx, landingBucketResourceName),
		}, pulumi.IgnoreChanges([]string{"Source", "Tags"}))

		if err != nil {
			return nil
		}

		ctx.Export("appBucketName", appBucket.Bucket)
		ctx.Export("appBucketArn", appBucket.Arn)
		ctx.Export("appBucketRegionalDomainName", appBucket.BucketRegionalDomainName)

		ctx.Export("landingBucketName", landingBucket.Bucket)
		ctx.Export("landingBucketArn", landingBucket.Arn)
		ctx.Export("landingBucketRegionalDomainName", landingBucket.BucketRegionalDomainName)

		return nil
	})
}
