package awscdkresourcesmongodbatlas


// Returns, adds, edits, and removes network peering containers and peering connections.
type CfnNetworkPeeringProps struct {
	// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that contains the specified network peering connection.
	ContainerId *string `field:"required" json:"containerId" yaml:"containerId"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Unique string that identifies the VPC on Amazon Web Services (AWS) that you want to peer with the MongoDB Cloud VPC.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Amazon Web Services (AWS) region where the Virtual Peering Connection (VPC) that you peered with the MongoDB Cloud VPC resides.
	//
	// The resource returns null if your VPC and the MongoDB Cloud VPC reside in the same region.
	AccepterRegionName *string `field:"optional" json:"accepterRegionName" yaml:"accepterRegionName"`
	// Unique twelve-digit string that identifies the Amazon Web Services (AWS) account that owns the VPC that you peered with the MongoDB Cloud VPC.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Internet Protocol (IP) addresses expressed in Classless Inter-Domain Routing (CIDR) notation of the VPC's subnet that you want to peer with the MongoDB Cloud VPC.
	RouteTableCidrBlock *string `field:"optional" json:"routeTableCidrBlock" yaml:"routeTableCidrBlock"`
}

