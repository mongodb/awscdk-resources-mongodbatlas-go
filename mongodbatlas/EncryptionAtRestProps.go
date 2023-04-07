// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type EncryptionAtRestProps struct {
	// The AWS customer master key used to encrypt and decrypt the MongoDB master keys.
	CustomerMasterKeyId *string `field:"required" json:"customerMasterKeyId" yaml:"customerMasterKeyId"`
	// ID of an AWS IAM role authorized to manage an AWS customer master key.
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
	// Specifies whether Encryption at Rest is enabled for an Atlas project.
	//
	// To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.
	// Default Value: true.
	EnabledEncryptionAtRest *bool `field:"optional" json:"enabledEncryptionAtRest" yaml:"enabledEncryptionAtRest"`
	// The AWS region in which the AWS customer master key exists.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

