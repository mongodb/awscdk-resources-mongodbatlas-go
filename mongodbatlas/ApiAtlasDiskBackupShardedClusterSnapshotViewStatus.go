// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Human-readable label that indicates the stage of the backup process for this snapshot.
type ApiAtlasDiskBackupShardedClusterSnapshotViewStatus string

const (
	// queued.
	ApiAtlasDiskBackupShardedClusterSnapshotViewStatus_QUEUED ApiAtlasDiskBackupShardedClusterSnapshotViewStatus = "QUEUED"
	// inProgress.
	ApiAtlasDiskBackupShardedClusterSnapshotViewStatus_IN_PROGRESS ApiAtlasDiskBackupShardedClusterSnapshotViewStatus = "IN_PROGRESS"
	// completed.
	ApiAtlasDiskBackupShardedClusterSnapshotViewStatus_COMPLETED ApiAtlasDiskBackupShardedClusterSnapshotViewStatus = "COMPLETED"
	// failed.
	ApiAtlasDiskBackupShardedClusterSnapshotViewStatus_FAILED ApiAtlasDiskBackupShardedClusterSnapshotViewStatus = "FAILED"
)

