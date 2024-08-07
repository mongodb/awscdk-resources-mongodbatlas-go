package awscdkresourcesmongodbatlas


// Returns and edits the Encryption at Rest using Customer Key Management configuration.
type CfnEncryptionAtRestProps struct {
	AwsKmsConfig *AwsKmsConfig `field:"required" json:"awsKmsConfig" yaml:"awsKmsConfig"`
	// Unique identifier of the Atlas project to which the user belongs.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

