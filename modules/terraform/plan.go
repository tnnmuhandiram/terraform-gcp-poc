package terraform

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func InitAndPlan(t *testing.T, options *Options) string {
	out, err := InitAndPlanE(t, options)
	require.NoError(t, err)
	return out
}

func InitAndPlanE(t *testing.T, options *Options) (string, error) {
	if _, err := InitE(options); err != nil {
		return "", err
	}

	if _, err := GetE(options); err != nil {
		return "", err
	}

	return PlanE(options)
}

func Plan(t *testing.T, options *Options) string {
	out, err := PlanE(options)
	require.NoError(t, err)
	return out
}

func PlanE(options *Options) (string, error) {
	return RunTerraformCommandE(options, FormatArgs(options, "plan", "-input=false", "-lock=false")...)
}

func InitAndPlanWithExitCode(t *testing.T, options *Options) int {
	exitCode, err := InitAndPlanWithExitCodeE(t, options)
	require.NoError(t, err)
	return exitCode
}

func InitAndPlanWithExitCodeE(t *testing.T, options *Options) (int, error) {
	if _, err := InitE(options); err != nil {
		return DefaultErrorExitCode, err
	}

	return PlanExitCodeE(t, options)
}

func PlanExitCode(t *testing.T, options *Options) int {
	exitCode, err := PlanExitCodeE(t, options)
	require.NoError(t, err)
	return exitCode
}

func PlanExitCodeE(t *testing.T, options *Options) (int, error) {
	return GetExitCodeForTerraformCommandE(t, options, FormatArgs(options, "plan", "-input=false", "-lock=true", "-detailed-exitcode")...)
}

func TgPlanAllExitCode(t *testing.T, options *Options) int {
	exitCode, err := TgPlanAllExitCodeE(t, options)
	require.NoError(t, err)
	return exitCode
}

func TgPlanAllExitCodeE(t *testing.T, options *Options) (int, error) {
	if options.TerraformBinary != "terragrunt" {
		return 1, fmt.Errorf("terragrunt must be set as TerraformBinary to use this method")
	}

	return GetExitCodeForTerraformCommandE(t, options, FormatArgs(options, "plan-all", "--input=false", "--lock=true", "--detailed-exitcode")...)
}
