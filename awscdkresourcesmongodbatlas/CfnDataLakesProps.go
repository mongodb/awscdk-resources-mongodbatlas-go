package awscdkresourcesmongodbatlas


// Returns, adds, edits, and removes Federated Database Instances.
type CfnDataLakesProps struct {
	// Cloud provider linked to this data lake.
	CloudProviderConfig *DataLakeCloudProviderConfigView `field:"optional" json:"cloudProviderConfig" yaml:"cloudProviderConfig"`
	// Information about the cloud provider region to which the data lake routes client connections.
	//
	// MongoDB Cloud supports AWS only.
	DataProcessRegion *DataLakeDataProcessRegionView `field:"optional" json:"dataProcessRegion" yaml:"dataProcessRegion"`
	// Timestamp that specifies the end point for the range of log messages to download.
	//
	// MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch.
	EndDate *float64 `field:"optional" json:"endDate" yaml:"endDate"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Flag that indicates whether this request should check if the requesting IAM role can read from the S3 bucket.
	//
	// AWS checks if the role can list the objects in the bucket before writing to it. Some IAM roles only need write permissions. This flag allows you to skip that check.
	SkipRoleValidation *bool `field:"optional" json:"skipRoleValidation" yaml:"skipRoleValidation"`
	// Configuration information for each data store and its mapping to MongoDB Cloud databases.
	Storage *DataLakeStorageView `field:"optional" json:"storage" yaml:"storage"`
	// Human-readable label that identifies the Federated Database to remove.
	TenantName *string `field:"optional" json:"tenantName" yaml:"tenantName"`
}

