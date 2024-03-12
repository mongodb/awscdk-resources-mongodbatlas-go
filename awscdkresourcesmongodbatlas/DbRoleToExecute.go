package awscdkresourcesmongodbatlas


// The name of a Built in or Custom DB Role to connect to an Atlas Cluster.
type DbRoleToExecute struct {
	// The name of the role to use.
	//
	// Can be a built in role or a custom role.
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Type of the DB role.
	//
	// Can be either BuiltIn or Custom.
	Type DbRoleToExecuteType `field:"optional" json:"type" yaml:"type"`
}

