package awscdkresourcesmongodbatlas


// Authentication configuration for Schema Registry.
type SchemaRegistryAuthentication struct {
	// Password or Private Key for authentication.
	//
	// Review [AWS security best practices for CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/security-best-practices.html#creds) to manage credentials.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Authentication type discriminator.
	//
	// Specifies the authentication mechanism for Schema Registry.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Username or Public Key for authentication.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

