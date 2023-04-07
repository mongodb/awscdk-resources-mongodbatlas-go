// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type ProjectApiKey struct {
	Key *string `field:"optional" json:"key" yaml:"key"`
	RoleNames *[]*string `field:"optional" json:"roleNames" yaml:"roleNames"`
}

