package awscdkresourcesmongodbatlas


type AtlasServerlessBasicProps struct {
	ProjectProps *ProjectProps `field:"required" json:"projectProps" yaml:"projectProps"`
	ServerlessProps *CfnServerlessInstanceProps `field:"required" json:"serverlessProps" yaml:"serverlessProps"`
	DbUserProps *DatabaseUserProps `field:"optional" json:"dbUserProps" yaml:"dbUserProps"`
	IpAccessListProps *IpAccessListProps `field:"optional" json:"ipAccessListProps" yaml:"ipAccessListProps"`
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

