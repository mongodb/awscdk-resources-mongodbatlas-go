package awscdkresourcesmongodbatlas


// Human-readable label that identifies the collection of privileges that MongoDB Cloud grants a specific API key, MongoDB Cloud user, or MongoDB Cloud team.
//
// These roles include organization- and project-level roles.
//
// Organization Roles
//
// * ORG_OWNER
// * ORG_MEMBER
// * ORG_GROUP_CREATOR
// * ORG_BILLING_ADMIN
// * ORG_READ_ONLY
//
// Project Roles
//
// * GROUP_CLUSTER_MANAGER
// * GROUP_DATA_ACCESS_ADMIN
// * GROUP_DATA_ACCESS_READ_ONLY
// * GROUP_DATA_ACCESS_READ_WRITE
// * GROUP_OWNER
// * GROUP_READ_ONLY.
type AtlasRoleRoleName string

const (
	// ORG_OWNER.
	AtlasRoleRoleName_ORG_OWNER AtlasRoleRoleName = "ORG_OWNER"
	// ORG_MEMBER.
	AtlasRoleRoleName_ORG_MEMBER AtlasRoleRoleName = "ORG_MEMBER"
	// ORG_GROUP_CREATOR.
	AtlasRoleRoleName_ORG_GROUP_CREATOR AtlasRoleRoleName = "ORG_GROUP_CREATOR"
	// ORG_BILLING_ADMIN.
	AtlasRoleRoleName_ORG_BILLING_ADMIN AtlasRoleRoleName = "ORG_BILLING_ADMIN"
	// ORG_READ_ONLY.
	AtlasRoleRoleName_ORG_READ_ONLY AtlasRoleRoleName = "ORG_READ_ONLY"
	// GROUP_CLUSTER_MANAGER.
	AtlasRoleRoleName_GROUP_CLUSTER_MANAGER AtlasRoleRoleName = "GROUP_CLUSTER_MANAGER"
	// GROUP_DATA_ACCESS_ADMIN.
	AtlasRoleRoleName_GROUP_DATA_ACCESS_ADMIN AtlasRoleRoleName = "GROUP_DATA_ACCESS_ADMIN"
	// GROUP_DATA_ACCESS_READ_ONLY.
	AtlasRoleRoleName_GROUP_DATA_ACCESS_READ_ONLY AtlasRoleRoleName = "GROUP_DATA_ACCESS_READ_ONLY"
	// GROUP_DATA_ACCESS_READ_WRITE.
	AtlasRoleRoleName_GROUP_DATA_ACCESS_READ_WRITE AtlasRoleRoleName = "GROUP_DATA_ACCESS_READ_WRITE"
	// GROUP_OWNER.
	AtlasRoleRoleName_GROUP_OWNER AtlasRoleRoleName = "GROUP_OWNER"
	// GROUP_READ_ONLY.
	AtlasRoleRoleName_GROUP_READ_ONLY AtlasRoleRoleName = "GROUP_READ_ONLY"
)

