// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type ScheduleViewType string

const (
	// DAILY.
	ScheduleViewType_DAILY ScheduleViewType = "DAILY"
	// MONTHLY.
	ScheduleViewType_MONTHLY ScheduleViewType = "MONTHLY"
	// DEFAULT.
	ScheduleViewType_DEFAULT ScheduleViewType = "DEFAULT"
	// WEEKLY.
	ScheduleViewType_WEEKLY ScheduleViewType = "WEEKLY"
)

