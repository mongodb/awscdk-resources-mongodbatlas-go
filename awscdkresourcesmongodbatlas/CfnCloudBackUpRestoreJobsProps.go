// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


// Returns, starts, and cancels Cloud Backup restore jobs.
type CfnCloudBackUpRestoreJobsProps struct {
	// The name of the Atlas cluster whose snapshot you want to restore or you want to retrieve restore jobs.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The unique identifier of the project for the Atlas cluster.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Indicates whether the restore job was canceled.
	Cancelled *bool `field:"optional" json:"cancelled" yaml:"cancelled"`
	// Type of restore job to create.The value can be any one of download,automated or point_in_time.
	DeliveryType CfnCloudBackUpRestoreJobsPropsDeliveryType `field:"optional" json:"deliveryType" yaml:"deliveryType"`
	// Indicates whether the restore job expired.
	Expired *bool `field:"optional" json:"expired" yaml:"expired"`
	// The instance name of the Serverless cluster whose snapshot you want to restore or you want to retrieve restore jobs.
	InstanceName *string `field:"optional" json:"instanceName" yaml:"instanceName"`
	// Oplog operation number from which to you want to restore this snapshot.
	//
	// This is the second part of an Oplog timestamp.
	OpLogInc *string `field:"optional" json:"opLogInc" yaml:"opLogInc"`
	// Timestamp in the number of seconds that have elapsed since the UNIX epoch from which to you want to restore this snapshot.
	//
	// This is the first part of an Oplog timestamp.
	OpLogTs *string `field:"optional" json:"opLogTs" yaml:"opLogTs"`
	// If you performed a Point-in-Time restores at a time specified by a Unix time in seconds since epoch, pointInTimeUTCSeconds indicates the Unix time used.
	PointInTimeUtcSeconds *float64 `field:"optional" json:"pointInTimeUtcSeconds" yaml:"pointInTimeUtcSeconds"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Unique identifier of the source snapshot ID of the restore job.
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// Name of the target Atlas cluster to which the restore job restores the snapshot.
	//
	// Only visible if deliveryType is automated.
	TargetClusterName *string `field:"optional" json:"targetClusterName" yaml:"targetClusterName"`
	// Name of the target Atlas project of the restore job.
	//
	// Only visible if deliveryType is automated.
	TargetProjectId *string `field:"optional" json:"targetProjectId" yaml:"targetProjectId"`
}

