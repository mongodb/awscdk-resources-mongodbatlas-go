package awscdkresourcesmongodbatlas


// Returns, adds, and edits organizational units in MongoDB Cloud.
type CfnOrganizationProps struct {
	// AwsSecretName used to set newly created Org credentials information.
	AwsSecretName *string `field:"required" json:"awsSecretName" yaml:"awsSecretName"`
	// Human-readable label that identifies the organization.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user that you want to assign the Organization Owner role.
	//
	// This user must be a member of the same organization as the calling API key. If you provide federationSettingsId, this user must instead have the Organization Owner role on an organization in the specified federation. This parameter is required only when you authenticate with Programmatic API Keys.
	OrgOwnerId *string `field:"required" json:"orgOwnerId" yaml:"orgOwnerId"`
	ApiKey *ApiKey `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Unique 24-hexadecimal digit string that identifies the federation to link the newly created organization to.
	//
	// If specified, the proposed Organization Owner of the new organization must have the Organization Owner role in an organization associated with the federation.
	FederatedSettingsId *string `field:"optional" json:"federatedSettingsId" yaml:"federatedSettingsId"`
	// Flag that indicates whether this organization has been deleted.
	IsDeleted *bool `field:"optional" json:"isDeleted" yaml:"isDeleted"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

