package terraform

type ResourceCount struct {
	Add     int
	Change  int
	Destroy int
}

const (
	applyRegexp             = `Apply complete! Resources: (\d+) added, (\d+) changed, (\d+) destroyed\.`
	destroyRegexp           = `Destroy complete! Resources: (\d+) destroyed\.`
	planWithChangesRegexp   = `(\033\[1m)?Plan:(\033\[0m)? (\d+) to add, (\d+) to change, (\d+) to destroy\.`
	planWithNoChangesRegexp = `No changes\. Infrastructure is up-to-date\.`
)

const getResourceCountErrMessage = "Can't parse Terraform output"
