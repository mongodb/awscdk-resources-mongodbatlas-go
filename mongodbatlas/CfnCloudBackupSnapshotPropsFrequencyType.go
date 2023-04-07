// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Human-readable label that identifies how often this snapshot triggers.
type CfnCloudBackupSnapshotPropsFrequencyType string

const (
	// hourly.
	CfnCloudBackupSnapshotPropsFrequencyType_HOURLY CfnCloudBackupSnapshotPropsFrequencyType = "HOURLY"
	// daily.
	CfnCloudBackupSnapshotPropsFrequencyType_DAILY CfnCloudBackupSnapshotPropsFrequencyType = "DAILY"
	// weekly.
	CfnCloudBackupSnapshotPropsFrequencyType_WEEKLY CfnCloudBackupSnapshotPropsFrequencyType = "WEEKLY"
	// monthly.
	CfnCloudBackupSnapshotPropsFrequencyType_MONTHLY CfnCloudBackupSnapshotPropsFrequencyType = "MONTHLY"
)

