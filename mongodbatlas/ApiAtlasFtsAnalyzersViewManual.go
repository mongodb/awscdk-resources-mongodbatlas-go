// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type ApiAtlasFtsAnalyzersViewManual struct {
	// Filters that examine text one character at a time and perform filtering operations.
	CharFilters *[]interface{} `field:"optional" json:"charFilters" yaml:"charFilters"`
	// Human-readable name that identifies the custom analyzer.
	//
	// Names must be unique within an index, and must not start with any of the following strings:
	// - `lucene.`
	// - `builtin.`
	// - `mongodb.`
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Filter that performs operations such as:.
	//
	// - Stemming, which reduces related words, such as "talking", "talked", and "talks" to their root word "talk".
	//
	// - Redaction, the removal of sensitive information from public documents.
	TokenFilters *[]interface{} `field:"optional" json:"tokenFilters" yaml:"tokenFilters"`
	// Tokenizer that you want to use to create tokens.
	//
	// Tokens determine how Atlas Search splits up text into discrete chunks for indexing.
	Tokenizer interface{} `field:"optional" json:"tokenizer" yaml:"tokenizer"`
}

