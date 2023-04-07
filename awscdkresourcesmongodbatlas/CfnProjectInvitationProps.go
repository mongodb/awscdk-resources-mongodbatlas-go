// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


// Returns, adds, and edits collections of clusters and users in MongoDB Cloud.
type CfnProjectInvitationProps struct {
	// Unique 24-hexadecimal digit string that identifies the invitation.
	InvitationId *string `field:"optional" json:"invitationId" yaml:"invitationId"`
	// Number of the page that displays the current set of the total objects that the response returns.
	PageNum *float64 `field:"optional" json:"pageNum" yaml:"pageNum"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// One or more organization or project level roles to assign to the MongoDB Cloud user.
	Roles *[]CfnProjectInvitationPropsRoles `field:"optional" json:"roles" yaml:"roles"`
	// Number of documents returned in this response.
	TotalCount *float64 `field:"optional" json:"totalCount" yaml:"totalCount"`
	// Email address of the user account invited to this project.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

