package terraform

import (
	"fmt"
)

func Init(options *Options) string {
	out, err := InitE(options)
	if err != nil {
		// t.Fatal(err)
	}
	return out
}

func InitE(options *Options) (string, error) {
	args := []string{"init", fmt.Sprintf("-upgrade=%t", options.Upgrade)}
	args = append(args, FormatTerraformBackendConfigAsArgs(options.BackendConfig)...)
	return RunTerraformCommandE(options, args...)
}
