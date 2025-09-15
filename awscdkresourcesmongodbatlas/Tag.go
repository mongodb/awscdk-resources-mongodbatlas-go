package awscdkresourcesmongodbatlas


// List that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the cluster.
type Tag struct {
	// Constant that defines the set of the tag.
	//
	// For example, environment in the environment : production tag.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Variable that belongs to the set of the tag.
	//
	// For example, production in the environment : production tag.
	Value *string `field:"required" json:"value" yaml:"value"`
}

