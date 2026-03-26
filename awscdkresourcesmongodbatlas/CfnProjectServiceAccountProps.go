package awscdkresourcesmongodbatlas


// Creates and manages a Project Service Account.
//
// This resource creates a Service Account and assigns it to a project. Note: Deleting this resource only unassigns the Service Account from the project, but doesn't delete it from the organization.
type CfnProjectServiceAccountProps struct {
	// Human readable description for the Service Account.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Human-readable name for the Service Account.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// A list of project-level roles for the Service Account.
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
	// Expiration time of the new Service Account secret in hours.
	SecretExpiresAfterHours *float64 `field:"required" json:"secretExpiresAfterHours" yaml:"secretExpiresAfterHours"`
	// Profile used to provide credentials information.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// List of secrets associated with the Service Account.
	Secrets *[]*SecretDefinition `field:"optional" json:"secrets" yaml:"secrets"`
}

