package awscdkresourcesmongodbatlas


// User credentials required to connect to a Kafka Cluster.
//
// Includes the authentication type, as well as the parameters for that authentication mode.
type StreamsKafkaAuthentication struct {
	// Style of authentication.
	//
	// Can be one of PLAIN, SCRAM-256, or SCRAM-512.
	Mechanism *string `field:"optional" json:"mechanism" yaml:"mechanism"`
	// Password of the account to connect to the Kafka cluster.
	//
	// Review [AWS security best practices for CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/security-best-practices.html#creds) to manage credentials.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Username of the account to connect to the Kafka cluster.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

