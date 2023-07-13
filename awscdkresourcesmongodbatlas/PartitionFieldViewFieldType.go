package awscdkresourcesmongodbatlas


// Data type of the parameter that that MongoDB Cloud uses to partition data.
//
// Partition parameters of type [UUID](http://bsonspec.org/spec.html) must be of binary subtype 4. MongoDB Cloud skips partition parameters of type UUID with subtype 3.
type PartitionFieldViewFieldType string

const (
	// date.
	PartitionFieldViewFieldType_DATE PartitionFieldViewFieldType = "DATE"
	// int.
	PartitionFieldViewFieldType_INT PartitionFieldViewFieldType = "INT"
	// long.
	PartitionFieldViewFieldType_LONG PartitionFieldViewFieldType = "LONG"
	// objectId.
	PartitionFieldViewFieldType_OBJECT_ID PartitionFieldViewFieldType = "OBJECT_ID"
	// string.
	PartitionFieldViewFieldType_STRING PartitionFieldViewFieldType = "STRING"
	// uuid.
	PartitionFieldViewFieldType_UUID PartitionFieldViewFieldType = "UUID"
)

