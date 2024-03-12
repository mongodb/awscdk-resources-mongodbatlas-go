package awscdkresourcesmongodbatlas


// Configuration options for an Atlas Stream Processing Instance.
type StreamConfig struct {
	// Selected tier for the Stream Instance.
	//
	// Configures Memory / VCPU allowances.
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

