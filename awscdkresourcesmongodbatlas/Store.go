package awscdkresourcesmongodbatlas


// Array that contains the data stores for the data lake.
type Store struct {
	// Human-readable label of the MongoDB Cloud cluster on which the store is based.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Human-readable label that identifies the data store.
	//
	// The databases.[n].collections.[n].dataSources.[n].storeName field references this values as part of the mapping configuration. To use MongoDB Cloud as a data store, the data lake requires a serverless instance or an M10 or higher cluster.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Unique 24-hexadecimal digit string that identifies the project.Regex ^([a-f0-9]{24})$ .
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Allowed values atlas, http, online_archive, s3 and DataLakeAzureBlobStore.
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// MongoDB Cloud cluster read preference, which describes how to route read requests to the cluster.
	ReadPreference *ReadPreference `field:"optional" json:"readPreference" yaml:"readPreference"`
}

