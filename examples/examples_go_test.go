// python tests

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestClusterGoAlone(t *testing.T) {
	test := getGoBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "cluster-go"),
		})

	integration.ProgramTest(t, &test)
}
func TestClusterGoDeploymet(t *testing.T) {
	test := getGoBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "cluster-deployment-go"),
		})

	integration.ProgramTest(t, &test)
}

func getGoBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseGo := base.With(integration.ProgramTestOptions{})

	return baseGo
}
