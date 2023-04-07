// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type AtlasEncryptionAtRestProps struct {
	// The AWS customer master key used to encrypt and decrypt the MongoDB master keys.
	CustomerMasterKeyId *string `field:"required" json:"customerMasterKeyId" yaml:"customerMasterKeyId"`
	// Unique identifier of the Atlas project to which the user belongs.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// ID of an AWS IAM role authorized to manage an AWS customer master key.
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
	// Specifies whether Encryption at Rest is enabled for an Atlas project.
	//
	// To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.
	// Default Value: true.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// The AWS region in which the AWS customer master key exists.
	//
	// Default Value: US_EAST_1.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

