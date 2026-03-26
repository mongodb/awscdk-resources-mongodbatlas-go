package awscdkresourcesmongodbatlas


// Type of the connection.
//
// Can be Cluster, Kafka, Sample, AWSLambda, or Https.
type CfnStreamConnectionPropsType string

const (
	// Kafka.
	CfnStreamConnectionPropsType_KAFKA CfnStreamConnectionPropsType = "KAFKA"
	// Cluster.
	CfnStreamConnectionPropsType_CLUSTER CfnStreamConnectionPropsType = "CLUSTER"
	// Sample.
	CfnStreamConnectionPropsType_SAMPLE CfnStreamConnectionPropsType = "SAMPLE"
	// AWSLambda.
	CfnStreamConnectionPropsType_AWS_LAMBDA CfnStreamConnectionPropsType = "AWS_LAMBDA"
	// Https.
	CfnStreamConnectionPropsType_HTTPS CfnStreamConnectionPropsType = "HTTPS"
)

