package awscdkresourcesmongodbatlas


// Returns, adds, edits, and removes private endpoints for serverless instances.
//
// To learn more, see the Atlas Administration API tab on the following tutorial.
type CfnServerlessPrivateEndpointProps struct {
	// Human-readable label that identifies the serverless instance for which the tenant endpoint will be created.
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Properties used to configure Aws private endpoint.
	AwsPrivateEndpointConfigurationProperties *AwsPrivateEndpointConfig `field:"optional" json:"awsPrivateEndpointConfigurationProperties" yaml:"awsPrivateEndpointConfigurationProperties"`
	// Unique string that identifies the private endpoint's network interface.
	CloudProviderEndpointId *string `field:"optional" json:"cloudProviderEndpointId" yaml:"cloudProviderEndpointId"`
	// Human-readable comment associated with the private endpoint.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// If true the resource will create the aws private endpoint and assign the Endpoint ID.
	CreateAndAssignAwsPrivateEndpoint *bool `field:"optional" json:"createAndAssignAwsPrivateEndpoint" yaml:"createAndAssignAwsPrivateEndpoint"`
	// IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service.
	PrivateEndpointIpAddress *string `field:"optional" json:"privateEndpointIpAddress" yaml:"privateEndpointIpAddress"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

