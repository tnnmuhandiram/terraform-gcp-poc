package structure

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tnnmuhandiram/terraform-gcp-poc/modules/files"
	"github.com/tnnmuhandiram/terraform-gcp-poc/modules/logger"
)

const SKIP_STAGE_ENV_VAR_PREFIX = "SKIP_"

func RunTestStage(t *testing.T, stageName string, stage func()) {
	envVarName := fmt.Sprintf("%s%s", SKIP_STAGE_ENV_VAR_PREFIX, stageName)
	if os.Getenv(envVarName) == "" {
		logger.Logf(t, "The '%s' environment variable is not set, so executing stage '%s'.", envVarName, stageName)
		stage()
	} else {
		logger.Logf(t, "The '%s' environment variable is set, so skipping stage '%s'.", envVarName, stageName)
	}
}

func SkipStageEnvVarSet() bool {
	for _, environmentVariable := range os.Environ() {
		if strings.HasPrefix(environmentVariable, SKIP_STAGE_ENV_VAR_PREFIX) {
			return true
		}
	}

	return false
}

func CopyTerraformFolderToTemp(rootFolder string, terraformModuleFolder string) string {
	if SkipStageEnvVarSet() {
		return filepath.Join(rootFolder, terraformModuleFolder)
	}

	tmpRootFolder, err := files.CopyTerraformFolderToTemp(rootFolder, cleanName("terraform-template"))
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
