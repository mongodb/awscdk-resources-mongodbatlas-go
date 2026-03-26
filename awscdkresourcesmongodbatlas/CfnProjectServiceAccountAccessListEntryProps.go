package awscdkresourcesmongodbatlas


// Manages IP access list entries for MongoDB Atlas Project Service Accounts.
type CfnProjectServiceAccountAccessListEntryProps struct {
	// The Client ID of the Service Account.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Unique 24-hexadecimal digit string that identifies the project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Range of IP addresses in CIDR notation to be added to the access list.
	//
	// You can set a value for this parameter or IPAddress, but not both.
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Single IP address to be added to the access list.
	//
	// You can set a value for this parameter or CIDRBlock, but not both.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

