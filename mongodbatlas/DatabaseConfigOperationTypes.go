// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type DatabaseConfigOperationTypes string

const (
	// INSERT.
	DatabaseConfigOperationTypes_INSERT DatabaseConfigOperationTypes = "INSERT"
	// UPDATE.
	DatabaseConfigOperationTypes_UPDATE DatabaseConfigOperationTypes = "UPDATE"
	// REPLACE.
	DatabaseConfigOperationTypes_REPLACE DatabaseConfigOperationTypes = "REPLACE"
	// DELETE.
	DatabaseConfigOperationTypes_DELETE DatabaseConfigOperationTypes = "DELETE"
)

