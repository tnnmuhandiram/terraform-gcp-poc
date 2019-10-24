package terraform

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Destroy(options *Options) string {
	out, err := DestroyE(options)
	print(err)
	return out
}

func TgDestroyAll(t *testing.T, options *Options) string {
	out, err := TgDestroyAllE(t, options)
	require.NoError(t, err)
	return out
}

func DestroyE(options *Options) (string, error) {
	return RunTerraformCommandE(options, FormatArgs(options, "destroy", "-auto-approve", "-input=false", "-lock=false")...)
}

func TgDestroyAllE(t *testing.T, options *Options) (string, error) {
	if options.TerraformBinary != "terragrunt" {
		return "", TgInvalidBinary(options.TerraformBinary)
	}

	return RunTerraformCommandE(options, FormatArgs(options, "destroy-all", "-force", "-input=false", "-lock=false")...)
}
