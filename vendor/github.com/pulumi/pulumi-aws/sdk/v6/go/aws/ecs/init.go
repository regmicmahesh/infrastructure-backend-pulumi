// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

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
	case "aws:ecs/accountSettingDefault:AccountSettingDefault":
		r = &AccountSettingDefault{}
	case "aws:ecs/capacityProvider:CapacityProvider":
		r = &CapacityProvider{}
	case "aws:ecs/cluster:Cluster":
		r = &Cluster{}
	case "aws:ecs/clusterCapacityProviders:ClusterCapacityProviders":
		r = &ClusterCapacityProviders{}
	case "aws:ecs/service:Service":
		r = &Service{}
	case "aws:ecs/tag:Tag":
		r = &Tag{}
	case "aws:ecs/taskDefinition:TaskDefinition":
		r = &TaskDefinition{}
	case "aws:ecs/taskSet:TaskSet":
		r = &TaskSet{}
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
		"ecs/accountSettingDefault",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ecs/capacityProvider",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ecs/cluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ecs/clusterCapacityProviders",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ecs/service",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ecs/tag",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ecs/taskDefinition",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ecs/taskSet",
		&module{version},
	)
}
