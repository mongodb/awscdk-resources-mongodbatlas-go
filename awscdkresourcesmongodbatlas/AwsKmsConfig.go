package awscdkresourcesmongodbatlas


// Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
type AwsKmsConfig struct {
	// Unique alphanumeric string that identifies the Amazon Web Services (AWS) Customer Master Key (CMK) you used to encrypt and decrypt the MongoDB master keys.
	CustomerMasterKeyId *string `field:"optional" json:"customerMasterKeyId" yaml:"customerMasterKeyId"`
	// Specifies whether Encryption at Rest is enabled for an Atlas project.
	//
	// To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Physical location where MongoDB Atlas deploys your AWS-hosted MongoDB cluster nodes.
	//
	// The region you choose can affect network latency for clients accessing your databases.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Enable connection to your Amazon Web Services (AWS) Key Management Service (KMS) over private networking.
	RequirePrivateNetworking *bool `field:"optional" json:"requirePrivateNetworking" yaml:"requirePrivateNetworking"`
	// Unique 24-hexadecimal digit string that identifies an Amazon Web Services (AWS) Identity and Access Management (IAM) role.
	//
	// This IAM role has the permissions required to manage your AWS customer master key.
	RoleId *string `field:"optional" json:"roleId" yaml:"roleId"`
	// Flag that indicates whether the Amazon Web Services (AWS) Key Management Service (KMS) encryption key can encrypt and decrypt data.
	Valid *bool `field:"optional" json:"valid" yaml:"valid"`
}

