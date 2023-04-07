// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Hardware specifications for nodes set for a given region.
//
// Each regionConfigs object describes the region's priority in elections and the number and type of MongoDB nodes that MongoDB Cloud deploys to the region. Each regionConfigs object must have either an analyticsSpecs object, electableSpecs object, or readOnlySpecs object. Tenant clusters only require electableSpecs. Dedicated clusters can specify any of these specifications, but must have at least one electableSpecs object within a replicationSpec. Every hardware specification must use the same instanceSize.
//
// Example:
//
// If you set "replicationSpecs[n].regionConfigs[m].analyticsSpecs.instanceSize" : "M30", set "replicationSpecs[n].regionConfigs[m].electableSpecs.instanceSize" : "M30"if you have electable nodes and"replicationSpecs[n].regionConfigs[m].readOnlySpecs.instanceSize" : "M30" if you have read-only nodes.",
type AdvancedRegionConfig struct {
	AnalyticsAutoScaling *AdvancedAutoScaling `field:"optional" json:"analyticsAutoScaling" yaml:"analyticsAutoScaling"`
	AnalyticsSpecs *Specs `field:"optional" json:"analyticsSpecs" yaml:"analyticsSpecs"`
	AutoScaling *AdvancedAutoScaling `field:"optional" json:"autoScaling" yaml:"autoScaling"`
	ElectableSpecs *Specs `field:"optional" json:"electableSpecs" yaml:"electableSpecs"`
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	ReadOnlySpecs *Specs `field:"optional" json:"readOnlySpecs" yaml:"readOnlySpecs"`
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

