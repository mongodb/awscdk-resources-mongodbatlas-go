// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type ApiAtlasFtsMappingsViewManual struct {
	// Flag that indicates whether the index uses dynamic or static mappings.
	//
	// Required if **mappings.fields** is omitted.
	Dynamic *bool `field:"optional" json:"dynamic" yaml:"dynamic"`
	// One or more field specifications for the Atlas Search index.
	//
	// The element of the array must have the format fieldName:fieldType. Required if **mappings.dynamic** is omitted or set to **false**.
	Fields *[]*string `field:"optional" json:"fields" yaml:"fields"`
}

