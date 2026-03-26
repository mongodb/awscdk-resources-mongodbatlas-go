package awscdkresourcesmongodbatlas


// Returns, adds, updates, and removes Atlas Stream Processing Workspaces.
//
// The DataProcessRegion.Region property can be updated after creation. Other properties (WorkspaceName, ProjectId, Profile, StreamConfig, DataProcessRegion.CloudProvider) are create-only and require resource replacement to change.
type CfnStreamWorkspaceProps struct {
	DataProcessRegion *StreamsDataProcessRegion `field:"required" json:"dataProcessRegion" yaml:"dataProcessRegion"`
	// Unique 24-hexadecimal character string that identifies the project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Human-readable label that identifies the stream workspace.
	WorkspaceName *string `field:"required" json:"workspaceName" yaml:"workspaceName"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	StreamConfig *StreamConfig `field:"optional" json:"streamConfig" yaml:"streamConfig"`
}

