// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type SynonymSource struct {
	// Human-readable label that identifies the MongoDB collection that stores words and their applicable synonyms.
	Collection *string `field:"optional" json:"collection" yaml:"collection"`
}

