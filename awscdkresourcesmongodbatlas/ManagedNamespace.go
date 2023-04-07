// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type ManagedNamespace struct {
	// Human-readable label of the collection to manage for this Global Cluster.
	Collection *string `field:"optional" json:"collection" yaml:"collection"`
	// Database parameter used to divide the *collection* into shards.
	//
	// Global clusters require a compound shard key. This compound shard key combines the location parameter and the user-selected custom key.
	CustomShardKey *string `field:"optional" json:"customShardKey" yaml:"customShardKey"`
	// Human-readable label of the database to manage for this Global Cluster.
	Db *string `field:"optional" json:"db" yaml:"db"`
	// Flag that indicates whether someone hashed the custom shard key for the specified collection.
	//
	// If you set this value to `false`, MongoDB Cloud uses ranged sharding.
	IsCustomShardKeyHashed *bool `field:"optional" json:"isCustomShardKeyHashed" yaml:"isCustomShardKeyHashed"`
	// Flag that indicates whether someone [hashed](https://www.mongodb.com/docs/manual/reference/method/sh.shardCollection/#hashed-shard-keys) the custom shard key. If this parameter returns `false`, this cluster uses [ranged sharding](https://www.mongodb.com/docs/manual/core/ranged-sharding/).
	IsShardKeyUnique *bool `field:"optional" json:"isShardKeyUnique" yaml:"isShardKeyUnique"`
}

