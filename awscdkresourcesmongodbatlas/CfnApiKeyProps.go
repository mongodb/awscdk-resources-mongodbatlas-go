package awscdkresourcesmongodbatlas


// Creates one API key for the specified organization.
//
// An organization API key grants programmatic access to an organization.
type CfnApiKeyProps struct {
	// Purpose or explanation provided when someone created this organization API key.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
	//
	// Use the /orgs endpoint to retrieve all organizations to which the authenticated user has access.
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// Name of the AWS Secrets Manager secret that stores the API key Details.
	//
	// The secret name must be unique to the AWS account and region. If you don't specify a name, AWS CloudFormation use APIUserId for the secret name.
	AwsSecretName *string `field:"optional" json:"awsSecretName" yaml:"awsSecretName"`
	ListOptions *ListOptions `field:"optional" json:"listOptions" yaml:"listOptions"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	ProjectAssignments *[]*ProjectAssignment `field:"optional" json:"projectAssignments" yaml:"projectAssignments"`
	// List of roles to grant this API key.
	//
	// If you provide this list, provide a minimum of one role and ensure each role applies to this organization.
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
}

