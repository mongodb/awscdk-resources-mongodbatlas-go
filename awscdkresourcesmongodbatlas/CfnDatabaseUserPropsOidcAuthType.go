package awscdkresourcesmongodbatlas


// Human-readable label that indicates whether the new database user or group authenticates with OIDC federated authentication.
//
// To create a federated authentication user, specify the value of USER in this field. To create a federated authentication group, specify the value of IDP_GROUP in this field. Default value is `NONE`.
type CfnDatabaseUserPropsOidcAuthType string

const (
	// NONE.
	CfnDatabaseUserPropsOidcAuthType_NONE CfnDatabaseUserPropsOidcAuthType = "NONE"
	// USER.
	CfnDatabaseUserPropsOidcAuthType_USER CfnDatabaseUserPropsOidcAuthType = "USER"
	// IDP_GROUP.
	CfnDatabaseUserPropsOidcAuthType_IDP_GROUP CfnDatabaseUserPropsOidcAuthType = "IDP_GROUP"
)

