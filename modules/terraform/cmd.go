package terraform

import (
	"fmt"
	"testing"

	"github.com/tnnmuhandiram/terraform-gcp-poc/modules/collections"
	"github.com/tnnmuhandiram/terraform-gcp-poc/modules/logger"
	"github.com/tnnmuhandiram/terraform-gcp-poc/modules/retry"
	"github.com/tnnmuhandiram/terraform-gcp-poc/modules/shell"
)

func GetCommonOptions(options *Options, args ...string) (*Options, []string) {
	if options.NoColor && !collections.ListContains(args, "-no-color") {
		args = append(args, "-no-color")
	}

	if options.TerraformBinary == "" {
		options.TerraformBinary = "terraform"
	}

	if options.TerraformBinary == "terragrunt" {
		args = append(args, "--terragrunt-non-interactive")
	}

	if options.SshAgent != nil {
		if options.EnvVars == nil {
			options.EnvVars = map[string]string{}
		}
		options.EnvVars["SSH_AUTH_SOCK"] = options.SshAgent.SocketFile()
	}
	return options, args
}

func RunTerraformCommand(t *testing.T, additionalOptions *Options, args ...string) string {
	out, err := RunTerraformCommandE(additionalOptions, args...)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

func RunTerraformCommandE(additionalOptions *Options, additionalArgs ...string) (string, error) {
	options, args := GetCommonOptions(additionalOptions, additionalArgs...)

	cmd := shell.Command{
		Command:           options.TerraformBinary,
		Args:              args,
		WorkingDir:        options.TerraformDir,
		Env:               options.EnvVars,
		OutputMaxLineSize: options.OutputMaxLineSize,
	}

	description := fmt.Sprintf("%s %v", options.TerraformBinary, args)
	return retry.DoWithRetryableErrorsE(description, options.RetryableTerraformErrors, options.MaxRetries, options.TimeBetweenRetries, func() (string, error) {
		return shell.RunCommandAndGetOutputE(cmd)
	})
}

func RunTerraformCommandAndGetStdoutE(t *testing.T, additionalOptions *Options, additionalArgs ...string) (string, error) {
	options, args := GetCommonOptions(additionalOptions, additionalArgs...)

	cmd := shell.Command{
		Command:    options.TerraformBinary,
		Args:       args,
		WorkingDir: options.TerraformDir,
		Env:        options.EnvVars,
	}

	description := fmt.Sprintf("%s %v", options.TerraformBinary, args)
	return retry.DoWithRetryableErrorsE(description, options.RetryableTerraformErrors, options.MaxRetries, options.TimeBetweenRetries, func() (string, error) {
		return shell.RunCommandAndGetStdOutE(cmd)
	})
}

func GetExitCodeForTerraformCommand(t *testing.T, additionalOptions *Options, args ...string) int {
	exitCode, err := GetExitCodeForTerraformCommandE(t, additionalOptions, args...)
	if err != nil {
		t.Fatal(err)
	}
	return exitCode
}

func GetExitCodeForTerraformCommandE(t *testing.T, additionalOptions *Options, additionalArgs ...string) (int, error) {
	options, args := GetCommonOptions(additionalOptions, additionalArgs...)

	logger.Logf(t, "Running %s with args %v", options.TerraformBinary, args)
	cmd := shell.Command{
		Command:           options.TerraformBinary,
		Args:              args,
		WorkingDir:        options.TerraformDir,
		Env:               options.EnvVars,
		OutputMaxLineSize: options.OutputMaxLineSize,
	}

	_, err := shell.RunCommandAndGetOutputE(cmd)
	if err == nil {
		return DefaultSuccessExitCode, nil
	}
	exitCode, getExitCodeErr := shell.GetExitCodeForRunCommandError(err)
	if getExitCodeErr == nil {
		return exitCode, nil
	}
	return DefaultErrorExitCode, getExitCodeErr
}
