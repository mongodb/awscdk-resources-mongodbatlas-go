// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type IpAccessListProps struct {
	AccessList *[]*AccessListDefinition `field:"required" json:"accessList" yaml:"accessList"`
	ListOptions *ListOptions `field:"optional" json:"listOptions" yaml:"listOptions"`
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	TotalCount *float64 `field:"optional" json:"totalCount" yaml:"totalCount"`
}

