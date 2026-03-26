package awscdkresourcesmongodbatlas


// Returns, adds, edits, and removes Atlas Stream Processors.
type CfnStreamProcessorProps struct {
	// Stream aggregation pipeline you want to apply to your streaming data.
	//
	// This should be a JSON-encoded array of pipeline stages. Refer to MongoDB Atlas Docs for more information on stream aggregation pipelines.
	Pipeline *string `field:"required" json:"pipeline" yaml:"pipeline"`
	// Label that identifies the stream processor.
	ProcessorName *string `field:"required" json:"processorName" yaml:"processorName"`
	// Unique 24-hexadecimal digit string that identifies your project.
	//
	// **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Label that identifies the stream processing workspace.
	WorkspaceName *string `field:"required" json:"workspaceName" yaml:"workspaceName"`
	// Indicates whether to delete the resource being created if a timeout is reached when waiting for completion.
	//
	// When set to `true` and timeout occurs, it triggers the deletion and returns immediately without waiting for deletion to complete. When set to `false`, the timeout will not trigger resource deletion. If you suspect a transient error when the value is `true`, wait before retrying to allow resource deletion to finish. Default is `true`.
	// Default: true`.
	//
	DeleteOnCreateTimeout *bool `field:"optional" json:"deleteOnCreateTimeout" yaml:"deleteOnCreateTimeout"`
	// The desired state of the stream processor.
	//
	// Used to start or stop the Stream Processor. Valid values are CREATED, STARTED or STOPPED. When a Stream Processor is created without specifying the desired state, it will default to CREATED state. When a Stream Processor is updated without specifying the desired state, it will default to the Previous state.
	//
	// **NOTE** When a Stream Processor is updated without specifying the desired state, it is stopped and then restored to previous state upon update completion.
	DesiredState CfnStreamProcessorPropsDesiredState `field:"optional" json:"desiredState" yaml:"desiredState"`
	Options *StreamsOptions `field:"optional" json:"options" yaml:"options"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Configurable timeouts for stream processor operations.
	Timeouts *Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

