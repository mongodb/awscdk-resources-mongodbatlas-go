package awscdkresourcesmongodbatlas


type TypeSet struct {
	// Human-readable label that identifies this type set.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Stringify json representation of types array.
	//
	// Each type defines a field type for the search index.
	Types *string `field:"optional" json:"types" yaml:"types"`
}

