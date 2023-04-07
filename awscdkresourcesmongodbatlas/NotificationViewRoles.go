// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type NotificationViewRoles string

const (
	// GROUP_CLUSTER_MANAGER.
	NotificationViewRoles_GROUP_CLUSTER_MANAGER NotificationViewRoles = "GROUP_CLUSTER_MANAGER"
	// GROUP_DATA_ACCESS_ADMIN.
	NotificationViewRoles_GROUP_DATA_ACCESS_ADMIN NotificationViewRoles = "GROUP_DATA_ACCESS_ADMIN"
	// GROUP_DATA_ACCESS_READ_ONLY.
	NotificationViewRoles_GROUP_DATA_ACCESS_READ_ONLY NotificationViewRoles = "GROUP_DATA_ACCESS_READ_ONLY"
	// GROUP_DATA_ACCESS_READ_WRITE.
	NotificationViewRoles_GROUP_DATA_ACCESS_READ_WRITE NotificationViewRoles = "GROUP_DATA_ACCESS_READ_WRITE"
	// GROUP_OWNER.
	NotificationViewRoles_GROUP_OWNER NotificationViewRoles = "GROUP_OWNER"
	// GROUP_READ_WRITE.
	NotificationViewRoles_GROUP_READ_WRITE NotificationViewRoles = "GROUP_READ_WRITE"
	// ORG_OWNER.
	NotificationViewRoles_ORG_OWNER NotificationViewRoles = "ORG_OWNER"
	// ORG_MEMBER.
	NotificationViewRoles_ORG_MEMBER NotificationViewRoles = "ORG_MEMBER"
	// ORG_GROUP_CREATOR.
	NotificationViewRoles_ORG_GROUP_CREATOR NotificationViewRoles = "ORG_GROUP_CREATOR"
	// ORG_BILLING_ADMIN.
	NotificationViewRoles_ORG_BILLING_ADMIN NotificationViewRoles = "ORG_BILLING_ADMIN"
	// ORG_READ_ONLY.
	NotificationViewRoles_ORG_READ_ONLY NotificationViewRoles = "ORG_READ_ONLY"
)

