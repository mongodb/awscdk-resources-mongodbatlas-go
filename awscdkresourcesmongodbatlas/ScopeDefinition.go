// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type ScopeDefinition struct {
	Name *string `field:"optional" json:"name" yaml:"name"`
	Type ScopeDefinitionType `field:"optional" json:"type" yaml:"type"`
}

