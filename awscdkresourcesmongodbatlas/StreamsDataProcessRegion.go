package awscdkresourcesmongodbatlas


// Information about the cloud provider region in which MongoDB Cloud processes the stream.
type StreamsDataProcessRegion struct {
	// Label that identifies the cloud service provider where MongoDB Cloud performs stream processing.
	//
	// Currently, this parameter supports AWS only.
	CloudProvider StreamsDataProcessRegionCloudProvider `field:"required" json:"cloudProvider" yaml:"cloudProvider"`
	Region *string `field:"required" json:"region" yaml:"region"`
}

