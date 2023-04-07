// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type ApiAtlasDiskBackupShardedClusterSnapshotMemberView struct {
	// Human-readable label that identifies the cloud provider that stores this snapshot.
	//
	// The resource returns this parameter when `"type": "replicaSet".`
	CloudProvider ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider `field:"optional" json:"cloudProvider" yaml:"cloudProvider"`
	// Unique 24-hexadecimal digit string that identifies the snapshot.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Human-readable label that identifies the shard or config host from which MongoDB Cloud took this snapshot.
	ReplicaSetName *string `field:"optional" json:"replicaSetName" yaml:"replicaSetName"`
}

