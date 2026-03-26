package awscdkresourcesmongodbatlas


// Creates a secret for the specified Service Account at the project level.
type CfnProjectServiceAccountSecretProps struct {
	// The Client ID of the Service Account.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// The expiration time of the new Service Account secret, provided in hours.
	//
	// The minimum and maximum allowed expiration times are subject to change and are controlled by the organization's settings.
	SecretExpiresAfterHours *float64 `field:"optional" json:"secretExpiresAfterHours" yaml:"secretExpiresAfterHours"`
}

