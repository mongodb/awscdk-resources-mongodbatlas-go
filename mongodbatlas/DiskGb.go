// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Automatic cluster storage settings that apply to this cluster.
type DiskGb struct {
	// Flag that indicates whether this cluster enables disk auto-scaling.
	//
	// The maximum memory allowed for the selected cluster tier and the oplog size can limit storage auto-scaling.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

