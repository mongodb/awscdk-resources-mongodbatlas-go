package awscdkresourcesmongodbatlas


// Returns, adds, edits, and removes one connection for a stream instance in the specified project.
//
// To use this resource, the requesting API Key must have the Project Owner roles. Note that Atlas Streams functionality is currently in [Public Preview](https://www.mongodb.com/blog/post/atlas-stream-processing-now-in-public-preview)
type CfnStreamConnectionProps struct {
	// Human-readable label that identifies the stream connection.
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
	// Human-readable label that identifies the stream instance.
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// Unique 24-hexadecimal digit string that identifies your project.
	//
	// Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.
	//
	// **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Type of the connection.
	//
	// Can be either Cluster or Kafka.
	Type CfnStreamConnectionPropsType `field:"required" json:"type" yaml:"type"`
	Authentication *StreamsKafkaAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Comma separated list of server addresses.
	BootstrapServers *string `field:"optional" json:"bootstrapServers" yaml:"bootstrapServers"`
	// Name of the cluster configured for this connection.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	Config interface{} `field:"optional" json:"config" yaml:"config"`
	DbRoleToExecute *DbRoleToExecute `field:"optional" json:"dbRoleToExecute" yaml:"dbRoleToExecute"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	Security *StreamsKafkaSecurity `field:"optional" json:"security" yaml:"security"`
}

