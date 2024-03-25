package awscdkresourcesmongodbatlas


// Type of the connection.
//
// Can be either Cluster, Kafka, or Sample.
type CfnStreamConnectionPropsType string

const (
	// Kafka.
	CfnStreamConnectionPropsType_KAFKA CfnStreamConnectionPropsType = "KAFKA"
	// Cluster.
	CfnStreamConnectionPropsType_CLUSTER CfnStreamConnectionPropsType = "CLUSTER"
	// Sample.
	CfnStreamConnectionPropsType_SAMPLE CfnStreamConnectionPropsType = "SAMPLE"
)

