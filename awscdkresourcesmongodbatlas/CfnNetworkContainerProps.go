// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


// Returns, adds, edits, and removes network peering containers.
type CfnNetworkContainerProps struct {
	// IP addresses expressed in Classless Inter-Domain Routing (CIDR) notation that MongoDB Cloud uses for the network peering containers in your project.
	//
	// MongoDB Cloud assigns all of the project's clusters deployed to this cloud provider an IP address from this range. MongoDB Cloud locks this value if an M10 or greater cluster or a network peering connection exists in this project.
	// These CIDR blocks must fall within the ranges reserved per RFC 1918. AWS further limits the block to between the /24 and /21 ranges.
	// To modify the CIDR block, the target project cannot have:
	// - Any M10 or greater clusters
	// - Any other VPC peering connections
	// You can also create a new project and create a network peering connection to set the desired MongoDB Cloud network peering container CIDR block for that project. MongoDB Cloud limits the number of MongoDB nodes per network peering connection based on the CIDR block and the region selected for the project.
	// Example: A project in an Amazon Web Services (AWS) region supporting three availability zones and an MongoDB CIDR network peering container block of limit of /24 equals 27 three-node replica sets.
	AtlasCidrBlock *string `field:"required" json:"atlasCidrBlock" yaml:"atlasCidrBlock"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Geographic area that Amazon Web Services (AWS) defines to which MongoDB Cloud deployed this network peering container.
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Boolean flag that indicates whether MongoDB Cloud clusters exist in the specified network peering container.
	Provisioned *bool `field:"optional" json:"provisioned" yaml:"provisioned"`
	// Unique string that identifies the MongoDB Cloud VPC on AWS.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

