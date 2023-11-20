package awscdkresourcesmongodbatlas


type ApiAtlasFtsMappingsViewManual struct {
	// Flag that indicates whether the index uses dynamic or static mappings.
	//
	// Required for search indexes if **mappings.fields** is omitted.
	Dynamic *bool `field:"optional" json:"dynamic" yaml:"dynamic"`
	// One or more field specifications for the Atlas Search index.
	//
	// Stringify json representation of field with types and properties. Required for search indexes if **mappings.dynamic** is omitted or set to **false**.
	Fields *string `field:"optional" json:"fields" yaml:"fields"`
}

