package awscdkresourcesmongodbatlas


type ReadPreference struct {
	// Maximum replication lag, or staleness, for reads from secondaries.
	MaxStalenessSeconds *string `field:"optional" json:"maxStalenessSeconds" yaml:"maxStalenessSeconds"`
	// "primary" "primaryPreferred" "secondary" "secondaryPreferred" "nearest" Read preference mode that specifies to which replica set member to route the read requests.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// List that contains tag sets or tag specification documents.
	//
	// If specified, Atlas Data Lake routes read requests to replica set member or members that are associated with the specified tags.
	TagSets *[]*[]*TagSet `field:"optional" json:"tagSets" yaml:"tagSets"`
}

