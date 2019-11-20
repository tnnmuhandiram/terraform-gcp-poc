package terraform

func Destroy(options *Options) string {
	out, err := DestroyE(options)
	print(err)
	return out
}

func DestroyE(options *Options) (string, error) {
	return RunTerraformCommandE(options, FormatArgs(options, "destroy", "-auto-approve", "-input=false", "-lock=false")...)
}
