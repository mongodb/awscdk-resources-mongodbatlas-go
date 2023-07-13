package awscdkresourcesmongodbatlas


type AtlasEncryptionAtRestExpressProps struct {
	EncryptionAtRest *EncryptionAtRestProps `field:"required" json:"encryptionAtRest" yaml:"encryptionAtRest"`
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	AccessList *IpAccessListProps `field:"optional" json:"accessList" yaml:"accessList"`
	Cluster *ClusterProps `field:"optional" json:"cluster" yaml:"cluster"`
	DatabaseUser *DatabaseUserProps `field:"optional" json:"databaseUser" yaml:"databaseUser"`
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

