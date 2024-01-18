package acmcertificate

import (
	"github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/label"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/acm"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AcmCertificate struct {
	pulumi.ResourceState
	CertificateArn pulumi.StringOutput `pulumi:"CertificateArn"`
}

type AcmCertificateArgs struct {
	DomainName     pulumi.StringInput
	RootDomainName string
}

func NewAcmCertificate(ctx *pulumi.Context, name string, args *AcmCertificateArgs, opts ...pulumi.ResourceOption) (*AcmCertificate, error) {

	tags := label.Tags(ctx, name)

	a := &AcmCertificate{}
	err := ctx.RegisterComponentResource("pkg:index:acm-certificate", name, a, opts...)
	if err != nil {
		return nil, err
	}

	cert, err := acm.NewCertificate(ctx, name, &acm.CertificateArgs{
		DomainName:       args.DomainName,
		ValidationMethod: pulumi.String("DNS"),
		Tags:             tags,
	}, pulumi.Parent(a))

	a.CertificateArn = cert.Arn

	rootZone, err := route53.LookupZone(ctx, &route53.LookupZoneArgs{
		Name:        pulumi.StringRef(args.RootDomainName),
		PrivateZone: pulumi.BoolRef(false),
	}, pulumi.Parent(a))

	domainValidationOption := cert.DomainValidationOptions.ApplyT(func(options []acm.CertificateDomainValidationOption) interface{} {
		return options[0]
	})

	validationRecord, err := route53.NewRecord(ctx, name, &route53.RecordArgs{
		Name: domainValidationOption.ApplyT(func(option interface{}) string {
			return *option.(acm.CertificateDomainValidationOption).ResourceRecordName
		}).(pulumi.StringOutput),
		Type: domainValidationOption.ApplyT(func(option interface{}) string {
			return *option.(acm.CertificateDomainValidationOption).ResourceRecordType
		}).(pulumi.StringOutput),
		Records: pulumi.StringArray{
			domainValidationOption.ApplyT(func(option interface{}) string {
				return *option.(acm.CertificateDomainValidationOption).ResourceRecordValue
			}).(pulumi.StringOutput),
		},
		Ttl:            pulumi.Int(10 * 60),
		ZoneId:         pulumi.String(rootZone.ZoneId),
		AllowOverwrite: pulumi.Bool(true),
	}, pulumi.Parent(a))

	_, err = acm.NewCertificateValidation(ctx, name, &acm.CertificateValidationArgs{
		CertificateArn: cert.Arn,
		ValidationRecordFqdns: pulumi.StringArray{
			validationRecord.Fqdn,
		},
	}, pulumi.Parent(a))

	if err != nil {
		return nil, err
	}

	ctx.RegisterResourceOutputs(a, pulumi.Map{
		"CertificateArn": cert.Arn,
	})

	return a, nil

}
