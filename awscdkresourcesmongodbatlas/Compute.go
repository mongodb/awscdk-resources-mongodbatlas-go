package awscdkresourcesmongodbatlas


// Automatic Compute Scaling.
type Compute struct {
	// Flag that indicates whether someone enabled instance size auto-scaling.
	//
	// Set to true to enable instance size auto-scaling. If enabled, you must specify a value for replicationSpecs[n].regionConfigs[m].autoScaling.compute.maxInstanceSize.
	// Set to false to disable instance size automatic scaling.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Maximum instance size to which your cluster can automatically scale.
	//
	// MongoDB Cloud requires this parameter if "replicationSpecs[n].regionConfigs[m].autoScaling.compute.enabled" : true.
	MaxInstanceSize *string `field:"optional" json:"maxInstanceSize" yaml:"maxInstanceSize"`
	// Minimum instance size to which your cluster can automatically scale.
	//
	// MongoDB Cloud requires this parameter if "replicationSpecs[n].regionConfigs[m].autoScaling.compute.enabled" : true.
	MinInstanceSize *string `field:"optional" json:"minInstanceSize" yaml:"minInstanceSize"`
	// Flag that indicates whether the instance size may scale down.
	//
	// MongoDB Cloud requires this parameter if "replicationSpecs[n].regionConfigs[m].autoScaling.compute.enabled" : true. If you enable this option, specify a value for replicationSpecs[n].regionConfigs[m].autoScaling.compute.minInstanceSize.
	ScaleDownEnabled *bool `field:"optional" json:"scaleDownEnabled" yaml:"scaleDownEnabled"`
}

