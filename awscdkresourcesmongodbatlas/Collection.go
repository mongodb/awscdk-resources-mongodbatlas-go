package awscdkresourcesmongodbatlas


// Array of collections and data sources that map to a stores data store.
type Collection struct {
	// Array that contains the data stores that map to a collection for this data lake.
	DataSources *[]*DataSource `field:"optional" json:"dataSources" yaml:"dataSources"`
	Name *string `field:"optional" json:"name" yaml:"name"`
}

