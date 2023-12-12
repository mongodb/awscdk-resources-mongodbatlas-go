package awscdkresourcesmongodbatlas


type ApiKey struct {
	// Purpose or explanation provided when someone created this organization API key.
	//
	// 1 to 250 characters.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of roles to grant this API key.
	//
	// If you provide this list, provide a minimum of one role and ensure each role applies to this organization.
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
}

