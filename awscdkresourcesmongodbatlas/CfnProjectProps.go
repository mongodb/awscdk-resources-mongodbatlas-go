package awscdkresourcesmongodbatlas


// Retrieves or creates projects in any given Atlas organization.
type CfnProjectProps struct {
	// Name of the project to create.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Unique identifier of the organization within which to create the project.
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// The number of Atlas clusters deployed in the project.
	ClusterCount *float64 `field:"optional" json:"clusterCount" yaml:"clusterCount"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// API keys that you assigned to the specified project.
	ProjectApiKeys *[]*ProjectApiKey `field:"optional" json:"projectApiKeys" yaml:"projectApiKeys"`
	ProjectSettings *ProjectSettings `field:"optional" json:"projectSettings" yaml:"projectSettings"`
	// Teams to which the authenticated user has access in the project specified using its unique 24-hexadecimal digit identifier.
	ProjectTeams *[]*ProjectTeam `field:"optional" json:"projectTeams" yaml:"projectTeams"`
	// Region usage restrictions that designate the project's AWS region.Enum: "GOV_REGIONS_ONLY" "COMMERCIAL_FEDRAMP_REGIONS_ONLY" "NONE".
	RegionUsageRestrictions *string `field:"optional" json:"regionUsageRestrictions" yaml:"regionUsageRestrictions"`
	// Flag that indicates whether to create the project with default alert settings.
	WithDefaultAlertsSettings *bool `field:"optional" json:"withDefaultAlertsSettings" yaml:"withDefaultAlertsSettings"`
}

