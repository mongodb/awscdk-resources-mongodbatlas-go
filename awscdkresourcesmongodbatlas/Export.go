package awscdkresourcesmongodbatlas


type Export struct {
	// Unique identifier of the AWS bucket to export the cloud backup snapshot to.
	ExportBucketId *string `field:"optional" json:"exportBucketId" yaml:"exportBucketId"`
	// Frequency associated with the export policy.
	//
	// Value can be daily, weekly, monthly or yearly.
	FrequencyType *string `field:"optional" json:"frequencyType" yaml:"frequencyType"`
}

