package awscdkresourcesmongodbatlas


// Returns, adds, edits, and removes Federated Database Instances.
type CfnFederatedDatabaseInstanceProps struct {
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Human-readable label that identifies the data federation.
	TenantName *string `field:"required" json:"tenantName" yaml:"tenantName"`
	// Cloud provider linked to this Atlas Data Federation.
	CloudProviderConfig *CloudProviderConfig `field:"optional" json:"cloudProviderConfig" yaml:"cloudProviderConfig"`
	// Information about the cloud provider region to which the Atlas Data Federation routes client connections.
	//
	// MongoDB Cloud supports AWS only.
	DataProcessRegion *DataProcessRegion `field:"optional" json:"dataProcessRegion" yaml:"dataProcessRegion"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Flag that indicates whether this request should check if the requesting IAM role can read from the S3 bucket.
	//
	// AWS checks if the role can list the objects in the bucket before writing to it. Some IAM roles only need write permissions. This flag allows you to skip that check.
	SkipRoleValidation *bool `field:"optional" json:"skipRoleValidation" yaml:"skipRoleValidation"`
	// Configuration information for each data store and its mapping to MongoDB Cloud databases.
	Storage *Storage `field:"optional" json:"storage" yaml:"storage"`
}

