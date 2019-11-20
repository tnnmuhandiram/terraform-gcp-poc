package terraform

func InitAndApply(options *Options) string {
	out, err := InitAndApplyE(options)

	print(err)
	print(out)
	return out
}

func InitAndApplyE(options *Options) (string, error) {
	if _, err := InitE(options); err != nil {
		return "", err
	}

	if _, err := GetE(options); err != nil {
		return "", err
	}

	return ApplyE(options)
}

func ApplyE(options *Options) (string, error) {
	return RunTerraformCommandE(options, FormatArgs(options, "apply", "-input=false", "-lock=false", "-auto-approve")...)
}
