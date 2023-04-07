// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type ScopeDefinitionType string

const (
	// CLUSTER.
	ScopeDefinitionType_CLUSTER ScopeDefinitionType = "CLUSTER"
	// DATA_LAKE.
	ScopeDefinitionType_DATA_LAKE ScopeDefinitionType = "DATA_LAKE"
)

