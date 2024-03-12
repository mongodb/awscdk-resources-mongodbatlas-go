package awscdkresourcesmongodbatlas


// Type of the connection.
//
// Can be either Cluster or Kafka.
type CfnStreamConnectionPropsType string

const (
	// Kafka.
	CfnStreamConnectionPropsType_KAFKA CfnStreamConnectionPropsType = "KAFKA"
	// Cluster.
	CfnStreamConnectionPropsType_CLUSTER CfnStreamConnectionPropsType = "CLUSTER"
)

