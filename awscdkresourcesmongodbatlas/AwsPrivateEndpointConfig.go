package awscdkresourcesmongodbatlas


// AWS Private endpoint configuration properties.
type AwsPrivateEndpointConfig struct {
	// Status of the AWS PrivateEndpoint connection.
	AwsPrivateEndpointStatus *string `field:"optional" json:"awsPrivateEndpointStatus" yaml:"awsPrivateEndpointStatus"`
	// Unique identifiers of the interface endpoints in your VPC that you added to the AWS PrivateLink connection.
	InterfaceEndpointId *string `field:"optional" json:"interfaceEndpointId" yaml:"interfaceEndpointId"`
	// Aws Region.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// List of string representing the AWS VPC Subnet ID (like: subnet-xxxxxxxxxxxxxxxxx) (Used For Creating the AWS VPC Endpoint).
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// String Representing the AWS VPC ID (like: vpc-xxxxxxxxxxxxxxxx) (Used For Creating the AWS VPC Endpoint).
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

