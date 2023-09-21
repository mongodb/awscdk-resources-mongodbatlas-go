package awscdkresourcesmongodbatlas


// Query limit for one federated database instance.
type CfnFederatedQueryLimitProps struct {
	// Human-readable label that identifies this data federation instance limit.
	LimitName CfnFederatedQueryLimitPropsLimitName `field:"required" json:"limitName" yaml:"limitName"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Human-readable label that identifies the data federated database instance to which the query limit applies.
	TenantName *string `field:"required" json:"tenantName" yaml:"tenantName"`
	// Amount to set the limit to.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Only used for Data Federation limits.
	//
	// Action to take when the usage limit is exceeded. If limit span is set to QUERY, this is ignored because MongoDB Cloud stops the query when it exceeds the usage limit. "enum" : [ "BLOCK", "BLOCK_AND_KILL" ]
	OverrunPolicy *string `field:"optional" json:"overrunPolicy" yaml:"overrunPolicy"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

