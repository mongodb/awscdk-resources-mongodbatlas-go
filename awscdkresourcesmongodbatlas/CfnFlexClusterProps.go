package awscdkresourcesmongodbatlas


// The flex cluster resource provides access to your flex cluster configurations.
//
// The resource lets you create, edit and delete flex clusters. The resource requires your Project ID.
type CfnFlexClusterProps struct {
	// Human-readable label that identifies the flex cluster.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Unique identifier of the project the cluster belongs to.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	ProviderSettings *ProviderSettings `field:"required" json:"providerSettings" yaml:"providerSettings"`
	BackupSettings *BackupSettings `field:"optional" json:"backupSettings" yaml:"backupSettings"`
	ConnectionStrings *ConnectionStrings `field:"optional" json:"connectionStrings" yaml:"connectionStrings"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Map that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the flex cluster.
	Tags *[]*Tag `field:"optional" json:"tags" yaml:"tags"`
	// Flag that indicates whether termination protection is enabled on the cluster.
	//
	// If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster.
	TerminationProtectionEnabled *bool `field:"optional" json:"terminationProtectionEnabled" yaml:"terminationProtectionEnabled"`
}

