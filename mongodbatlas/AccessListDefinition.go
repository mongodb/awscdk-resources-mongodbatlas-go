// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type AccessListDefinition struct {
	// Unique string of the Amazon Web Services (AWS) security group that you want to add to the project's IP access list.
	//
	// Your IP access list entry can be one awsSecurityGroup, one cidrBlock, or one ipAddress. You must configure Virtual Private Connection (VPC) peering for your project before you can add an AWS security group to an IP access list. You cannot set AWS security groups as temporary access list entries. Don't set this parameter if you set cidrBlock or ipAddress.
	AwsSecurityGroup *string `field:"optional" json:"awsSecurityGroup" yaml:"awsSecurityGroup"`
	// Range of IP addresses in Classless Inter-Domain Routing (CIDR) notation that you want to add to the project's IP access list.
	//
	// Your IP access list entry can be one awsSecurityGroup, one cidrBlock, or one ipAddress. Don't set this parameter if you set awsSecurityGroup or ipAddress
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Remark that explains the purpose or scope of this IP access list entry.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Date and time after which MongoDB Cloud deletes the temporary access list entry.
	//
	// This parameter expresses its value in the ISO 8601 timestamp format in UTC and can include the time zone designation. The date must be later than the current date but no later than one week after you submit this request. The resource returns this parameter if you specified an expiration date when creating this IP access list entry.
	DeleteAfterDate *string `field:"optional" json:"deleteAfterDate" yaml:"deleteAfterDate"`
	// IP address that you want to add to the project's IP access list.
	//
	// Your IP access list entry can be one awsSecurityGroup, one cidrBlock, or one ipAddress. Don't set this parameter if you set awsSecurityGroup or cidrBlock.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

