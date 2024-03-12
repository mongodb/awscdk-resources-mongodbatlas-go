package awscdkresourcesmongodbatlas


// Ingestion destination of a Data Lake Pipeline.
type Source struct {
	// Human-readable name that identifies the cluster.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Human-readable name that identifies the collection.
	CollectionName *string `field:"optional" json:"collectionName" yaml:"collectionName"`
	// Human-readable name that identifies the database.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Unique 24-hexadecimal character string that identifies the project.
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Type of ingestion source of this Data Lake Pipeline.
	Type SourceType `field:"optional" json:"type" yaml:"type"`
}

