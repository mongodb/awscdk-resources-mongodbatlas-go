package awscdkresourcesmongodbatlas


// Information about the cloud provider region to which the data lake routes client connections.
//
// MongoDB Cloud supports AWS only.
type DataProcessRegion struct {
	// Name of the cloud service that hosts the data lake's data stores.
	CloudProvider *string `field:"optional" json:"cloudProvider" yaml:"cloudProvider"`
	// Name of the region to which the data lake routes client connections.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

