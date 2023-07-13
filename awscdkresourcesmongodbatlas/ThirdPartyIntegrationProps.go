package awscdkresourcesmongodbatlas


type ThirdPartyIntegrationProps struct {
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Atlas API keys.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

