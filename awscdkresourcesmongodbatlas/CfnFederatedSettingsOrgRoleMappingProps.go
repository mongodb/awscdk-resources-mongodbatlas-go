package awscdkresourcesmongodbatlas


// Returns, adds, edits, and removes federation-related features such as role mappings and connected organization configurations.
type CfnFederatedSettingsOrgRoleMappingProps struct {
	// Unique human-readable label that identifies the identity provider group to whichthis role mapping applies.
	ExternalGroupName *string `field:"required" json:"externalGroupName" yaml:"externalGroupName"`
	// Unique 24-hexadecimal digit string that identifies your federation.
	FederationSettingsId *string `field:"required" json:"federationSettingsId" yaml:"federationSettingsId"`
	// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Atlas roles and the unique identifiers of the groups and organizations associated with each role.
	RoleAssignments *[]*RoleAssignment `field:"optional" json:"roleAssignments" yaml:"roleAssignments"`
}

