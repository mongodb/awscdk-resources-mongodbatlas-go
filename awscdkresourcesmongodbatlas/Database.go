package awscdkresourcesmongodbatlas


type Database struct {
	// Array of collections and data sources that map to a stores data store.
	Collections *[]*Collection `field:"optional" json:"collections" yaml:"collections"`
	// Maximum number of wildcard collections in the database.
	//
	// This only applies to S3 data sources.
	MaxWildcardCollections *string `field:"optional" json:"maxWildcardCollections" yaml:"maxWildcardCollections"`
	// Human-readable label that identifies the database to which the Atlas Data Federation maps data.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Array of aggregation pipelines that apply to the collection.
	//
	// This only applies to S3 data sources.
	Views *[]*View `field:"optional" json:"views" yaml:"views"`
}

