// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


// AWS Automatic Cluster Scaling.
type AdvancedAutoScaling struct {
	Compute *Compute `field:"optional" json:"compute" yaml:"compute"`
	DiskGb *DiskGb `field:"optional" json:"diskGb" yaml:"diskGb"`
}

