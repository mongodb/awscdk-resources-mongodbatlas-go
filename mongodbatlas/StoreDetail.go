// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Configuration information for each data store and its mapping to MongoDB Cloud databases.
type StoreDetail struct {
	// Human-readable label that identifies the Federated Database to update.
	AdditionalStorageClasses *[]*string `field:"optional" json:"additionalStorageClasses" yaml:"additionalStorageClasses"`
	// Human-readable label that identifies the Federated Database to update.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Human-readable label that identifies the Federated Database to update.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Human-readable label that identifies the Federated Database to update.
	IncludeTags *bool `field:"optional" json:"includeTags" yaml:"includeTags"`
	// Human-readable label that identifies the data store.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Human-readable label that identifies the Federated Database to update.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Human-readable label that identifies the Federated Database to update.
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// Human-readable label that identifies the Federated Database to update.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

