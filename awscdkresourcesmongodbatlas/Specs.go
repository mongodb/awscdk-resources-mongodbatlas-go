package awscdkresourcesmongodbatlas


type Specs struct {
	// Target throughput desired for storage attached to your AWS-provisioned cluster. Only change this parameter if you:.
	//
	// set "replicationSpecs[n].regionConfigs[m].providerName" : "AWS".
	// set "replicationSpecs[n].regionConfigs[m].electableSpecs.instanceSize" : "M30" or greater not including Mxx_NVME tiers.
	// The maximum input/output operations per second (IOPS) depend on the selected .instanceSize and .diskSizeGB. This parameter defaults to the cluster tier's standard IOPS value. Changing this value impacts cluster cost. MongoDB Cloud enforces minimum ratios of storage capacity to system memory for given cluster tiers. This keeps cluster performance consistent with large datasets.
	//
	// Instance sizes M10 to M40 have a ratio of disk capacity to system memory of 60:1.
	// Instance sizes greater than M40 have a ratio of 120:1.
	DiskIops *string `field:"optional" json:"diskIops" yaml:"diskIops"`
	// Type of storage you want to attach to your AWS-provisioned cluster.
	//
	// STANDARD volume types can't exceed the default input/output operations per second (IOPS) rate for the selected volume size.
	//
	// PROVISIONED volume types must fall within the allowable IOPS range for the selected volume size."
	EbsVolumeType *string `field:"optional" json:"ebsVolumeType" yaml:"ebsVolumeType"`
	// Hardware specification for the instance sizes in this region.
	//
	// Each instance size has a default storage and memory capacity. The instance size you select applies to all the data-bearing hosts in your instance size. If you deploy a Global Cluster, you must choose a instance size of M30 or greater.
	InstanceSize *string `field:"optional" json:"instanceSize" yaml:"instanceSize"`
	// Number of read-only nodes for MongoDB Cloud deploys to the region.
	//
	// Read-only nodes can never become the primary, but can enable local reads.
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
}

