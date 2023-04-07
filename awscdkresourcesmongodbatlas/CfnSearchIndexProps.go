// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
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
	// Index specifications for the collection's fields.
	Mappings *ApiAtlasFtsMappingsViewManual `field:"required" json:"mappings" yaml:"mappings"`
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
	// Human-readable label that identifies this index.
	//
	// Within each namespace, names of all indexes in the namespace must be unique.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Method applied to identify words when searching this index.
	SearchAnalyzer *string `field:"optional" json:"searchAnalyzer" yaml:"searchAnalyzer"`
	// Rule sets that map words to their synonyms in this index.
	Synonyms *[]*ApiAtlasFtsSynonymMappingDefinitionView `field:"optional" json:"synonyms" yaml:"synonyms"`
}

