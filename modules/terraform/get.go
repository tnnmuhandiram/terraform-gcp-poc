package terraform

func GetE(options *Options) (string, error) {
	return RunTerraformCommandE(options, "get", "-update")
}
