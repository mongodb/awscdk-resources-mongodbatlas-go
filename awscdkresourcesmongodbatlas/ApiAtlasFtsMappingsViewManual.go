package awscdkresourcesmongodbatlas


type ApiAtlasFtsMappingsViewManual struct {
	// Flag that indicates whether the index uses dynamic or static mappings.
	//
	// If DynamicConfig is specified, this field is ignored (DynamicConfig takes precedence). Required for search indexes if **mappings.fields** is omitted and **mappings.dynamicConfig** is not specified.
	Dynamic *bool `field:"optional" json:"dynamic" yaml:"dynamic"`
	// Stringify json representation of dynamic mapping configuration object.
	//
	// This allows for more complex dynamic mapping configurations beyond a simple boolean. If both Dynamic and DynamicConfig are specified, DynamicConfig takes precedence.
	DynamicConfig *string `field:"optional" json:"dynamicConfig" yaml:"dynamicConfig"`
	// One or more field specifications for the Atlas Search index.
	//
	// Stringify json representation of field with types and properties. Required for search indexes if **mappings.dynamic** and **mappings.dynamicConfig** are omitted or if **mappings.dynamic** is set to **false**.
	Fields *string `field:"optional" json:"fields" yaml:"fields"`
}

