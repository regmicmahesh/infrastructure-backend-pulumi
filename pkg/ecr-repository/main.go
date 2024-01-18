package ecrrepository

import (
	"github.com/regmicmahesh/infrastructure-backend-pulumi/pkg/label"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EcrRepository struct {
	pulumi.ResourceState
	Repository ecr.RepositoryOutput
}

type EcrRepositoryArgs struct {
}

func NewEcrRepository(ctx *pulumi.Context, name string, args *EcrRepositoryArgs, opts ...pulumi.ResourceOption) (*EcrRepository, error) {

	tags := label.Tags(ctx, name)

	a := &EcrRepository{}
	err := ctx.RegisterComponentResource("pkg:index:ecr-repository", name, a, opts...)
	if err != nil {
		return nil, err
	}

	repo, err := ecr.NewRepository(ctx, name, &ecr.RepositoryArgs{
		Name: pulumi.String(name),
		Tags: tags,
	})
	a.Repository = repo.ToRepositoryOutput()

	return a, nil

}
