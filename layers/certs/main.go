package main

import (
	acmcertificate "github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/acm-certificate"
	label "github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/label"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		conf := config.New(ctx, "")

		wildCardCertResourceName := label.ID(ctx, "wildcard-cert")
		wildCardCertificate, err := acmcertificate.NewAcmCertificate(ctx, wildCardCertResourceName, &acmcertificate.AcmCertificateArgs{
			DomainName:     pulumi.String(conf.Require("wildcard-domain")),
			RootDomainName: conf.Require("root-domain"),
		})
		if err != nil {
			return err
		}

		rootCertResourceName := label.ID(ctx, "root-cert")
		rootCertificate, err := acmcertificate.NewAcmCertificate(ctx, rootCertResourceName, &acmcertificate.AcmCertificateArgs{
			DomainName:     pulumi.String(conf.Require("root-domain")),
			RootDomainName: conf.Require("root-domain"),
		})
		if err != nil {
			return err
		}

		ctx.Export("rootCertificateArn", rootCertificate.CertificateArn)
		ctx.Export("wildCardCertificateArn", wildCardCertificate.CertificateArn)

		return nil

	})
}
