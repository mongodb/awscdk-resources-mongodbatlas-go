package awscdkresourcesmongodbatlas


type SynonymSource struct {
	// Human-readable label that identifies the MongoDB collection that stores words and their applicable synonyms.
	Collection *string `field:"required" json:"collection" yaml:"collection"`
}

