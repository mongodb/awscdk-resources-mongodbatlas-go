package awscdkresourcesmongodbatlas


// Creates and manages a Service Account for an organization.
type CfnServiceAccountProps struct {
	// Human readable description for the Service Account.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Human-readable name for the Service Account.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Unique 24-hexadecimal digit string that identifies the organization.
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// List of organization-level roles for the Service Account.
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
	// Expiration time of the new Service Account secret in hours.
	SecretExpiresAfterHours *float64 `field:"required" json:"secretExpiresAfterHours" yaml:"secretExpiresAfterHours"`
	// Profile used to provide credentials information.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// List of secrets associated with the Service Account.
	Secrets *[]*Secret `field:"optional" json:"secrets" yaml:"secrets"`
}

