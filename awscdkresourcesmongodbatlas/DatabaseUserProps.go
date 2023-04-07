// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type DatabaseUserProps struct {
	AwsiamType CfnDatabaseUserPropsAwsiamType `field:"optional" json:"awsiamType" yaml:"awsiamType"`
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	DeleteAfterDate *string `field:"optional" json:"deleteAfterDate" yaml:"deleteAfterDate"`
	Labels *[]*LabelDefinition `field:"optional" json:"labels" yaml:"labels"`
	LdapAuthType CfnDatabaseUserPropsLdapAuthType `field:"optional" json:"ldapAuthType" yaml:"ldapAuthType"`
	Password *string `field:"optional" json:"password" yaml:"password"`
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	Roles *[]*RoleDefinition `field:"optional" json:"roles" yaml:"roles"`
	Scopes *[]*ScopeDefinition `field:"optional" json:"scopes" yaml:"scopes"`
	Username *string `field:"optional" json:"username" yaml:"username"`
	X509Type CfnDatabaseUserPropsX509Type `field:"optional" json:"x509Type" yaml:"x509Type"`
}

