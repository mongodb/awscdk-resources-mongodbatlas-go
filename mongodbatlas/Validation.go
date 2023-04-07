// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type Validation struct {
	Status *string `field:"optional" json:"status" yaml:"status"`
	ValidationType *string `field:"optional" json:"validationType" yaml:"validationType"`
}

