// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type ApiAtlasDiskBackupCopySettingView struct {
	// A label that identifies the cloud provider that stores the snapshot copy.
	CloudProvider *string `field:"optional" json:"cloudProvider" yaml:"cloudProvider"`
	// List that describes which types of snapshots to copy.
	Frequencies *[]*string `field:"optional" json:"frequencies" yaml:"frequencies"`
	// Target region to copy snapshots belonging to replicationSpecId to.
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster.
	ReplicationSpecId *string `field:"optional" json:"replicationSpecId" yaml:"replicationSpecId"`
	// Flag that indicates whether to copy the oplogs to the target region.
	ShouldCopyOplogs *bool `field:"optional" json:"shouldCopyOplogs" yaml:"shouldCopyOplogs"`
}

