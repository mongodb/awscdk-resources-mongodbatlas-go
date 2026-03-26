package awscdkresourcesmongodbatlas


// The desired state of the stream processor.
//
// Used to start or stop the Stream Processor. Valid values are CREATED, STARTED or STOPPED. When a Stream Processor is created without specifying the desired state, it will default to CREATED state. When a Stream Processor is updated without specifying the desired state, it will default to the Previous state.
//
// **NOTE** When a Stream Processor is updated without specifying the desired state, it is stopped and then restored to previous state upon update completion.
type CfnStreamProcessorPropsDesiredState string

const (
	// CREATED.
	CfnStreamProcessorPropsDesiredState_CREATED CfnStreamProcessorPropsDesiredState = "CREATED"
	// STARTED.
	CfnStreamProcessorPropsDesiredState_STARTED CfnStreamProcessorPropsDesiredState = "STARTED"
	// STOPPED.
	CfnStreamProcessorPropsDesiredState_STOPPED CfnStreamProcessorPropsDesiredState = "STOPPED"
)

