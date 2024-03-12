package awscdkresourcesmongodbatlas


// Properties for the secure transport connection to Kafka.
//
// For SSL, this can include the trusted certificate to use.
type StreamsKafkaSecurity struct {
	// A trusted, public x509 certificate for connecting to Kafka over SSL.
	BrokerPublicCertificate *string `field:"optional" json:"brokerPublicCertificate" yaml:"brokerPublicCertificate"`
	// Describes the transport type.
	//
	// Can be either PLAINTEXT or SSL.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

