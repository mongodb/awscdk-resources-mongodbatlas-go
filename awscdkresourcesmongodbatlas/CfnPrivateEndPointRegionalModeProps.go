package awscdkresourcesmongodbatlas


// An example resource schema demonstrating some basic constructs and validation rules.
type CfnPrivateEndPointRegionalModeProps struct {
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

