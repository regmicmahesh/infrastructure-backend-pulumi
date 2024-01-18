package label

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ID(ctx *pulumi.Context, resourceName string) string {
	stackName := ctx.Stack()

	// already has stack-name attached in the name.
	prefix := fmt.Sprintf("%s-", stackName)
	if strings.HasPrefix(resourceName, prefix) {
		return resourceName
	}

	name := fmt.Sprintf("%s-%s", stackName, resourceName)
	return name
}

func Tags(ctx *pulumi.Context, name string) pulumi.StringMap {

	return pulumi.StringMap{
		"Name":        pulumi.String(ID(ctx, name)),
		"environment": pulumi.String(ctx.Stack()),
		"project":     pulumi.String(ctx.Project()),
	}

}
