package awscdkresourcesmongodbatlas


// Adds one private endpoint for Federated Database Instances and Online Archives to the specified projects.
type CfnPrivatelinkEndpointServiceDataFederationOnlineArchiveProps struct {
	// Unique 22-character alphanumeric string that identifies the private endpoint.Reg ex ^vpce-[0-9a-f]{17}$ .
	//
	// Atlas Data Lake supports Amazon Web Services private endpoints using the AWS PrivateLink feature.
	EndpointId *string `field:"required" json:"endpointId" yaml:"endpointId"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Human-readable string to associate with this private endpoint.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Human-readable label that identifies the resource type associated with this private endpoint.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

