package awscdkresourcesmongodbatlas


// Settings needed to configure the MongoDB Connector for Business Intelligence for this cluster.
type CfnClusterPropsBiConnector struct {
	// Flag that indicates whether MongoDB Connector for Business Intelligence is enabled on the specified cluster.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Data source node designated for the MongoDB Connector for Business Intelligence on MongoDB Cloud.
	//
	// The MongoDB Connector for Business Intelligence on MongoDB Cloud reads data from the primary, secondary, or analytics node based on your read preferences. Defaults to ANALYTICS node, or SECONDARY if there are no ANALYTICS nodes.
	ReadPreference *string `field:"optional" json:"readPreference" yaml:"readPreference"`
}

