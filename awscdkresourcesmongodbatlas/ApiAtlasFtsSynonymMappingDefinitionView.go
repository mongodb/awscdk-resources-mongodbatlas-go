package awscdkresourcesmongodbatlas


type ApiAtlasFtsSynonymMappingDefinitionView struct {
	// Specific pre-defined method chosen to apply to the synonyms to be searched.
	Analyzer *string `field:"optional" json:"analyzer" yaml:"analyzer"`
	// Human-readable label that identifies the synonym definition.
	//
	// Each **synonym.name** must be unique within the same index definition.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Data set that stores the mapping one or more words map to one or more synonyms of those words.
	Source *SynonymSource `field:"optional" json:"source" yaml:"source"`
}

