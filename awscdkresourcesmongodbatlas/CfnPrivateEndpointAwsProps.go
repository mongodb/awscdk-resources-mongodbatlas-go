package awscdkresourcesmongodbatlas


// Creates one private endpoint for the specified cloud service provider.
//
// At this current version only AWS is supported.
type CfnPrivateEndpointAwsProps struct {
	// Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to create a private endpoint.
	EndpointServiceId *string `field:"required" json:"endpointServiceId" yaml:"endpointServiceId"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// State of the Amazon Web Service PrivateLink connection when MongoDB Cloud received this request.
	ConnectionStatus *string `field:"optional" json:"connectionStatus" yaml:"connectionStatus"`
	// If this proper is set to TRUE, the cloud formation resource will return success Only if the private connection is Succeeded.
	EnforceConnectionSuccess *bool `field:"optional" json:"enforceConnectionSuccess" yaml:"enforceConnectionSuccess"`
	// Error message returned when requesting private connection resource.
	//
	// The resource returns null if the request succeeded.
	ErrorMessage *string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
	// Unique string that identifies the private endpoint.
	//
	// for AWS is the VPC endpoint ID, example: vpce-xxxxxxxx.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml)
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

