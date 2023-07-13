package awscdkresourcesmongodbatlas


// Human-readable label that indicates whether the new database user authenticates with the Amazon Web Services (AWS) Identity and Access Management (IAM) credentials associated with the user or the user's role.
type CfnDatabaseUserPropsAwsiamType string

const (
	// NONE.
	CfnDatabaseUserPropsAwsiamType_NONE CfnDatabaseUserPropsAwsiamType = "NONE"
	// USER.
	CfnDatabaseUserPropsAwsiamType_USER CfnDatabaseUserPropsAwsiamType = "USER"
	// ROLE.
	CfnDatabaseUserPropsAwsiamType_ROLE CfnDatabaseUserPropsAwsiamType = "ROLE"
)

