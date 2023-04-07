// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Method that briefs who owns the certificate provided.
//
// If no value is given while using X509Type, Atlas uses the default value of MANAGED.
type CfnDatabaseUserPropsX509Type string

const (
	// NONE.
	CfnDatabaseUserPropsX509Type_NONE CfnDatabaseUserPropsX509Type = "NONE"
	// MANAGED.
	CfnDatabaseUserPropsX509Type_MANAGED CfnDatabaseUserPropsX509Type = "MANAGED"
	// CUSTOMER.
	CfnDatabaseUserPropsX509Type_CUSTOMER CfnDatabaseUserPropsX509Type = "CUSTOMER"
)

