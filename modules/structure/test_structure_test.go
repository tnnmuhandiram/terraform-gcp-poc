package structure

import "testing"

func TestCopyToTempFolder(t *testing.T) {
	tempFolder := CopyTerraformFolderToTemp(t, "../../", "scripts")
	t.Log(tempFolder)
}

func TestCopySubtestToTempFolder(t *testing.T) {
	t.Run("Subtest", func(t *testing.T) {
		tempFolder := CopyTerraformFolderToTemp(t, "../../", "scripts")
		t.Log(tempFolder)
	})
}
