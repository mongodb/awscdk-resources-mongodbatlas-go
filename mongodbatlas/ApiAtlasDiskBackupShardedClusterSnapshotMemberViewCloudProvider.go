// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Human-readable label that identifies the cloud provider that stores this snapshot.
//
// The resource returns this parameter when `"type": "replicaSet".`
type ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider string

const (
	// AWS.
	ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider_AWS ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider = "AWS"
	// AZURE.
	ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider_AZURE ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider = "AZURE"
	// GCP.
	ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider_GCP ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider = "GCP"
)

