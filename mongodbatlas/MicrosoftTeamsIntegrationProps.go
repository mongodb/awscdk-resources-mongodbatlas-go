// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type MicrosoftTeamsIntegrationProps struct {
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Atlas API keys.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Endpoint web address of the Microsoft Teams webhook to which MongoDB Cloud sends notifications.
	MicrosoftTeamsWebhookUrl *string `field:"required" json:"microsoftTeamsWebhookUrl" yaml:"microsoftTeamsWebhookUrl"`
}

