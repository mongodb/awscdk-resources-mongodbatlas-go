package awscdkresourcesmongodbatlas


// Configurable timeouts for stream processor operations.
type Timeouts struct {
	// Timeout for create operation in Go duration format (e.g., '5m', '10s'). Default is 20 minutes.
	// Default: 20 minutes.
	//
	Create *string `field:"optional" json:"create" yaml:"create"`
}

