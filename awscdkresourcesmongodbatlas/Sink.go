package awscdkresourcesmongodbatlas


// Ingestion destination of a Data Lake Pipeline.
type Sink struct {
	// Target cloud provider for this Data Lake Pipeline.
	MetadataProvider *string `field:"optional" json:"metadataProvider" yaml:"metadataProvider"`
	// Target cloud provider region for this Data Lake Pipeline.
	MetadataRegion *string `field:"optional" json:"metadataRegion" yaml:"metadataRegion"`
	// Ordered fields used to physically organize data in the destination.
	PartitionFields *[]*PartitionFields `field:"optional" json:"partitionFields" yaml:"partitionFields"`
	// Type of ingestion destination of this Data Lake Pipeline.
	Type SinkType `field:"optional" json:"type" yaml:"type"`
}

