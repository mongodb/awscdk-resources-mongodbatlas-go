// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type AtlasBasicPrivateEndpointProps struct {
	AtlasBasicProps *AtlasBasicProps `field:"required" json:"atlasBasicProps" yaml:"atlasBasicProps"`
	PrivateEndpointProps *PrivateEndpointProps `field:"required" json:"privateEndpointProps" yaml:"privateEndpointProps"`
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	Region *string `field:"optional" json:"region" yaml:"region"`
}

