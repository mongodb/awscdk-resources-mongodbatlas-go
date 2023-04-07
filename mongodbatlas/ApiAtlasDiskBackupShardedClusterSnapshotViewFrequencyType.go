// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Human-readable label that identifies how often this snapshot triggers.
type ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType string

const (
	// hourly.
	ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType_HOURLY ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType = "HOURLY"
	// daily.
	ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType_DAILY ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType = "DAILY"
	// weekly.
	ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType_WEEKLY ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType = "WEEKLY"
	// monthly.
	ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType_MONTHLY ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType = "MONTHLY"
)

