package awscdkresourcesmongodbatlas


type PartitionFieldView struct {
	// Human-readable label that identifies the parameter that MongoDB Cloud uses to partition data.
	//
	// To specify a nested parameter, use the dot notation.
	FieldName *string `field:"optional" json:"fieldName" yaml:"fieldName"`
	// Data type of the parameter that that MongoDB Cloud uses to partition data.
	//
	// Partition parameters of type [UUID](http://bsonspec.org/spec.html) must be of binary subtype 4. MongoDB Cloud skips partition parameters of type UUID with subtype 3.
	FieldType PartitionFieldViewFieldType `field:"optional" json:"fieldType" yaml:"fieldType"`
	// Sequence in which MongoDB Cloud slices the collection data to create partitions.
	//
	// The resource expresses this sequence starting with zero. The value of the **criteria.dateField** parameter defaults as the first item in the partition sequence.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
}

