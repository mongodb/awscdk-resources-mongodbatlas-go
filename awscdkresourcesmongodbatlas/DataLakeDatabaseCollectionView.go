package awscdkresourcesmongodbatlas


type DataLakeDatabaseCollectionView struct {
	// Array that contains the data stores that map to a collection for this data lake.
	DataSources *[]*DataLakeDatabaseDataSourceView `field:"optional" json:"dataSources" yaml:"dataSources"`
	// Human-readable label that identifies the collection to which MongoDB Cloud maps the data in the data stores.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

