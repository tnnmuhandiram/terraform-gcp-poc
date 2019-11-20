package terraform

import (
	"fmt"
	"regexp"
	"strings"
)

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
