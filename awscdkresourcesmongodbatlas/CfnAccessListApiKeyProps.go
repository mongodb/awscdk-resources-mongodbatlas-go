package awscdkresourcesmongodbatlas


// Creates the access list entries for the specified organization API key.
type CfnAccessListApiKeyProps struct {
	// Unique 24-hexadecimal digit string that identifies this organization API key for which you want to return access list entries.
	ApiUserId *string `field:"required" json:"apiUserId" yaml:"apiUserId"`
	// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// Range of network addresses that you want to add to the access list for the API key.
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Network address that you want to add to the access list for the API key.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Network address that issued the most recent request to the API.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Number of documents returned in this response.
	TotalCount *float64 `field:"optional" json:"totalCount" yaml:"totalCount"`
}

