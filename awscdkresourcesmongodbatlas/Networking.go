package awscdkresourcesmongodbatlas


// Networking configuration for connections.
type Networking struct {
	// Network access configuration.
	Access *NetworkingAccess `field:"required" json:"access" yaml:"access"`
}

