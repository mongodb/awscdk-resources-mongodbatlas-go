// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type ProjectTeam struct {
	RoleNames *[]*string `field:"optional" json:"roleNames" yaml:"roleNames"`
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
}

