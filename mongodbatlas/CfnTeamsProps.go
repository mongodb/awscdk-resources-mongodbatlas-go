// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Adds one team to the specified project.
//
// All members of the team share the same project access. To use this resource, the requesting API Key must have the Project User Admin role. This resource doesn't require the API Key to have an Access List.
type CfnTeamsProps struct {
	// Unique 24-hexadecimal character string that identifies the team.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Unique 24-hexadecimal character string that identifies the team.
	OrgId *string `field:"optional" json:"orgId" yaml:"orgId"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Unique 24-hexadecimal character string that identifies the project.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// One or more organization- or project-level roles to assign to the MongoDB Cloud user.
	RoleNames *[]CfnTeamsPropsRoleNames `field:"optional" json:"roleNames" yaml:"roleNames"`
	// One or more organization- or project-level roles to assign to the MongoDB Cloud user.
	Usernames *[]*string `field:"optional" json:"usernames" yaml:"usernames"`
}

