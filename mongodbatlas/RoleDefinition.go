// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type RoleDefinition struct {
	CollectionName *string `field:"optional" json:"collectionName" yaml:"collectionName"`
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
}

