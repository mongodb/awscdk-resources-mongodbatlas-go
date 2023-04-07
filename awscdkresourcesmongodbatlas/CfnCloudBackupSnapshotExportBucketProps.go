// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


// The exportBuckets resource allows you to grant Atlas access to the specified bucket for exporting backup snapshots.
type CfnCloudBackupSnapshotExportBucketProps struct {
	// Human-readable label that identifies the AWS bucket that the role is authorized to access.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Unique 24-hexadecimal character string that identifies the AWS IAM role that MongoDB Cloud uses to access the AWS S3 bucket.
	IamRoleId *string `field:"required" json:"iamRoleId" yaml:"iamRoleId"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

