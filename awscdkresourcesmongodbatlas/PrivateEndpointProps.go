package awscdkresourcesmongodbatlas


type PrivateEndpointProps struct {
	AwsSubnetId *string `field:"required" json:"awsSubnetId" yaml:"awsSubnetId"`
	AwsVpcId *string `field:"required" json:"awsVpcId" yaml:"awsVpcId"`
}

