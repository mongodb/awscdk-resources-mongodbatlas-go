package awscdkresourcesmongodbatlas


// Retrieves or creates projects in any given Atlas organization.
type CfnProjectProps struct {
	// Name of the project to create.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Unique identifier of the organization within which to create the project.
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	ProjectApiKeys *[]*ProjectApiKey `field:"optional" json:"projectApiKeys" yaml:"projectApiKeys"`
	ProjectSettings *ProjectSettings `field:"optional" json:"projectSettings" yaml:"projectSettings"`
	ProjectTeams *[]*ProjectTeam `field:"optional" json:"projectTeams" yaml:"projectTeams"`
	// Flag that indicates whether to create the project with default alert settings.
	WithDefaultAlertsSettings *bool `field:"optional" json:"withDefaultAlertsSettings" yaml:"withDefaultAlertsSettings"`
}

