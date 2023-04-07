// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type PrivateEndpointProps struct {
	EndpointServiceName *string `field:"optional" json:"endpointServiceName" yaml:"endpointServiceName"`
	ErrorMessage *string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	PrivateEndpoints *[]*PrivateEndpoint `field:"optional" json:"privateEndpoints" yaml:"privateEndpoints"`
	Region *string `field:"optional" json:"region" yaml:"region"`
	Status *string `field:"optional" json:"status" yaml:"status"`
}

