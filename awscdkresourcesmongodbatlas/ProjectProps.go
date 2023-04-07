// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type ProjectProps struct {
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	ClusterCount *float64 `field:"optional" json:"clusterCount" yaml:"clusterCount"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	ProjectApiKeys *[]*ProjectApiKey `field:"optional" json:"projectApiKeys" yaml:"projectApiKeys"`
	ProjectOwnerId *string `field:"optional" json:"projectOwnerId" yaml:"projectOwnerId"`
	ProjectSettings *ProjectSettings `field:"optional" json:"projectSettings" yaml:"projectSettings"`
	ProjectTeams *[]*ProjectTeam `field:"optional" json:"projectTeams" yaml:"projectTeams"`
	WithDefaultAlertsSettings *bool `field:"optional" json:"withDefaultAlertsSettings" yaml:"withDefaultAlertsSettings"`
}

