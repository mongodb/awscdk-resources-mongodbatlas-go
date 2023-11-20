package awscdkresourcesmongodbatlas


type ApiAtlasFtsSynonymMappingDefinitionView struct {
	// Specific pre-defined method chosen to apply to the synonyms to be searched.
	Analyzer *string `field:"required" json:"analyzer" yaml:"analyzer"`
	// Human-readable label that identifies the synonym definition.
	//
	// Each **synonym.name** must be unique within the same index definition.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Data set that stores the mapping one or more words map to one or more synonyms of those words.
	Source *SynonymSource `field:"required" json:"source" yaml:"source"`
}

