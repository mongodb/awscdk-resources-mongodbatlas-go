// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type RoleAssignment struct {
	// List that contains comma-separated key value pairs to map zones to geographic regions.
	//
	// These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to a unique 24-hexadecimal string that identifies the custom zone.
	//
	// This parameter returns an empty object if no custom zones exist.
	OrgId *string `field:"optional" json:"orgId" yaml:"orgId"`
	// List that contains comma-separated key value pairs to map zones to geographic regions.
	//
	// These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to a unique 24-hexadecimal string that identifies the custom zone.
	//
	// This parameter returns an empty object if no custom zones exist.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	Role *string `field:"optional" json:"role" yaml:"role"`
}

