package awscdkresourcesmongodbatlas


type ProjectApiKey struct {
	// Unique 24-hexadecimal digit string that identifies this organization API key assigned to this project.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// List of roles to grant this API key.
	//
	// If you provide this list, provide a minimum of one role and ensure each role applies to this project.Items Enum: "ORG_OWNER" "ORG_MEMBER" "ORG_GROUP_CREATOR" "ORG_BILLING_ADMIN" "ORG_READ_ONLY" "ORG_TEAM_MEMBERS_ADMIN" "GROUP_ATLAS_ADMIN" "GROUP_AUTOMATION_ADMIN" "GROUP_BACKUP_ADMIN" "GROUP_MONITORING_ADMIN" "GROUP_OWNER" "GROUP_READ_ONLY" "GROUP_USER_ADMIN" "GROUP_BILLING_ADMIN" "GROUP_DATA_ACCESS_ADMIN" "GROUP_DATA_ACCESS_READ_ONLY" "GROUP_DATA_ACCESS_READ_WRITE" "GROUP_CHARTS_ADMIN" "GROUP_CLUSTER_MANAGER" "GROUP_SEARCH_INDEX_EDITOR"
	RoleNames *[]*string `field:"optional" json:"roleNames" yaml:"roleNames"`
}

