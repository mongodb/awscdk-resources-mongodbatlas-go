// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type DataLakeAwsCloudProviderConfigView struct {
	// Unique identifier associated with the Identity and Access Management (IAM) role that the data lake assumes when accessing the data stores.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that the data lake assumes when accessing data stores.
	IamAssumedRoleArn *string `field:"optional" json:"iamAssumedRoleArn" yaml:"iamAssumedRoleArn"`
	// Amazon Resource Name (ARN) of the user that the data lake assumes when accessing data stores.
	IamUserArn *string `field:"optional" json:"iamUserArn" yaml:"iamUserArn"`
	// Unique identifier of the role that the data lake can use to access the data stores.Required if specifying cloudProviderConfig.
	RoleId *string `field:"optional" json:"roleId" yaml:"roleId"`
	// Name of the S3 data bucket that the provided role ID is authorized to access.Required if specifying cloudProviderConfig.
	TestS3Bucket *string `field:"optional" json:"testS3Bucket" yaml:"testS3Bucket"`
}

