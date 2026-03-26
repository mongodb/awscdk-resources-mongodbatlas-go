package awscdkresourcesmongodbatlas


// Creates a secret for the specified Service Account at the organization level.
type CfnServiceAccountSecretProps struct {
	// The Client ID of the Service Account.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// The expiration time of the new Service Account secret, provided in hours.
	//
	// The minimum and maximum allowed expiration times are subject to change and are controlled by the organization's settings.
	SecretExpiresAfterHours *float64 `field:"optional" json:"secretExpiresAfterHours" yaml:"secretExpiresAfterHours"`
}

