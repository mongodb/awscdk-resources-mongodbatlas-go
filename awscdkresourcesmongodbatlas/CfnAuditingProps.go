package awscdkresourcesmongodbatlas


// Returns and edits database auditing settings for MongoDB Cloud projects.
type CfnAuditingProps struct {
	// Unique 24-hexadecimal digit string that identifies your project.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

