package terraform

func PlanE(options *Options) (string, error) {
	return RunTerraformCommandE(options, FormatArgs(options, "plan", "-input=false", "-lock=false")...)
}
