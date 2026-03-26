package awscdkresourcesmongodbatlas

import (
	"time"
)

// Grants temporary access to MongoDB employees for a specific cluster.
//
// This allows MongoDB support engineers to access cluster infrastructure, database logs, or app services sync data for troubleshooting purposes.
type CfnMongoDbEmployeeAccessGrantProps struct {
	// Human-readable label that identifies this cluster.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Expiration date for the employee access grant in ISO 8601 date and time format in UTC.
	ExpirationTime *time.Time `field:"required" json:"expirationTime" yaml:"expirationTime"`
	// Level of access to grant to MongoDB Employees.
	//
	// Possible values are CLUSTER_DATABASE_LOGS, CLUSTER_INFRASTRUCTURE or CLUSTER_INFRASTRUCTURE_AND_APP_SERVICES_SYNC_DATA.
	GrantType *string `field:"required" json:"grantType" yaml:"grantType"`
	// Unique 24-hexadecimal digit string that identifies your project.
	//
	// Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.
	//
	// **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

