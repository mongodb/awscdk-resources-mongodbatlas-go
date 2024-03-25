package awscdkresourcesmongodbatlas


// Type of the DB role.
//
// Can be either BuiltIn or Custom.
type DbRoleToExecuteType string

const (
	// BUILT_IN.
	DbRoleToExecuteType_BUILT_IN DbRoleToExecuteType = "BUILT_IN"
	// CUSTOM.
	DbRoleToExecuteType_CUSTOM DbRoleToExecuteType = "CUSTOM"
)

