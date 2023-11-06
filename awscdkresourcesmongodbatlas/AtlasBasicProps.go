package awscdkresourcesmongodbatlas


type AtlasBasicProps struct {
	ClusterProps *ClusterProps `field:"required" json:"clusterProps" yaml:"clusterProps"`
	IpAccessListProps *IpAccessListProps `field:"required" json:"ipAccessListProps" yaml:"ipAccessListProps"`
	ProjectProps *ProjectProps `field:"required" json:"projectProps" yaml:"projectProps"`
	DbUserProps *DatabaseUserProps `field:"optional" json:"dbUserProps" yaml:"dbUserProps"`
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

