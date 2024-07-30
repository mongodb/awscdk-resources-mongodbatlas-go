package awscdkresourcesmongodbatlas


// Returns, adds, edits, and removes network access limits to database deployments in Atlas.
type CfnProjectIpAccessListProps struct {
	AccessList *[]*AccessListDefinition `field:"required" json:"accessList" yaml:"accessList"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	ListOptions *ListOptions `field:"optional" json:"listOptions" yaml:"listOptions"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

