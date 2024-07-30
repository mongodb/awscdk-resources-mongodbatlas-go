package awscdkresourcesmongodbatlas


// Returns, adds, and edits organizational units in MongoDB Cloud.
type CfnOrgInvitationProps struct {
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"required" json:"profile" yaml:"profile"`
	// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
	IncludeCount *bool `field:"optional" json:"includeCount" yaml:"includeCount"`
	// Unique 24-hexadecimal digit string that identifies the invitation.
	InvitationId *string `field:"optional" json:"invitationId" yaml:"invitationId"`
	// Number of items that the response returns per page.
	ItemsPerPage *float64 `field:"optional" json:"itemsPerPage" yaml:"itemsPerPage"`
	// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
	OrgId *string `field:"optional" json:"orgId" yaml:"orgId"`
	// Human-readable label that identifies this organization.
	OrgName *string `field:"optional" json:"orgName" yaml:"orgName"`
	// Number of the page that displays the current set of the total objects that the response returns.
	PageNum *float64 `field:"optional" json:"pageNum" yaml:"pageNum"`
	// One or more organization or project level roles to assign to the MongoDB Cloud user.
	Roles *[]CfnOrgInvitationPropsRoles `field:"optional" json:"roles" yaml:"roles"`
	// List of unique 24-hexadecimal digit strings that identifies each team.
	TeamIds *[]*string `field:"optional" json:"teamIds" yaml:"teamIds"`
	// Number of documents returned in this response.
	TotalCount *float64 `field:"optional" json:"totalCount" yaml:"totalCount"`
	// Email address of the MongoDB Cloud user invited to join the organization.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

