// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// The type of authentication event that the trigger listens for.
type AuthConfigOperationType string

const (
	// LOGIN.
	AuthConfigOperationType_LOGIN AuthConfigOperationType = "LOGIN"
	// CREATE.
	AuthConfigOperationType_CREATE AuthConfigOperationType = "CREATE"
	// DELETE.
	AuthConfigOperationType_DELETE AuthConfigOperationType = "DELETE"
)

