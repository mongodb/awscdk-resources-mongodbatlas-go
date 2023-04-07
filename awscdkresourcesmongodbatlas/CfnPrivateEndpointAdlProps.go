// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


// Adds one private endpoint for Federated Database Instances and Online Archives to the specified projects.
//
// To use this resource, the requesting API Key must have the Project Atlas Admin or Project Charts Admin roles. This resource doesn't require the API Key to have an Access List.
type CfnPrivateEndpointAdlProps struct {
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Human-readable label that identifies the cloud service provider.
	//
	// Atlas Data Lake supports Amazon Web Services only.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// Human-readable string to associate with this private endpoint.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Unique 22-character alphanumeric string that identifies the private endpoint.
	EndpointId *string `field:"optional" json:"endpointId" yaml:"endpointId"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Human-readable label that identifies the resource type associated with this private endpoint.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

