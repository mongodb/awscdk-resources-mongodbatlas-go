package awscdkresourcesmongodbatlas


type Action struct {
	// Human-readable label that identifies the privilege action.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// List of resources on which you grant the action.
	Resources *[]*Resource `field:"optional" json:"resources" yaml:"resources"`
}

