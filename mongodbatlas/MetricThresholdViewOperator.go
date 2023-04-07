// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Comparison operator to apply when checking the current metric value.
type MetricThresholdViewOperator string

const (
	// GREATER_THAN.
	MetricThresholdViewOperator_GREATER_THAN MetricThresholdViewOperator = "GREATER_THAN"
	// LESS_THAN.
	MetricThresholdViewOperator_LESS_THAN MetricThresholdViewOperator = "LESS_THAN"
)

