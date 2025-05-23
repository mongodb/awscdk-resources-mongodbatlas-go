package awscdkresourcesmongodbatlas


type AtlasBasicPrivateEndpointProps struct {
	AtlasBasicProps *AtlasBasicProps `field:"required" json:"atlasBasicProps" yaml:"atlasBasicProps"`
	PrivateEndpointProps *PrivateEndpointProps `field:"required" json:"privateEndpointProps" yaml:"privateEndpointProps"`
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Default: us-east-1.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

