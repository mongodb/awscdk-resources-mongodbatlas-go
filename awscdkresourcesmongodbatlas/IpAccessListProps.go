package awscdkresourcesmongodbatlas


type IpAccessListProps struct {
	AccessList *[]*AccessListDefinition `field:"required" json:"accessList" yaml:"accessList"`
	ListOptions *ListOptions `field:"optional" json:"listOptions" yaml:"listOptions"`
	// Default: allow-all.
	//
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	TotalCount *float64 `field:"optional" json:"totalCount" yaml:"totalCount"`
}

