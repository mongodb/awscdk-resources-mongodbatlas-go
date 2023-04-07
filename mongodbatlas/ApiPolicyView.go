// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type ApiPolicyView struct {
	Id *string `field:"optional" json:"id" yaml:"id"`
	PolicyItems *[]*ApiPolicyItemView `field:"optional" json:"policyItems" yaml:"policyItems"`
}

