// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type EventAwseventbridgeAwsConfig struct {
	// An AWS Account ID.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// If `true`, event objects are serialized using EJSON.
	ExtendedJsonEnabled *bool `field:"optional" json:"extendedJsonEnabled" yaml:"extendedJsonEnabled"`
	// An AWS region.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

