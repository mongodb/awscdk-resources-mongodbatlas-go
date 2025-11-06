package awscdkresourcesmongodbatlas


// Information about the cloud provider region to which the Atlas Data Federation routes client connections.
//
// MongoDB Cloud supports AWS only.
type DataProcessRegion struct {
	// Name of the region to which the Atlas Data Federation routes client connections.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Name of the cloud service that hosts the Atlas Data Federation data stores.
	CloudProvider *string `field:"optional" json:"cloudProvider" yaml:"cloudProvider"`
}

