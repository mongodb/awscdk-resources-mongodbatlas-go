package awscdkresourcesmongodbatlas


type AtlasBasicPrivateEndpointProps struct {
	AtlasBasicProps *AtlasBasicProps `field:"required" json:"atlasBasicProps" yaml:"atlasBasicProps"`
	PrivateEndpointProps *PrivateEndpointProps `field:"required" json:"privateEndpointProps" yaml:"privateEndpointProps"`
	// Default: us-east-1.
	//
	Region *string `field:"required" json:"region" yaml:"region"`
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

