// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Advanced configuration details to add for one cluster in the specified project.
type ProcessArgs struct {
	// Default level of acknowledgment requested from MongoDB for read operations set for this cluster.
	DefaultReadConcern *string `field:"optional" json:"defaultReadConcern" yaml:"defaultReadConcern"`
	// Default level of acknowledgment requested from MongoDB for write operations set for this cluster.
	DefaultWriteConcern *string `field:"optional" json:"defaultWriteConcern" yaml:"defaultWriteConcern"`
	// Flag that indicates whether you can insert or update documents where all indexed entries don't exceed 1024 bytes.
	//
	// If you set this to false, mongod writes documents that exceed this limit but doesn't index them.
	FailIndexKeyTooLong *bool `field:"optional" json:"failIndexKeyTooLong" yaml:"failIndexKeyTooLong"`
	// Flag that indicates whether the cluster allows execution of operations that perform server-side executions of JavaScript.
	JavascriptEnabled *bool `field:"optional" json:"javascriptEnabled" yaml:"javascriptEnabled"`
	// Minimum Transport Layer Security (TLS) version that the cluster accepts for incoming connections.
	//
	// Clusters using TLS 1.0 or 1.1 should consider setting TLS 1.2 as the minimum TLS protocol version.
	MinimumEnabledTlsProtocol *string `field:"optional" json:"minimumEnabledTlsProtocol" yaml:"minimumEnabledTlsProtocol"`
	// Flag that indicates whether the cluster disables executing any query that requires a collection scan to return results.
	NoTableScan *bool `field:"optional" json:"noTableScan" yaml:"noTableScan"`
	// Storage limit of cluster's oplog expressed in megabytes.
	//
	// A value of null indicates that the cluster uses the default oplog size that MongoDB Cloud calculates.
	OplogSizeMb *float64 `field:"optional" json:"oplogSizeMb" yaml:"oplogSizeMb"`
	// Number of documents per database to sample when gathering schema information.
	SampleRefreshIntervalBiConnector *float64 `field:"optional" json:"sampleRefreshIntervalBiConnector" yaml:"sampleRefreshIntervalBiConnector"`
	// Interval in seconds at which the mongosqld process re-samples data to create its relational schema.
	SampleSizeBiConnector *float64 `field:"optional" json:"sampleSizeBiConnector" yaml:"sampleSizeBiConnector"`
}

