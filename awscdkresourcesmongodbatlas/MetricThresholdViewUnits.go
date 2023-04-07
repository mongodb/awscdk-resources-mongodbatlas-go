// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


// Element used to express the quantity.
//
// This can be an element of time, storage capacity, and the like.
type MetricThresholdViewUnits string

const (
	// BITS.
	MetricThresholdViewUnits_BITS MetricThresholdViewUnits = "BITS"
	// BYTES.
	MetricThresholdViewUnits_BYTES MetricThresholdViewUnits = "BYTES"
	// DAYS.
	MetricThresholdViewUnits_DAYS MetricThresholdViewUnits = "DAYS"
	// GIGABITS.
	MetricThresholdViewUnits_GIGABITS MetricThresholdViewUnits = "GIGABITS"
	// GIGABYTES.
	MetricThresholdViewUnits_GIGABYTES MetricThresholdViewUnits = "GIGABYTES"
	// HOURS.
	MetricThresholdViewUnits_HOURS MetricThresholdViewUnits = "HOURS"
	// KILOBITS.
	MetricThresholdViewUnits_KILOBITS MetricThresholdViewUnits = "KILOBITS"
	// KILOBYTES.
	MetricThresholdViewUnits_KILOBYTES MetricThresholdViewUnits = "KILOBYTES"
	// MEGABITS.
	MetricThresholdViewUnits_MEGABITS MetricThresholdViewUnits = "MEGABITS"
	// MEGABYTES.
	MetricThresholdViewUnits_MEGABYTES MetricThresholdViewUnits = "MEGABYTES"
	// MILLISECONDS.
	MetricThresholdViewUnits_MILLISECONDS MetricThresholdViewUnits = "MILLISECONDS"
	// MINUTES.
	MetricThresholdViewUnits_MINUTES MetricThresholdViewUnits = "MINUTES"
	// PETABYTES.
	MetricThresholdViewUnits_PETABYTES MetricThresholdViewUnits = "PETABYTES"
	// RAW.
	MetricThresholdViewUnits_RAW MetricThresholdViewUnits = "RAW"
	// SECONDS.
	MetricThresholdViewUnits_SECONDS MetricThresholdViewUnits = "SECONDS"
	// TERABYTES.
	MetricThresholdViewUnits_TERABYTES MetricThresholdViewUnits = "TERABYTES"
)

