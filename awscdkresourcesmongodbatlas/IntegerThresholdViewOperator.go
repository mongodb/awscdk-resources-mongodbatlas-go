// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


// Comparison operator to apply when checking the current metric value.
type IntegerThresholdViewOperator string

const (
	// GREATER_THAN.
	IntegerThresholdViewOperator_GREATER_THAN IntegerThresholdViewOperator = "GREATER_THAN"
	// LESS_THAN.
	IntegerThresholdViewOperator_LESS_THAN IntegerThresholdViewOperator = "LESS_THAN"
)

