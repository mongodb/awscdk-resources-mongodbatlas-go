package awscdkresourcesmongodbatlas


type View struct {
	// Human-readable label that identifies the view, which corresponds to an aggregation pipeline on a collection.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Aggregation pipeline stages to apply to the source collection.
	Pipeline *string `field:"optional" json:"pipeline" yaml:"pipeline"`
	// Human-readable label that identifies the source collection for the view.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

