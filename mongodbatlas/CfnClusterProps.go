// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// The cluster resource provides access to your cluster configurations.
//
// The resource lets you create, edit and delete clusters. The resource requires your Project ID.
type CfnClusterProps struct {
	// Human-readable label that identifies the advanced cluster.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Unique identifier of the project the cluster belongs to.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	AdvancedSettings *ProcessArgs `field:"optional" json:"advancedSettings" yaml:"advancedSettings"`
	// Flag that indicates whether the cluster can perform backups.
	//
	// If set to true, the cluster can perform backups. You must set this value to true for NVMe clusters. Backup uses Cloud Backups for dedicated clusters and Shared Cluster Backups for tenant clusters. If set to false, the cluster doesn't use backups.
	BackupEnabled *bool `field:"optional" json:"backupEnabled" yaml:"backupEnabled"`
	// Settings needed to configure the MongoDB Connector for Business Intelligence for this cluster.
	BiConnector *CfnClusterPropsBiConnector `field:"optional" json:"biConnector" yaml:"biConnector"`
	// Configuration of nodes that comprise the cluster.
	ClusterType *string `field:"optional" json:"clusterType" yaml:"clusterType"`
	// Set of connection strings that your applications use to connect to this cluster.
	//
	// Use the parameters in this object to connect your applications to this cluster. See the MongoDB [Connection String URI Format](https://docs.mongodb.com/manual/reference/connection-string/) reference for further details.
	ConnectionStrings *ConnectionStrings `field:"optional" json:"connectionStrings" yaml:"connectionStrings"`
	// Storage capacity that the host's root volume possesses expressed in gigabytes.
	//
	// Increase this number to add capacity. MongoDB Cloud requires this parameter if you set replicationSpecs. If you specify a disk size below the minimum (10 GB), this parameter defaults to the minimum disk size value. Storage charge calculations depend on whether you choose the default value or a custom value. The maximum value for disk storage cannot exceed 50 times the maximum RAM for the selected cluster. If you require more storage space, consider upgrading your cluster to a higher tier.
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Cloud service provider that manages your customer keys to provide an additional layer of encryption at rest for the cluster.
	//
	// To enable customer key management for encryption at rest, the cluster replicationSpecs[n].regionConfigs[m].{type}Specs.instanceSize setting must be M10 or higher and "backupEnabled" : false or omitted entirely.
	EncryptionAtRestProvider CfnClusterPropsEncryptionAtRestProvider `field:"optional" json:"encryptionAtRestProvider" yaml:"encryptionAtRestProvider"`
	// Collection of key-value pairs between 1 to 255 characters in length that tag and categorize the cluster.
	//
	// The MongoDB Cloud console doesn't display your labels.
	Labels *[]*CfnClusterPropsLabels `field:"optional" json:"labels" yaml:"labels"`
	// Major MongoDB version of the cluster.
	//
	// MongoDB Cloud deploys the cluster with the latest stable release of the specified version.
	MongoDbMajorVersion *string `field:"optional" json:"mongoDbMajorVersion" yaml:"mongoDbMajorVersion"`
	// Flag that indicates whether the cluster is paused or not.
	Paused *bool `field:"optional" json:"paused" yaml:"paused"`
	// Flag that indicates whether the cluster uses continuous cloud backups.
	PitEnabled *bool `field:"optional" json:"pitEnabled" yaml:"pitEnabled"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// List of settings that configure your cluster regions.
	//
	// For Global Clusters, each object in the array represents a zone where your clusters nodes deploy. For non-Global replica sets and sharded clusters, this array has one object representing where your clusters nodes deploy.
	ReplicationSpecs *[]*AdvancedReplicationSpec `field:"optional" json:"replicationSpecs" yaml:"replicationSpecs"`
	// Root Certificate Authority that MongoDB Cloud cluster uses.
	//
	// MongoDB Cloud supports Internet Security Research Group.
	RootCertType *string `field:"optional" json:"rootCertType" yaml:"rootCertType"`
	// Flag that indicates whether termination protection is enabled on the cluster.
	//
	// If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster.
	TerminationProtectionEnabled *bool `field:"optional" json:"terminationProtectionEnabled" yaml:"terminationProtectionEnabled"`
	// Method by which the cluster maintains the MongoDB versions.
	//
	// If value is CONTINUOUS, you must not specify mongoDBMajorVersion.
	VersionReleaseSystem *string `field:"optional" json:"versionReleaseSystem" yaml:"versionReleaseSystem"`
}

