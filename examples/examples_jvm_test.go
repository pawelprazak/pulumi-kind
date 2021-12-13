package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestClusterJVMAlone(t *testing.T) {
	test := getJVMBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "cluster-jvm"),
		})

	integration.ProgramTest(t, &test)
}

// TODO
/*func TestClusterJVMDeployment(t *testing.T) {
	test := getJVMBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "cluster-deployment-jvm"),
		})

	integration.ProgramTest(t, &test)
}*/

func getJVMBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseJVM := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"io.pulumi:pulumi:3.6.0+",
		},
	})

	return baseJVM
}
