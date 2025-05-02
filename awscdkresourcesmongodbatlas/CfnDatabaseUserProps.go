package awscdkresourcesmongodbatlas


// Returns, adds, edits, and removes database users.
type CfnDatabaseUserProps struct {
	// MongoDB database against which the MongoDB database user authenticates.
	//
	// MongoDB database users must provide both a username and authentication database to log into MongoDB.  Default value is `admin`.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Unique 24-hexadecimal digit string that identifies your Atlas Project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// List that provides the pairings of one role with one applicable database.
	Roles *[]*RoleDefinition `field:"required" json:"roles" yaml:"roles"`
	// Human-readable label that represents the user that authenticates to MongoDB.
	//
	// The format of this label depends on the method of authentication. This will be USER_ARN or ROLE_ARN if AWSIAMType is USER or ROLE. Refer https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Database-Users/operation/createDatabaseUser for details.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Human-readable label that indicates whether the new database user authenticates with the Amazon Web Services (AWS) Identity and Access Management (IAM) credentials associated with the user or the user's role.
	//
	// Default value is `NONE`.
	AwsiamType CfnDatabaseUserPropsAwsiamType `field:"optional" json:"awsiamType" yaml:"awsiamType"`
	// Date and time when MongoDB Cloud deletes the user.
	//
	// This parameter expresses its value in the ISO 8601 timestamp format in UTC and can include the time zone designation. You must specify a future date that falls within one week of making the Application Programming Interface (API) request.
	DeleteAfterDate *string `field:"optional" json:"deleteAfterDate" yaml:"deleteAfterDate"`
	// Description of this database user.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List that contains the key-value pairs for tagging and categorizing the MongoDB database user.
	//
	// The labels that you define do not appear in the console.
	Labels *[]*LabelDefinition `field:"optional" json:"labels" yaml:"labels"`
	// Method by which the provided username is authenticated.
	//
	// Default value is `NONE`.
	LdapAuthType CfnDatabaseUserPropsLdapAuthType `field:"optional" json:"ldapAuthType" yaml:"ldapAuthType"`
	// The user’s password.
	//
	// This field is not included in the entity returned from the server.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided `default` is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// List that contains clusters and MongoDB Atlas Data Lakes that this database user can access.
	//
	// If omitted, MongoDB Cloud grants the database user access to all the clusters and MongoDB Atlas Data Lakes in the project.
	Scopes *[]*ScopeDefinition `field:"optional" json:"scopes" yaml:"scopes"`
	// Method that briefs who owns the certificate provided.
	//
	// Default value is `NONE`.
	X509Type CfnDatabaseUserPropsX509Type `field:"optional" json:"x509Type" yaml:"x509Type"`
}

