package awscdkresourcesmongodbatlas


// Ordered fields used to physically organize data in the destination.
type PartitionFields struct {
	// Human-readable label that identifies the field name used to partition data.
	FieldName *string `field:"optional" json:"fieldName" yaml:"fieldName"`
	// Sequence in which MongoDB Cloud slices the collection data to create partitions.
	//
	// The resource expresses this sequence starting with zero.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
}

