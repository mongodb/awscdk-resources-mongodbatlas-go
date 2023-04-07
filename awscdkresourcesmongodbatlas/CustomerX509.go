// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type CustomerX509 struct {
	// PEM string containing one or more customer CAs for database user authentication.
	Cas *string `field:"optional" json:"cas" yaml:"cas"`
}

