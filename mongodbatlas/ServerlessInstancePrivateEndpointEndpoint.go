// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type ServerlessInstancePrivateEndpointEndpoint struct {
	// Unique provider identifier of the private endpoint.
	EndpointId *string `field:"optional" json:"endpointId" yaml:"endpointId"`
	// Cloud provider where the private endpoint is deployed.
	ProviderName *string `field:"optional" json:"providerName" yaml:"providerName"`
	// Region where the private endpoint is deployed.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

