package awscdkresourcesmongodbatlas


// Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
type AwsKmsConfig struct {
	// The AWS customer master key used to encrypt and decrypt the MongoDB master keys.
	CustomerMasterKeyId *string `field:"optional" json:"customerMasterKeyId" yaml:"customerMasterKeyId"`
	// Specifies whether Encryption at Rest is enabled for an Atlas project.
	//
	// To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The AWS region in which the AWS customer master key exists.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// ID of an AWS IAM role authorized to manage an AWS customer master key.
	RoleId *string `field:"optional" json:"roleId" yaml:"roleId"`
}

