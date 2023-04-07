// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// The Private Endpoint creation flow consists of the creation of three related resources in the next order: 1.
//
// Atlas Private Endpoint Service 2. Aws VPC private Endpoint 3. Atlas Private Endpoint
type CfnPrivateEndpointProps struct {
	// Unique 24-hexadecimal digit string that identifies your project.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// Aws Region.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Name of the AWS PrivateLink endpoint service.
	//
	// Atlas returns null while it is creating the endpoint service.
	EndpointServiceName *string `field:"optional" json:"endpointServiceName" yaml:"endpointServiceName"`
	// Error message pertaining to the AWS PrivateLink connection.
	//
	// Returns null if there are no errors.
	ErrorMessage *string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
	// List of private endpoint associated to the service.
	PrivateEndpoints *[]*PrivateEndpoint `field:"optional" json:"privateEndpoints" yaml:"privateEndpoints"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup (../../../examples/profile-secret.yaml)
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Status of the Atlas PrivateEndpoint service connection.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

