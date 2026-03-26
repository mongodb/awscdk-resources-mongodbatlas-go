package awscdkresourcesmongodbatlas


// The log integration resource provides access to push-based log export configurations for MongoDB Atlas.
//
// The resource allows you to create, edit and delete log export integrations (e.g. to cloud storage or log services). This release supports S3 integration; additional integration types will be added as the API supports them. The resource requires your Project ID.
type CfnLogIntegrationProps struct {
	// Human-readable label that identifies the bucket (or storage container) name for storing log files.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Unique 24-hexadecimal digit string that identifies the AWS IAM role that MongoDB Cloud uses to access your S3 bucket.
	IamRoleId *string `field:"required" json:"iamRoleId" yaml:"iamRoleId"`
	// Array of log types to export.
	//
	// Valid values depend on the integration type (e.g. MONGOD, MONGOS, MONGOD_AUDIT, MONGOS_AUDIT for S3).
	LogTypes *[]*string `field:"required" json:"logTypes" yaml:"logTypes"`
	// Path prefix where log files will be stored.
	//
	// MongoDB Cloud adds further sub-directories based on the log type.
	PrefixPath *string `field:"required" json:"prefixPath" yaml:"prefixPath"`
	// Unique 24-hexadecimal digit string that identifies your project.
	//
	// Use the /groups endpoint to retrieve all projects to which the authenticated user has access. Groups and projects are synonymous terms. Your group id is the same as your project id.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud.
	//
	// The value must match the log integration type.
	Type *string `field:"required" json:"type" yaml:"type"`
	// AWS KMS key ID or ARN for server-side encryption (optional).
	//
	// If not provided, uses bucket default encryption settings.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

