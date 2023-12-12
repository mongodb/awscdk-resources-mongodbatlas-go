package awscdkresourcesmongodbatlas


// An example resource schema demonstrating some basic constructs and validation rules.
type CfnClusterOutageSimulationProps struct {
	// Human-readable label that identifies the cluster .
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// List of settings that configure your cluster regions.
	//
	// For Global Clusters, each object in the array represents a zone where your clusters nodes deploy. For non-Global replica sets and sharded clusters, this array has one object representing where your clusters nodes deploy.
	OutageFilters *[]*Filter `field:"required" json:"outageFilters" yaml:"outageFilters"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"required" json:"profile" yaml:"profile"`
	// Human-readable label that identifies the project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

