package awscdkresourcesmongodbatlas


// Ordered fields used to physically organize data in the destination.
type Transformations struct {
	// Key in the document.
	Field *string `field:"optional" json:"field" yaml:"field"`
	// Type of transformation applied during the export of the namespace in a Data Lake Pipeline.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

