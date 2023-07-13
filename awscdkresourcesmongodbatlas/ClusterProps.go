package awscdkresourcesmongodbatlas


type ClusterProps struct {
	AdvancedSettings *ProcessArgs `field:"optional" json:"advancedSettings" yaml:"advancedSettings"`
	BackupEnabled *bool `field:"optional" json:"backupEnabled" yaml:"backupEnabled"`
	BiConnector *CfnClusterPropsBiConnector `field:"optional" json:"biConnector" yaml:"biConnector"`
	ClusterType *string `field:"optional" json:"clusterType" yaml:"clusterType"`
	ConnectionStrings *ConnectionStrings `field:"optional" json:"connectionStrings" yaml:"connectionStrings"`
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	EncryptionAtRestProvider CfnClusterPropsEncryptionAtRestProvider `field:"optional" json:"encryptionAtRestProvider" yaml:"encryptionAtRestProvider"`
	Labels *[]*CfnClusterPropsLabels `field:"optional" json:"labels" yaml:"labels"`
	MongoDbMajorVersion *string `field:"optional" json:"mongoDbMajorVersion" yaml:"mongoDbMajorVersion"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Paused *bool `field:"optional" json:"paused" yaml:"paused"`
	PitEnabled *bool `field:"optional" json:"pitEnabled" yaml:"pitEnabled"`
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	ReplicationSpecs *[]*AdvancedReplicationSpec `field:"optional" json:"replicationSpecs" yaml:"replicationSpecs"`
	RootCertType *string `field:"optional" json:"rootCertType" yaml:"rootCertType"`
	TerminationProtectionEnabled *bool `field:"optional" json:"terminationProtectionEnabled" yaml:"terminationProtectionEnabled"`
	VersionReleaseSystem *string `field:"optional" json:"versionReleaseSystem" yaml:"versionReleaseSystem"`
}

