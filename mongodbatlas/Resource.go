// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// List of resources on which you grant the action.
type Resource struct {
	// Flag that indicates whether to grant the action on the cluster resource.
	//
	// If true, MongoDB Cloud ignores the actions.resources.collection and actions.resources.db parameters.
	Cluster *bool `field:"optional" json:"cluster" yaml:"cluster"`
	// Human-readable label that identifies the collection on which you grant the action to one MongoDB user.
	//
	// If you don't set this parameter, you grant the action to all collections in the database specified in the actions.resources.db parameter. If you set "actions.resources.cluster" : true, MongoDB Cloud ignores this parameter.
	Collection *string `field:"optional" json:"collection" yaml:"collection"`
	// Human-readable label that identifies the database on which you grant the action to one MongoDB user.
	//
	// If you set "actions.resources.cluster" : true, MongoDB Cloud ignores this parameter.
	Db *string `field:"optional" json:"db" yaml:"db"`
}

