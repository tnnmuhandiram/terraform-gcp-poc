package terraform

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func WorkspaceSelectOrNew(t *testing.T, options *Options, name string) string {
	out, err := WorkspaceSelectOrNewE(t, options, name)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

func WorkspaceSelectOrNewE(t *testing.T, options *Options, name string) (string, error) {
	out, err := RunTerraformCommandE(options, "workspace", "list")
	if err != nil {
		return "", err
	}

	if isExistingWorkspace(out, name) {
		_, err = RunTerraformCommandE(options, "workspace", "select", name)
	} else {
		_, err = RunTerraformCommandE(options, "workspace", "new", name)
	}
	if err != nil {
		return "", err
	}

	return RunTerraformCommandE(options, "workspace", "show")
}

func isExistingWorkspace(out string, name string) bool {
	workspaces := strings.Split(out, "\n")
	for _, ws := range workspaces {
		if nameMatchesWorkspace(name, ws) {
			return true
		}
	}
	return false
}

func nameMatchesWorkspace(name string, workspace string) bool {
	match, _ := regexp.MatchString(fmt.Sprintf("^\\*?\\s*%s$", name), workspace)
	return match
}
