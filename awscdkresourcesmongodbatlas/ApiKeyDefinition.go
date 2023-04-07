// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type ApiKeyDefinition struct {
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
}

