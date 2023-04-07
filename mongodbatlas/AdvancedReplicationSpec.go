// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// List of settings that configure your cluster regions.
//
// For Global Clusters, each object in the array represents a zone where your clusters nodes deploy. For non-Global replica sets and sharded clusters, this array has one object representing where your clusters nodes deploy.
type AdvancedReplicationSpec struct {
	// Hardware specifications for nodes set for a given region.
	//
	// Each regionConfigs object describes the region's priority in elections and the number and type of MongoDB nodes that MongoDB Cloud deploys to the region. Each regionConfigs object must have either an analyticsSpecs object, electableSpecs object, or readOnlySpecs object. Tenant clusters only require electableSpecs. Dedicated clusters can specify any of these specifications, but must have at least one electableSpecs object within a replicationSpec. Every hardware specification must use the same instanceSize.
	//
	// Example:
	//
	// If you set "replicationSpecs[n].regionConfigs[m].analyticsSpecs.instanceSize" : "M30", set "replicationSpecs[n].regionConfigs[m].electableSpecs.instanceSize" : "M30"if you have electable nodes and"replicationSpecs[n].regionConfigs[m].readOnlySpecs.instanceSize" : "M30" if you have read-only nodes.",
	AdvancedRegionConfigs *[]*AdvancedRegionConfig `field:"optional" json:"advancedRegionConfigs" yaml:"advancedRegionConfigs"`
	// Unique 24-hexadecimal digit string that identifies the replication object for a zone in a Multi-Cloud Cluster.
	//
	// If you include existing zones in the request, you must specify this parameter. If you add a new zone to an existing Multi-Cloud Cluster, you may specify this parameter. The request deletes any existing zones in the Multi-Cloud Cluster that you exclude from the request.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Positive integer that specifies the number of shards to deploy in each specified zone.
	//
	// If you set this value to 1 and "clusterType" : "SHARDED", MongoDB Cloud deploys a single-shard sharded cluster. Don't create a sharded cluster with a single shard for production environments. Single-shard sharded clusters don't provide the same benefits as multi-shard configurations.
	NumShards *float64 `field:"optional" json:"numShards" yaml:"numShards"`
	// Human-readable label that identifies the zone in a Global Cluster.
	//
	// Provide this value only if "clusterType" : "GEOSHARDED".
	ZoneName *string `field:"optional" json:"zoneName" yaml:"zoneName"`
}

