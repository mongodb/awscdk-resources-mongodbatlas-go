// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type IntegerThresholdView struct {
	// Comparison operator to apply when checking the current metric value.
	Operator IntegerThresholdViewOperator `field:"optional" json:"operator" yaml:"operator"`
	// Value of metric that, when exceeded, triggers an alert.
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// Element used to express the quantity.
	//
	// This can be an element of time, storage capacity, and the like.
	Units IntegerThresholdViewUnits `field:"optional" json:"units" yaml:"units"`
}

