package awscdkresourcesmongodbatlas


// Creates one API key for the specified organization.
//
// An organization API key grants programmatic access to an organization.
type CfnApiKeyProps struct {
	// Name of the AWS Secrets Manager secret that stores the API key Details.
	AwsSecretName *string `field:"required" json:"awsSecretName" yaml:"awsSecretName"`
	// Purpose or explanation provided when someone created this organization API key.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
	//
	// Use the /orgs endpoint to retrieve all organizations to which the authenticated user has access.
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// ARN of the AWS Secrets Manager secret that stores the API key Details.
	AwsSecretArn *string `field:"optional" json:"awsSecretArn" yaml:"awsSecretArn"`
	ListOptions *ListOptions `field:"optional" json:"listOptions" yaml:"listOptions"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	ProjectAssignments *[]*ProjectAssignment `field:"optional" json:"projectAssignments" yaml:"projectAssignments"`
	// List of roles to grant this API key.
	//
	// If you provide this list, provide a minimum of one role and ensure each role applies to this organization.
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
}

