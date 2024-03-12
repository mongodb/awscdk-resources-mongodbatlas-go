package awscdkresourcesmongodbatlas


// Type of the connection.
//
// Can be either Cluster or Kafka.
type StreamsConnectionType string

const (
	// Kafka.
	StreamsConnectionType_KAFKA StreamsConnectionType = "KAFKA"
	// Cluster.
	StreamsConnectionType_CLUSTER StreamsConnectionType = "CLUSTER"
	// Sample.
	StreamsConnectionType_SAMPLE StreamsConnectionType = "SAMPLE"
)

