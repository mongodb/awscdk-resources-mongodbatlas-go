package awscdkresourcesmongodbatlas


// User credentials required to connect to a Kafka Cluster.
//
// Includes the authentication type, as well as the parameters for that authentication mode.
type StreamsKafkaAuthentication struct {
	// OAuth client ID.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// OAuth client secret.
	//
	// Review [AWS security best practices for CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/security-best-practices.html#creds) to manage credentials.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Style of authentication.
	//
	// Can be one of PLAIN, SCRAM-256, SCRAM-512, or OAUTHBEARER.
	Mechanism *string `field:"optional" json:"mechanism" yaml:"mechanism"`
	// OAuth authentication method.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Password of the account to connect to the Kafka cluster.
	//
	// Review [AWS security best practices for CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/security-best-practices.html#creds) to manage credentials.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// SASL OAuth bearer extensions.
	SaslOauthbearerExtensions *string `field:"optional" json:"saslOauthbearerExtensions" yaml:"saslOauthbearerExtensions"`
	// OAuth scope.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// SSL certificate for client authentication to Kafka.
	SslCertificate *string `field:"optional" json:"sslCertificate" yaml:"sslCertificate"`
	// SSL key for client authentication to Kafka.
	//
	// Review [AWS security best practices for CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/security-best-practices.html#creds) to manage credentials.
	SslKey *string `field:"optional" json:"sslKey" yaml:"sslKey"`
	// Password for the SSL key, if it is password protected.
	//
	// Review [AWS security best practices for CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/security-best-practices.html#creds) to manage credentials.
	SslKeyPassword *string `field:"optional" json:"sslKeyPassword" yaml:"sslKeyPassword"`
	// OAuth token endpoint URL.
	TokenEndpointUrl *string `field:"optional" json:"tokenEndpointUrl" yaml:"tokenEndpointUrl"`
	// Username of the account to connect to the Kafka cluster.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

