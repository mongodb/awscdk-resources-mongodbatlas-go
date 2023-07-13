package awscdkresourcesmongodbatlas


// Returns, adds, edits, and removes custom database user privilege roles.
type CfnCustomDbRoleProps struct {
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Human-readable label that identifies the role for the request.
	//
	// This name must be unique for this custom role in this project.
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// List of the individual privilege actions that the role grants.
	Actions *[]*Action `field:"optional" json:"actions" yaml:"actions"`
	// List of the built-in roles that this custom role inherits.
	InheritedRoles *[]*InheritedRole `field:"optional" json:"inheritedRoles" yaml:"inheritedRoles"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

