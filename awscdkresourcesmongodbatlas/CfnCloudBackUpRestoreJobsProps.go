package awscdkresourcesmongodbatlas


// Returns, starts, and cancels Cloud Backup restore jobs.
type CfnCloudBackUpRestoreJobsProps struct {
	// Type of restore job to create.The value can be any one of download,automated or point_in_time.
	DeliveryType CfnCloudBackUpRestoreJobsPropsDeliveryType `field:"required" json:"deliveryType" yaml:"deliveryType"`
	// The instance name of the Serverless/Cluster whose snapshot you want to restore or you want to retrieve restore jobs.
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// Type of instance specified on the Instance Name serverless or cluster.
	//
	// **WARNING:** `serverless` instance type is deprecated and will be removed in January 2026. For more details, see [Migrate your programmatic tools from M2, M5, or Serverless Instances to Flex Clusters](https://www.mongodb.com/docs/atlas/flex-migration/).
	InstanceType CfnCloudBackUpRestoreJobsPropsInstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// The unique identifier of the project for the Atlas cluster.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Unique identifier of the source snapshot ID of the restore job.
	SnapshotId *string `field:"required" json:"snapshotId" yaml:"snapshotId"`
	// If set to true, the CloudFormation resource will wait until the job is completed, WARNING: if the snapshot has a big load of data, the cloud formation resource might take a long time to finish leading to high costs.
	EnableSynchronousCreation *bool `field:"optional" json:"enableSynchronousCreation" yaml:"enableSynchronousCreation"`
	// One or more links to sub-resources and/or related resources.
	Links *[]*CfnCloudBackUpRestoreJobsPropsLinks `field:"optional" json:"links" yaml:"links"`
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
	// Options that needs to be set to control the synchronous creation flow, this options need to be set if EnableSynchronousCreation is se to TRUE.
	SynchronousCreationOptions *SynchronousCreationOptions `field:"optional" json:"synchronousCreationOptions" yaml:"synchronousCreationOptions"`
	// Name of the target Atlas cluster to which the restore job restores the snapshot.
	//
	// Only visible if deliveryType is automated.
	TargetClusterName *string `field:"optional" json:"targetClusterName" yaml:"targetClusterName"`
	// Name of the target Atlas project of the restore job.
	//
	// Only visible if deliveryType is automated.
	TargetProjectId *string `field:"optional" json:"targetProjectId" yaml:"targetProjectId"`
}

