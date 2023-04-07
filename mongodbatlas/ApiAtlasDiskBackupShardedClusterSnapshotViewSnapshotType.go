// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Human-readable label that identifies when this snapshot triggers.
type ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType string

const (
	// onDemand.
	ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType_ON_DEMAND ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType = "ON_DEMAND"
	// scheduled.
	ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType_SCHEDULED ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType = "SCHEDULED"
)

