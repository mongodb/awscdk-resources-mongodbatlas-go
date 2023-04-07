// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type AtlasBasicProps struct {
	ClusterProps *ClusterProps `field:"required" json:"clusterProps" yaml:"clusterProps"`
	ProjectProps *ProjectProps `field:"required" json:"projectProps" yaml:"projectProps"`
	DbUserProps *DatabaseUserProps `field:"optional" json:"dbUserProps" yaml:"dbUserProps"`
	IpAccessListProps *IpAccessListProps `field:"optional" json:"ipAccessListProps" yaml:"ipAccessListProps"`
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

