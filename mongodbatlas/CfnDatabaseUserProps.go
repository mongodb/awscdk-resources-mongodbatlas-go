// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Returns, adds, edits, and removes database users.
type CfnDatabaseUserProps struct {
	// MongoDB database against which the MongoDB database user authenticates.
	//
	// MongoDB database users must provide both a username and authentication database to log into MongoDB.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Unique identifier of the Atlas project to which the user belongs.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// List that provides the pairings of one role with one applicable database.
	Roles *[]*RoleDefinition `field:"required" json:"roles" yaml:"roles"`
	// Human-readable label that represents the user that authenticates to MongoDB.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Human-readable label that indicates whether the new database user authenticates with the Amazon Web Services (AWS) Identity and Access Management (IAM) credentials associated with the user or the user's role.
	AwsiamType CfnDatabaseUserPropsAwsiamType `field:"optional" json:"awsiamType" yaml:"awsiamType"`
	// Date and time when MongoDB Cloud deletes the user.
	//
	// This parameter expresses its value in the ISO 8601 timestamp format in UTC and can include the time zone designation. You must specify a future date that falls within one week of making the Application Programming Interface (API) request.
	DeleteAfterDate *string `field:"optional" json:"deleteAfterDate" yaml:"deleteAfterDate"`
	// List that contains the key-value pairs for tagging and categorizing the MongoDB database user.
	//
	// The labels that you define do not appear in the console.
	Labels *[]*LabelDefinition `field:"optional" json:"labels" yaml:"labels"`
	// Method by which the provided username is authenticated.
	//
	// If no value is given, Atlas uses the default value of NONE.
	LdapAuthType CfnDatabaseUserPropsLdapAuthType `field:"optional" json:"ldapAuthType" yaml:"ldapAuthType"`
	// The user’s password.
	//
	// This field is not included in the entity returned from the server.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// List that contains clusters and MongoDB Atlas Data Lakes that this database user can access.
	//
	// If omitted, MongoDB Cloud grants the database user access to all the clusters and MongoDB Atlas Data Lakes in the project.
	Scopes *[]*ScopeDefinition `field:"optional" json:"scopes" yaml:"scopes"`
	// Method that briefs who owns the certificate provided.
	//
	// If no value is given while using X509Type, Atlas uses the default value of MANAGED.
	X509Type CfnDatabaseUserPropsX509Type `field:"optional" json:"x509Type" yaml:"x509Type"`
}

