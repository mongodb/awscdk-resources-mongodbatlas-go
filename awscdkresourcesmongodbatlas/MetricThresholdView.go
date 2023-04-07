// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type MetricThresholdView struct {
	// Human-readable label that identifies the metric against which MongoDB Cloud checks the configured **metricThreshold.threshold**.
	MetricName MetricThresholdViewMetricName `field:"optional" json:"metricName" yaml:"metricName"`
	// MongoDB Cloud computes the current metric value as an average.
	Mode MetricThresholdViewMode `field:"optional" json:"mode" yaml:"mode"`
	// Comparison operator to apply when checking the current metric value.
	Operator MetricThresholdViewOperator `field:"optional" json:"operator" yaml:"operator"`
	// Value of metric that, when exceeded, triggers an alert.
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// Element used to express the quantity.
	//
	// This can be an element of time, storage capacity, and the like.
	Units MetricThresholdViewUnits `field:"optional" json:"units" yaml:"units"`
}

