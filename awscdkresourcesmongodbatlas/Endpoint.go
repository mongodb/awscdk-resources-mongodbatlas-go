// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type Endpoint struct {
	// Unique string that the cloud provider uses to identify the private endpoint.
	EndpointId *string `field:"optional" json:"endpointId" yaml:"endpointId"`
	// Cloud provider in which MongoDB Cloud deploys the private endpoint.
	ProviderName *string `field:"optional" json:"providerName" yaml:"providerName"`
	// Region in which MongoDB Cloud deploys the private endpoint.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

