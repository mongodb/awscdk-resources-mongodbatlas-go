package awscdkresourcesmongodbatlas


// Tokenizer that you want to use to create tokens.
//
// Tokens determine how Atlas Search splits up text into discrete chunks for indexing.
type ApiAtlasFtsAnalyzersTokenizer struct {
	// Index of the character group within the matching expression to extract into tokens.
	//
	// Use `0` to extract all character groups.
	Group *float64 `field:"optional" json:"group" yaml:"group"`
	// Characters to include in the longest token that Atlas Search creates.
	MaxGram *float64 `field:"optional" json:"maxGram" yaml:"maxGram"`
	// Maximum number of characters in a single token.
	//
	// Tokens greater than this length are split at this length into multiple tokens.
	MaxTokenLength *float64 `field:"optional" json:"maxTokenLength" yaml:"maxTokenLength"`
	// Characters to include in the shortest token that Atlas Search creates.
	MinGram *float64 `field:"optional" json:"minGram" yaml:"minGram"`
	// Regular expression to match against.
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// Human-readable label that identifies this tokenizer type.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

