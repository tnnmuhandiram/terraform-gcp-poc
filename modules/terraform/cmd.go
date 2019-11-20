package terraform

import (
	"fmt"

	"github.com/tnnmuhandiram/terraform-gcp-poc/modules/collections"
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

func RunTerraformCommandAndGetStdoutE(additionalOptions *Options, additionalArgs ...string) (string, error) {
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
