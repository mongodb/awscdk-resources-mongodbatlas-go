// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// State of this alert at the time you requested its details.
type AlertViewStatus string

const (
	// CANCELLED.
	AlertViewStatus_CANCELLED AlertViewStatus = "CANCELLED"
	// CLOSED.
	AlertViewStatus_CLOSED AlertViewStatus = "CLOSED"
	// OPEN.
	AlertViewStatus_OPEN AlertViewStatus = "OPEN"
	// TRACKING.
	AlertViewStatus_TRACKING AlertViewStatus = "TRACKING"
)

