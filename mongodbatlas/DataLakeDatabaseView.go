// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type DataLakeDatabaseView struct {
	// Array of collections and data sources that map to a ``stores`` data store.
	Collections *[]*DataLakeDatabaseCollectionView `field:"optional" json:"collections" yaml:"collections"`
	// Maximum number of wildcard collections in the database.
	//
	// This only applies to S3 data sources.
	MaxWildcardCollections *float64 `field:"optional" json:"maxWildcardCollections" yaml:"maxWildcardCollections"`
	// Human-readable label that identifies the database to which the data lake maps data.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Array of aggregation pipelines that apply to the collection.
	//
	// This only applies to S3 data sources.
	Views *[]*DataLakeViewView `field:"optional" json:"views" yaml:"views"`
}

