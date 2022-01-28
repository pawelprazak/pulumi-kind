// python tests

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestClusterDotnetAlone(t *testing.T) {
	test := getDontNetBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "cluster-dotnet"),
		})

	integration.ProgramTest(t, &test)
}

func TestClusterDotnetDeployment(t *testing.T) {
	test := getDontNetBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "cluster-deployment-dotnet"),
		})

	integration.ProgramTest(t, &test)
}

func getDontNetBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseDotnet := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"Pulumi.Kind",
		},
	})

	return baseDotnet
}
