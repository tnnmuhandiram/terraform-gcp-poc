package terraform

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func InitAndApply(options *Options) string {
	out, err := InitAndApplyE(options)

	print(err)
	print(out)
	print("------------------------")
	return out
}

func InitAndApplyE(options *Options) (string, error) {
	if _, err := InitE(options); err != nil {
		return "", err
	}

	if _, err := GetE(options); err != nil {
		return "", err
	}

	return ApplyE(options)
}

func Apply(t *testing.T, options *Options) string {
	out, err := ApplyE(options)
	require.NoError(t, err)
	return out
}

func TgApplyAll(t *testing.T, options *Options) string {
	out, err := TgApplyAllE(t, options)
	require.NoError(t, err)
	return out
}

func ApplyE(options *Options) (string, error) {
	return RunTerraformCommandE(options, FormatArgs(options, "apply", "-input=false", "-lock=false", "-auto-approve")...)
}

func TgApplyAllE(t *testing.T, options *Options) (string, error) {
	if options.TerraformBinary != "terragrunt" {
		return "", TgInvalidBinary(options.TerraformBinary)
	}

	return RunTerraformCommandE(options, FormatArgs(options, "apply-all", "-input=false", "-lock=false", "-auto-approve")...)
}
