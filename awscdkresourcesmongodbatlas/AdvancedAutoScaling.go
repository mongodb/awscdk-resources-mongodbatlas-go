package awscdkresourcesmongodbatlas


// AWS Automatic Cluster Scaling.
type AdvancedAutoScaling struct {
	Compute *Compute `field:"optional" json:"compute" yaml:"compute"`
	DiskGb *DiskGb `field:"optional" json:"diskGb" yaml:"diskGb"`
}

