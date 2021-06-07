// python tests

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestClusterPythonAlone(t *testing.T) {
	test := getPyBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "clusterPy"),
		})

	integration.ProgramTest(t, &test)
}

func TestClusterPythonDeployment(t *testing.T) {
	test := getPyBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "clusterDeploymentPy"),
		})

	integration.ProgramTest(t, &test)
}

func getPyBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	basePython := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			path.Join("..", "sdk", "python", "bin"),
		},
	})

	return basePython
}
