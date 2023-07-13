package awscdkresourcesmongodbatlas


type PagerDutyIntegrationProps struct {
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Atlas API keys.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// PagerDuty region that indicates the API Uniform Resource Locator (URL) to use.
	Region PagerDutyRegion `field:"required" json:"region" yaml:"region"`
	// Service key associated with your PagerDuty account.
	ServiceKey *string `field:"required" json:"serviceKey" yaml:"serviceKey"`
}

