package awscdkresourcesmongodbatlas


// Configuration options for an Atlas Stream Processing Workspace.
type StreamConfig struct {
	// Selected tier for the Stream Workspace.
	//
	// Configures Memory / VCPU allowances. Valid values: SP2, SP5, SP10, SP30, SP50.
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// Max tier size for the Stream Workspace.
	//
	// Configures Memory / VCPU allowances. Valid values: SP2, SP5, SP10, SP30, SP50.
	MaxTierSize *string `field:"optional" json:"maxTierSize" yaml:"maxTierSize"`
}

