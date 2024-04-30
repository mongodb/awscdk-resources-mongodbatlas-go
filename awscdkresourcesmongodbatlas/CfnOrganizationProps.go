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
	// Flag that indicates whether to require API operations to originate from an IP Address added to the API access list for the specified organization.
	ApiAccessListRequired *bool `field:"optional" json:"apiAccessListRequired" yaml:"apiAccessListRequired"`
	ApiKey *ApiKey `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Unique 24-hexadecimal digit string that identifies the federation to link the newly created organization to.
	//
	// If specified, the proposed Organization Owner of the new organization must have the Organization Owner role in an organization associated with the federation.
	FederatedSettingsId *string `field:"optional" json:"federatedSettingsId" yaml:"federatedSettingsId"`
	// Flag that indicates whether this organization has been deleted.
	IsDeleted *bool `field:"optional" json:"isDeleted" yaml:"isDeleted"`
	// Flag that indicates whether to require users to set up Multi-Factor Authentication (MFA) before accessing the specified organization.
	//
	// To learn more, see: https://www.mongodb.com/docs/atlas/security-multi-factor-authentication/.
	MultiFactorAuthRequired *bool `field:"optional" json:"multiFactorAuthRequired" yaml:"multiFactorAuthRequired"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Flag that indicates whether to block MongoDB Support from accessing Atlas infrastructure for any deployment in the specified organization without explicit permission.
	//
	// Once this setting is turned on, you can grant MongoDB Support a 24-hour bypass access to the Atlas deployment to resolve support issues. To learn more, see: https://www.mongodb.com/docs/atlas/security-restrict-support-access/.
	RestrictEmployeeAccess *bool `field:"optional" json:"restrictEmployeeAccess" yaml:"restrictEmployeeAccess"`
}

