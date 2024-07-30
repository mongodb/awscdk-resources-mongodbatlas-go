package awscdkresourcesmongodbatlas


// Returns, adds, and removes Global Cluster managed namespaces and custom zone mappings.
//
// This resource can only be used with Atlas-managed clusters, see doc for `GlobalClusterSelfManagedSharding` attribute in `Mongodb::Atlas::Cluster` resource for more info.
type CfnGlobalClusterConfigProps struct {
	// The name of the Atlas cluster that contains the snapshots you want to retrieve.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// List that contains comma-separated key value pairs to map zones to geographic regions.
	//
	// These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to the human-readable label for the desired custom zone. MongoDB Cloud maps the ISO 3166-1a2 code to the nearest geographical zone by default. Include this parameter to override the default mappings.
	//
	// This parameter returns an empty object if no custom zones exist.
	CustomZoneMappings *[]*ZoneMapping `field:"optional" json:"customZoneMappings" yaml:"customZoneMappings"`
	// List that contains comma-separated key value pairs to map zones to geographic regions.
	//
	// These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to the human-readable label for the desired custom zone. MongoDB Cloud maps the ISO 3166-1a2 code to the nearest geographical zone by default. Include this parameter to override the default mappings.
	//
	// This parameter returns an empty object if no custom zones exist.
	ManagedNamespaces *[]*ManagedNamespace `field:"optional" json:"managedNamespaces" yaml:"managedNamespaces"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// The unique identifier of the project for the Atlas cluster.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

