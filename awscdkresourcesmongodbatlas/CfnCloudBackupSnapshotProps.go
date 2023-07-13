package awscdkresourcesmongodbatlas


// Returns, takes, and removes Cloud Backup snapshots.
type CfnCloudBackupSnapshotProps struct {
	// Human-readable label that identifies the cluster.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Human-readable phrase or sentence that explains the purpose of the snapshot.
	//
	// The resource returns this parameter when `"status": "onDemand"`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Human-readable label that identifies how often this snapshot triggers.
	FrequencyType CfnCloudBackupSnapshotPropsFrequencyType `field:"optional" json:"frequencyType" yaml:"frequencyType"`
	// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
	IncludeCount *bool `field:"optional" json:"includeCount" yaml:"includeCount"`
	// Human-readable label that identifies the serverless instance.
	InstanceName *string `field:"optional" json:"instanceName" yaml:"instanceName"`
	// Number of items that the response returns per page.
	ItemsPerPage *float64 `field:"optional" json:"itemsPerPage" yaml:"itemsPerPage"`
	// List that includes the snapshots and the cloud provider that stores the snapshots.
	//
	// The resource returns this parameter when `"type" : "SHARDED_CLUSTER"`.
	Members *[]*ApiAtlasDiskBackupShardedClusterSnapshotMemberView `field:"optional" json:"members" yaml:"members"`
	// Number of the page that displays the current set of the total objects that the response returns.
	PageNum *float64 `field:"optional" json:"pageNum" yaml:"pageNum"`
	// List that contains unique identifiers for the policy items.
	PolicyItems *[]*string `field:"optional" json:"policyItems" yaml:"policyItems"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// List of returned documents that MongoDB Cloud provides when completing this request.
	Results *[]*ApiAtlasDiskBackupShardedClusterSnapshotView `field:"optional" json:"results" yaml:"results"`
	// Number of days that MongoDB Cloud should retain the on-demand snapshot.
	//
	// Must be at least **1**.
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// Human-readable label that identifies when this snapshot triggers.
	SnapshotType CfnCloudBackupSnapshotPropsSnapshotType `field:"optional" json:"snapshotType" yaml:"snapshotType"`
	// Number of documents returned in this response.
	TotalCount *float64 `field:"optional" json:"totalCount" yaml:"totalCount"`
}

