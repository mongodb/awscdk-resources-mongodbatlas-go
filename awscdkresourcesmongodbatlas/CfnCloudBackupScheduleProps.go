// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


// An example resource schema demonstrating some basic constructs and validation rules.
type CfnCloudBackupScheduleProps struct {
	// The name of the Atlas cluster that contains the snapshots you want to retrieve.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"required" json:"profile" yaml:"profile"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Flag that indicates whether automatic export of cloud backup snapshots to the AWS bucket is enabled.
	AutoExportEnabled *bool `field:"optional" json:"autoExportEnabled" yaml:"autoExportEnabled"`
	// List that contains a document for each copy setting item in the desired backup policy.
	CopySettings *[]*ApiAtlasDiskBackupCopySettingView `field:"optional" json:"copySettings" yaml:"copySettings"`
	// List that contains a document for each deleted copy setting whose backup copies you want to delete.
	DeleteCopiedBackups *[]*ApiDeleteCopiedBackupsView `field:"optional" json:"deleteCopiedBackups" yaml:"deleteCopiedBackups"`
	// Policy for automatically exporting cloud backup snapshots.
	Export *Export `field:"optional" json:"export" yaml:"export"`
	// Unique identifier of the snapshot.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both.
	//
	// RFC 5988 outlines these relationships.
	Links *[]*Link `field:"optional" json:"links" yaml:"links"`
	// Rules set for this backup schedule.
	Policies *[]*ApiPolicyView `field:"optional" json:"policies" yaml:"policies"`
	// UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.
	ReferenceHourOfDay *float64 `field:"optional" json:"referenceHourOfDay" yaml:"referenceHourOfDay"`
	// UTC Minute of day between 0 and 59 representing which minute of the referenceHourOfDay that Atlas takes the snapshot.
	ReferenceMinuteOfHour *float64 `field:"optional" json:"referenceMinuteOfHour" yaml:"referenceMinuteOfHour"`
	// Number of days back in time you can restore to with Continuous Cloud Backup accuracy.
	//
	// Must be a positive, non-zero integer.
	RestoreWindowDays *float64 `field:"optional" json:"restoreWindowDays" yaml:"restoreWindowDays"`
	// Flag indicating if updates to retention in the backup policy were applied to snapshots that Atlas took earlier.
	UpdateSnapshots *bool `field:"optional" json:"updateSnapshots" yaml:"updateSnapshots"`
	// Specify true to use organization and project names instead of organization and project UUIDs in the path for the metadata files that Atlas uploads to your S3 bucket after it finishes exporting the snapshots.
	UseOrgAndGroupNamesInExportPrefix *bool `field:"optional" json:"useOrgAndGroupNamesInExportPrefix" yaml:"useOrgAndGroupNamesInExportPrefix"`
}

