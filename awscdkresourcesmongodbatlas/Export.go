// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type Export struct {
	// Unique identifier of the AWS bucket to export the cloud backup snapshot to.
	ExportBucketId *string `field:"optional" json:"exportBucketId" yaml:"exportBucketId"`
	// Frequency associated with the export policy.
	//
	// Value can be daily, weekly, or monthly.
	FrequencyType *string `field:"optional" json:"frequencyType" yaml:"frequencyType"`
}

