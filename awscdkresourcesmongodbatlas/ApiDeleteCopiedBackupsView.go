// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type ApiDeleteCopiedBackupsView struct {
	// A label that identifies the cloud provider for the deleted copy setting whose backup copies you want to delete.
	CloudProvider *string `field:"optional" json:"cloudProvider" yaml:"cloudProvider"`
	// Target region for the deleted copy setting whose backup copies you want to delete.
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster.
	ReplicationSpecId *string `field:"optional" json:"replicationSpecId" yaml:"replicationSpecId"`
}

