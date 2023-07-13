package awscdkresourcesmongodbatlas


// Returns, adds, edits, and removes network access limits to database deployments in MongoDB Cloud.
type CfnProjectIpAccessListProps struct {
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	AccessList *[]*AccessListDefinition `field:"optional" json:"accessList" yaml:"accessList"`
	ListOptions *ListOptions `field:"optional" json:"listOptions" yaml:"listOptions"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Number of documents returned in this response.
	TotalCount *float64 `field:"optional" json:"totalCount" yaml:"totalCount"`
}

