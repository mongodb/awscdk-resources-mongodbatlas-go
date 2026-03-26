package awscdkresourcesmongodbatlas


// Returns, adds, edits, and removes Atlas Search indexes.
//
// Also returns and updates user-defined analyzers.
type CfnSearchIndexProps struct {
	// Name of the cluster that contains the database and collection with one or more Application Search indexes.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Human-readable label that identifies the collection that contains one or more Atlas Search indexes.
	CollectionName *string `field:"required" json:"collectionName" yaml:"collectionName"`
	// Human-readable label that identifies the database that contains the collection with one or more Atlas Search indexes.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Human-readable label that identifies this index.
	//
	// Within each namespace, names of all indexes in the namespace must be unique.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specific pre-defined method chosen to convert database field text into searchable words.
	//
	// This conversion reduces the text of fields into the smallest units of text. These units are called a **term** or **token**. This process, known as tokenization, involves a variety of changes made to the text in fields:
	//
	// - extracting words
	// - removing punctuation
	// - removing accents
	// - changing to lowercase
	// - removing common words
	// - reducing words to their root form (stemming)
	// - changing words to their base form (lemmatization)
	// MongoDB Cloud uses the selected process to build the Atlas Search index.
	Analyzer *string `field:"optional" json:"analyzer" yaml:"analyzer"`
	// List of user-defined methods to convert database field text into searchable words.
	Analyzers *[]*ApiAtlasFtsAnalyzersViewManual `field:"optional" json:"analyzers" yaml:"analyzers"`
	// Array of [Fields](https://www.mongodb.com/docs/atlas/atlas-search/field-types/knn-vector/#std-label-fts-data-types-knn-vector) to configure this vectorSearch index. Stringify json representation of field with types and properties. Required for vector indexes. It must contain at least one **vector** type field.
	Fields *string `field:"optional" json:"fields" yaml:"fields"`
	// Index specifications for the collection's fields.
	Mappings *ApiAtlasFtsMappingsViewManual `field:"optional" json:"mappings" yaml:"mappings"`
	// Number of partitions for the index.
	//
	// This is used to improve search performance for large datasets by distributing the index across multiple partitions.
	NumPartitions *float64 `field:"optional" json:"numPartitions" yaml:"numPartitions"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Method applied to identify words when searching this index.
	SearchAnalyzer *string `field:"optional" json:"searchAnalyzer" yaml:"searchAnalyzer"`
	// Flag that indicates whether to store the original document in the index.
	//
	// Can be a boolean ("true" or "false") or a stringified JSON object specifying which fields to include/exclude. When stored, this allows the index to return the original document for queries.
	StoredSource *string `field:"optional" json:"storedSource" yaml:"storedSource"`
	// Rule sets that map words to their synonyms in this index.
	Synonyms *[]*ApiAtlasFtsSynonymMappingDefinitionView `field:"optional" json:"synonyms" yaml:"synonyms"`
	// Type of index: **search** or **vectorSearch**.
	//
	// Default type is **search**.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Array of type sets that define alternate types for fields in the index.
	//
	// Each type set allows you to group related fields under a common name.
	TypeSets *[]*TypeSet `field:"optional" json:"typeSets" yaml:"typeSets"`
}

