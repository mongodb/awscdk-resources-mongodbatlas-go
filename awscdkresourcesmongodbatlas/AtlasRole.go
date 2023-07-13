package awscdkresourcesmongodbatlas


type AtlasRole struct {
	// Unique 24-hexadecimal digit string that identifies the organization to which this role belongs.
	//
	// You can set a value for this parameter or **groupId** but not both in the same request.
	OrgId *string `field:"optional" json:"orgId" yaml:"orgId"`
	// Unique 24-hexadecimal digit string that identifies the project to which this role belongs.
	//
	// You can set a value for this parameter or **orgId** but not both in the same request.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Human-readable label that identifies the collection of privileges that MongoDB Cloud grants a specific API key, MongoDB Cloud user, or MongoDB Cloud team.
	//
	// These roles include organization- and project-level roles.
	//
	// Organization Roles
	//
	// * ORG_OWNER
	// * ORG_MEMBER
	// * ORG_GROUP_CREATOR
	// * ORG_BILLING_ADMIN
	// * ORG_READ_ONLY
	//
	// Project Roles
	//
	// * GROUP_CLUSTER_MANAGER
	// * GROUP_DATA_ACCESS_ADMIN
	// * GROUP_DATA_ACCESS_READ_ONLY
	// * GROUP_DATA_ACCESS_READ_WRITE
	// * GROUP_OWNER
	// * GROUP_READ_ONLY.
	RoleName AtlasRoleRoleName `field:"optional" json:"roleName" yaml:"roleName"`
}

