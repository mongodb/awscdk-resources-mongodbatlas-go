package awscdkresourcesmongodbatlas


// Returns, adds, edits, and removes Atlas Stream Processing Instances.
//
// Note that Atlas Streams functionality is currently in [Public Preview](https://www.mongodb.com/blog/post/atlas-stream-processing-now-in-public-preview)
type CfnStreamInstanceProps struct {
	DataProcessRegion *StreamsDataProcessRegion `field:"required" json:"dataProcessRegion" yaml:"dataProcessRegion"`
	// Human-readable label that identifies the stream connection.
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// Unique 24-hexadecimal character string that identifies the project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	StreamConfig *StreamConfig `field:"required" json:"streamConfig" yaml:"streamConfig"`
	Connections *[]*StreamsConnection `field:"optional" json:"connections" yaml:"connections"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

