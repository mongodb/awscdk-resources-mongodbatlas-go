package awscdkresourcesmongodbatlas


// Atlas Resource Policies.
type CfnResourcePolicyProps struct {
	// Human-readable label that describes the atlas resource policy.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
	//
	// Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// List of policies that make up the atlas resource policy.
	Policies *[]*ApiAtlasPolicy `field:"required" json:"policies" yaml:"policies"`
	// The user that last updated the atlas resource policy.
	CreatedByUser *ApiAtlasUserMetadata `field:"optional" json:"createdByUser" yaml:"createdByUser"`
	// Description of the Atlas resource policy.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The user that last updated the atlas resource policy.
	LastUpdatedByUser *ApiAtlasUserMetadata `field:"optional" json:"lastUpdatedByUser" yaml:"lastUpdatedByUser"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

