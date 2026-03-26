package awscdkresourcesmongodbatlas


// Assigns a Service Account to a Project with specified roles in MongoDB Atlas.
type CfnServiceAccountProjectAssignmentProps struct {
	// The Client ID of the Service Account.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The Project permissions for the Service Account in the specified Project.
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
	// Unique 24-hexadecimal digit string that identifies the organization.
	//
	// Required for List operation.
	OrgId *string `field:"optional" json:"orgId" yaml:"orgId"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

