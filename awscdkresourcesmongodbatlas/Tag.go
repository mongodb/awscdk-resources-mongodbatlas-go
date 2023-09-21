package awscdkresourcesmongodbatlas


// Advanced configuration details to add for one cluster in the specified project.
type Tag struct {
	// Constant that defines the set of the tag.
	//
	// For example, environment in the environment : production tag.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Variable that belongs to the set of the tag.
	//
	// For example, production in the environment : production tag.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

