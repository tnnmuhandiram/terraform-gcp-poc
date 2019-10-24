package terraform

import (
	"testing"
)

func Get(t *testing.T, options *Options) string {
	out, err := GetE(options)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

func GetE(options *Options) (string, error) {
	return RunTerraformCommandE(options, "get", "-update")
}
