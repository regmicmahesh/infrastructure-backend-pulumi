// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws:cloudwatch/compositeAlarm:CompositeAlarm":
		r = &CompositeAlarm{}
	case "aws:cloudwatch/dashboard:Dashboard":
		r = &Dashboard{}
	case "aws:cloudwatch/eventApiDestination:EventApiDestination":
		r = &EventApiDestination{}
	case "aws:cloudwatch/eventArchive:EventArchive":
		r = &EventArchive{}
	case "aws:cloudwatch/eventBus:EventBus":
		r = &EventBus{}
	case "aws:cloudwatch/eventBusPolicy:EventBusPolicy":
		r = &EventBusPolicy{}
	case "aws:cloudwatch/eventConnection:EventConnection":
		r = &EventConnection{}
	case "aws:cloudwatch/eventEndpoint:EventEndpoint":
		r = &EventEndpoint{}
	case "aws:cloudwatch/eventPermission:EventPermission":
		r = &EventPermission{}
	case "aws:cloudwatch/eventRule:EventRule":
		r = &EventRule{}
	case "aws:cloudwatch/eventTarget:EventTarget":
		r = &EventTarget{}
	case "aws:cloudwatch/internetMonitor:InternetMonitor":
		r = &InternetMonitor{}
	case "aws:cloudwatch/logDataProtectionPolicy:LogDataProtectionPolicy":
		r = &LogDataProtectionPolicy{}
	case "aws:cloudwatch/logDestination:LogDestination":
		r = &LogDestination{}
	case "aws:cloudwatch/logDestinationPolicy:LogDestinationPolicy":
		r = &LogDestinationPolicy{}
	case "aws:cloudwatch/logGroup:LogGroup":
		r = &LogGroup{}
	case "aws:cloudwatch/logMetricFilter:LogMetricFilter":
		r = &LogMetricFilter{}
	case "aws:cloudwatch/logResourcePolicy:LogResourcePolicy":
		r = &LogResourcePolicy{}
	case "aws:cloudwatch/logStream:LogStream":
		r = &LogStream{}
	case "aws:cloudwatch/logSubscriptionFilter:LogSubscriptionFilter":
		r = &LogSubscriptionFilter{}
	case "aws:cloudwatch/metricAlarm:MetricAlarm":
		r = &MetricAlarm{}
	case "aws:cloudwatch/metricStream:MetricStream":
		r = &MetricStream{}
	case "aws:cloudwatch/queryDefinition:QueryDefinition":
		r = &QueryDefinition{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/compositeAlarm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/dashboard",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/eventApiDestination",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/eventArchive",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/eventBus",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/eventBusPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/eventConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/eventEndpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/eventPermission",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/eventRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/eventTarget",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/internetMonitor",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/logDataProtectionPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/logDestination",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/logDestinationPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/logGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/logMetricFilter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/logResourcePolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/logStream",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/logSubscriptionFilter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/metricAlarm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/metricStream",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudwatch/queryDefinition",
		&module{version},
	)
}
