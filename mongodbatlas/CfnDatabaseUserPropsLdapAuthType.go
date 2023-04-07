// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Method by which the provided username is authenticated.
//
// If no value is given, Atlas uses the default value of NONE.
type CfnDatabaseUserPropsLdapAuthType string

const (
	// NONE.
	CfnDatabaseUserPropsLdapAuthType_NONE CfnDatabaseUserPropsLdapAuthType = "NONE"
	// USER.
	CfnDatabaseUserPropsLdapAuthType_USER CfnDatabaseUserPropsLdapAuthType = "USER"
	// GROUP.
	CfnDatabaseUserPropsLdapAuthType_GROUP CfnDatabaseUserPropsLdapAuthType = "GROUP"
)

