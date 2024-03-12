package awscdkresourcesmongodbatlas


// Settings that define a connection to an external data store.
type StreamsConnection struct {
	Authentication *StreamsKafkaAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Comma separated list of server addresses.
	BootstrapServers *string `field:"optional" json:"bootstrapServers" yaml:"bootstrapServers"`
	// Name of the cluster configured for this connection.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	DbRoleToExecute *DbRoleToExecute `field:"optional" json:"dbRoleToExecute" yaml:"dbRoleToExecute"`
	// Human-readable label that identifies the stream connection.
	Name *string `field:"optional" json:"name" yaml:"name"`
	Security *StreamsKafkaSecurity `field:"optional" json:"security" yaml:"security"`
	// Type of the connection.
	//
	// Can be either Cluster or Kafka.
	Type StreamsConnectionType `field:"optional" json:"type" yaml:"type"`
}

