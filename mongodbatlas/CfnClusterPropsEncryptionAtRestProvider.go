// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Cloud service provider that manages your customer keys to provide an additional layer of encryption at rest for the cluster.
//
// To enable customer key management for encryption at rest, the cluster replicationSpecs[n].regionConfigs[m].{type}Specs.instanceSize setting must be M10 or higher and "backupEnabled" : false or omitted entirely.
type CfnClusterPropsEncryptionAtRestProvider string

const (
	// AWS.
	CfnClusterPropsEncryptionAtRestProvider_AWS CfnClusterPropsEncryptionAtRestProvider = "AWS"
	// GCP.
	CfnClusterPropsEncryptionAtRestProvider_GCP CfnClusterPropsEncryptionAtRestProvider = "GCP"
	// AZURE.
	CfnClusterPropsEncryptionAtRestProvider_AZURE CfnClusterPropsEncryptionAtRestProvider = "AZURE"
	// NONE.
	CfnClusterPropsEncryptionAtRestProvider_NONE CfnClusterPropsEncryptionAtRestProvider = "NONE"
)

