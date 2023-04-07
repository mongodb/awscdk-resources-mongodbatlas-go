// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type DatadogIntegrationProps struct {
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Atlas API keys.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Key that allows MongoDB Cloud to access your Datadog account.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// Two-letter code that indicates which regional URL MongoDB uses to access the Datadog API.
	Region DatadogRegion `field:"required" json:"region" yaml:"region"`
}

