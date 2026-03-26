package awscdkresourcesmongodbatlas


// Creates one private endpoint for the specified cloud service provider.
//
// At this current version only AWS is supported.
type CfnPrivateEndpointAwsProps struct {
	// Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to create a private endpoint.
	EndpointServiceId *string `field:"required" json:"endpointServiceId" yaml:"endpointServiceId"`
	// Unique string that identifies the private endpoint.
	//
	// for AWS is the VPC endpoint ID, example: vpce-xxxxxxxx.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// If this proper is set to TRUE, the cloud formation resource will return success Only if the private connection is Succeeded.
	EnforceConnectionSuccess *bool `field:"optional" json:"enforceConnectionSuccess" yaml:"enforceConnectionSuccess"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml)
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

