package structure

import (
	"path/filepath"
	"strings"

	"github.com/tnnmuhandiram/terraform/modules/files"
)

func CopyTerraformFolderToTemp(rootFolder string, terraformModuleFolder string) string {

	tmpRootFolder, err := files.CopyTerraformFolderToTemp(rootFolder, cleanName("gcp-instance-testing"))
	if err != nil {
		// t.Fatal(err)
	}

	tmpTestFolder := filepath.Join(tmpRootFolder, terraformModuleFolder)
	return tmpTestFolder
}

func cleanName(originalName string) string {
	parts := strings.Split(originalName, "/")
	return parts[len(parts)-1]
}
