package awscdkresourcesmongodbatlas


// AWS Lambda connection configuration.
type Aws struct {
	// Amazon Resource Name (ARN) of the IAM role for AWS Lambda connection.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of an S3 bucket used to check authorization of the passed-in IAM role ARN.
	TestBucket *string `field:"optional" json:"testBucket" yaml:"testBucket"`
}

