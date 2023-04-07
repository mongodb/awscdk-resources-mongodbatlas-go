// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type AuthConfig struct {
	// The type of authentication event that the trigger listens for.
	OperationType AuthConfigOperationType `field:"required" json:"operationType" yaml:"operationType"`
	// The type(s) of authentication provider that the trigger listens to.
	Providers *[]AuthConfigProviders `field:"required" json:"providers" yaml:"providers"`
}

