package awscdkresourcesmongodbatlas


type ProjectAssignment struct {
	// Unique 24-hexadecimal digit string that identifies the project in an organization.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// List of roles to grant this API key.
	//
	// If you provide this list, provide a minimum of one role and ensure each role applies to this organization.
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
}

