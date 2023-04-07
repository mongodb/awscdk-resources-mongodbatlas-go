// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type ApiAtlasDiskBackupShardedClusterSnapshotView struct {
	// Date and time when MongoDB Cloud took the snapshot.
	//
	// This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Human-readable phrase or sentence that explains the purpose of the snapshot.
	//
	// The resource returns this parameter when `"status": "onDemand"`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Date and time when MongoDB Cloud deletes the snapshot.
	//
	// This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	ExpiresAt *string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// Human-readable label that identifies how often this snapshot triggers.
	FrequencyType ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType `field:"optional" json:"frequencyType" yaml:"frequencyType"`
	// Unique 24-hexadecimal digit string that identifies the snapshot.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both.
	//
	// RFC 5988 outlines these relationships.
	Links *[]*Link `field:"optional" json:"links" yaml:"links"`
	// Unique string that identifies the Amazon Web Services (AWS) Key Management Service (KMS) Customer Master Key (CMK) used to encrypt the snapshot.
	//
	// The resource returns this value when `"encryptionEnabled" : true`.
	MasterKeyUuid *string `field:"optional" json:"masterKeyUuid" yaml:"masterKeyUuid"`
	// List that includes the snapshots and the cloud provider that stores the snapshots.
	//
	// The resource returns this parameter when `"type" : "SHARDED_CLUSTER"`.
	Members *[]*ApiAtlasDiskBackupShardedClusterSnapshotMemberView `field:"optional" json:"members" yaml:"members"`
	// Version of the MongoDB host that this snapshot backs up.
	MongodVersion *string `field:"optional" json:"mongodVersion" yaml:"mongodVersion"`
	// List that contains unique identifiers for the policy items.
	PolicyItems *[]*string `field:"optional" json:"policyItems" yaml:"policyItems"`
	// List that contains the unique identifiers of the snapshots created for the shards and config host for a sharded cluster.
	//
	// The resource returns this parameter when `"type": "SHARDED_CLUSTER"`. These identifiers should match the ones specified in the **members[n].id** parameters. This allows you to map a snapshot to its shard or config host name.
	SnapshotIds *[]*string `field:"optional" json:"snapshotIds" yaml:"snapshotIds"`
	// Human-readable label that identifies when this snapshot triggers.
	SnapshotType ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType `field:"optional" json:"snapshotType" yaml:"snapshotType"`
	// Human-readable label that indicates the stage of the backup process for this snapshot.
	Status ApiAtlasDiskBackupShardedClusterSnapshotViewStatus `field:"optional" json:"status" yaml:"status"`
	// Number of bytes taken to store the backup snapshot.
	StorageSizeBytes *string `field:"optional" json:"storageSizeBytes" yaml:"storageSizeBytes"`
	// Human-readable label that categorizes the cluster as a replica set or sharded cluster.
	Type ApiAtlasDiskBackupShardedClusterSnapshotViewType `field:"optional" json:"type" yaml:"type"`
}

