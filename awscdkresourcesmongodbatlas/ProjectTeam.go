package awscdkresourcesmongodbatlas


type ProjectTeam struct {
	// One or more organization- or project-level roles to assign to the MongoDB Cloud user.
	//
	// tems Enum: "GROUP_CLUSTER_MANAGER" "GROUP_DATA_ACCESS_ADMIN" "GROUP_DATA_ACCESS_READ_ONLY" "GROUP_DATA_ACCESS_READ_WRITE" "GROUP_OWNER" "GROUP_READ_ONLY".
	RoleNames *[]*string `field:"optional" json:"roleNames" yaml:"roleNames"`
	// Unique 24-hexadecimal character string that identifies the team.
	//
	// string = 24 characters ^([a-f0-9]{24})$.
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
}

