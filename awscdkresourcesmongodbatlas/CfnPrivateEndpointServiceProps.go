package awscdkresourcesmongodbatlas


// Creates one private endpoint service for the specified cloud service provider.
//
// This cloud service provider manages the private endpoint service for the project. When you create a private endpoint service, MongoDB Cloud creates a network container in the project for the cloud provider for which you create the private endpoint service if one doesn't already exist.
type CfnPrivateEndpointServiceProps struct {
	// Cloud service provider that manages this private endpoint, default : AWS.
	CloudProvider CfnPrivateEndpointServicePropsCloudProvider `field:"required" json:"cloudProvider" yaml:"cloudProvider"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Aws Region.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup (../../../examples/profile-secret.yaml)
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

