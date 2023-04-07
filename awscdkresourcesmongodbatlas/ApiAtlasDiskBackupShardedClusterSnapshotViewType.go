// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


// Human-readable label that categorizes the cluster as a replica set or sharded cluster.
type ApiAtlasDiskBackupShardedClusterSnapshotViewType string

const (
	// REPLICA_SET.
	ApiAtlasDiskBackupShardedClusterSnapshotViewType_REPLICA_SET ApiAtlasDiskBackupShardedClusterSnapshotViewType = "REPLICA_SET"
	// SHARDED_CLUSTER.
	ApiAtlasDiskBackupShardedClusterSnapshotViewType_SHARDED_CLUSTER ApiAtlasDiskBackupShardedClusterSnapshotViewType = "SHARDED_CLUSTER"
)

