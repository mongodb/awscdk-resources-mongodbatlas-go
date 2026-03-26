package awscdkresourcesmongodbatlas


// Returns, adds, edits, and removes one connection for a stream workspace in the specified project.
//
// To use this resource, the requesting API Key must have the Project Owner roles.
type CfnStreamConnectionProps struct {
	// Human-readable label that identifies the stream connection.
	//
	// In the case of the Sample type, this is the name of the sample source.
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
	// Unique 24-hexadecimal digit string that identifies your project.
	//
	// Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.
	//
	// **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Type of the connection.
	//
	// Can be Cluster, Kafka, Sample, AWSLambda, or Https.
	Type CfnStreamConnectionPropsType `field:"required" json:"type" yaml:"type"`
	Authentication *StreamsKafkaAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	Aws *Aws `field:"optional" json:"aws" yaml:"aws"`
	// Comma separated list of server addresses.
	BootstrapServers *string `field:"optional" json:"bootstrapServers" yaml:"bootstrapServers"`
	// Name of the cluster configured for this connection.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Unique 24-hexadecimal digit string that identifies the project containing the cluster for cross-project cluster connections.
	ClusterProjectId *string `field:"optional" json:"clusterProjectId" yaml:"clusterProjectId"`
	Config interface{} `field:"optional" json:"config" yaml:"config"`
	DbRoleToExecute *DbRoleToExecute `field:"optional" json:"dbRoleToExecute" yaml:"dbRoleToExecute"`
	// HTTP headers for HTTPS type connections.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Human-readable label that identifies the stream instance.
	//
	// WARNING: This field is deprecated and will be removed in the next major release. Please use WorkspaceName instead.
	InstanceName *string `field:"optional" json:"instanceName" yaml:"instanceName"`
	Networking *Networking `field:"optional" json:"networking" yaml:"networking"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// The Schema Registry provider.
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	SchemaRegistryAuthentication *SchemaRegistryAuthentication `field:"optional" json:"schemaRegistryAuthentication" yaml:"schemaRegistryAuthentication"`
	// List of Schema Registry endpoint URLs.
	//
	// Each URL must use the http or https scheme and specify a valid host and optional port.
	SchemaRegistryUrls *[]*string `field:"optional" json:"schemaRegistryUrls" yaml:"schemaRegistryUrls"`
	Security *StreamsKafkaSecurity `field:"optional" json:"security" yaml:"security"`
	// URL endpoint for HTTPS type connections.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Human-readable label that identifies the stream workspace.
	WorkspaceName *string `field:"optional" json:"workspaceName" yaml:"workspaceName"`
}

