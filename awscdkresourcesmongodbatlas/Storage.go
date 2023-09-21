package awscdkresourcesmongodbatlas


// Configuration information for each data store and its mapping to MongoDB Cloud databases.
type Storage struct {
	// Array that contains the queryable databases and collections for this data lake.
	Databases *[]*Database `field:"optional" json:"databases" yaml:"databases"`
	// Array that contains the data stores for the data lake.
	Stores *[]*Store `field:"optional" json:"stores" yaml:"stores"`
}

