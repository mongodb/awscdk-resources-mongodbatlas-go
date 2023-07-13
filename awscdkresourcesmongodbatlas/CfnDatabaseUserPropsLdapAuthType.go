package awscdkresourcesmongodbatlas


// Method by which the provided username is authenticated.
//
// Default value is `NONE`.
type CfnDatabaseUserPropsLdapAuthType string

const (
	// NONE.
	CfnDatabaseUserPropsLdapAuthType_NONE CfnDatabaseUserPropsLdapAuthType = "NONE"
	// USER.
	CfnDatabaseUserPropsLdapAuthType_USER CfnDatabaseUserPropsLdapAuthType = "USER"
	// GROUP.
	CfnDatabaseUserPropsLdapAuthType_GROUP CfnDatabaseUserPropsLdapAuthType = "GROUP"
)

