// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


// Element used to express the quantity.
//
// This can be an element of time, storage capacity, and the like.
type IntegerThresholdViewUnits string

const (
	// BITS.
	IntegerThresholdViewUnits_BITS IntegerThresholdViewUnits = "BITS"
	// BYTES.
	IntegerThresholdViewUnits_BYTES IntegerThresholdViewUnits = "BYTES"
	// DAYS.
	IntegerThresholdViewUnits_DAYS IntegerThresholdViewUnits = "DAYS"
	// GIGABITS.
	IntegerThresholdViewUnits_GIGABITS IntegerThresholdViewUnits = "GIGABITS"
	// GIGABYTES.
	IntegerThresholdViewUnits_GIGABYTES IntegerThresholdViewUnits = "GIGABYTES"
	// HOURS.
	IntegerThresholdViewUnits_HOURS IntegerThresholdViewUnits = "HOURS"
	// KILOBITS.
	IntegerThresholdViewUnits_KILOBITS IntegerThresholdViewUnits = "KILOBITS"
	// KILOBYTES.
	IntegerThresholdViewUnits_KILOBYTES IntegerThresholdViewUnits = "KILOBYTES"
	// MEGABITS.
	IntegerThresholdViewUnits_MEGABITS IntegerThresholdViewUnits = "MEGABITS"
	// MEGABYTES.
	IntegerThresholdViewUnits_MEGABYTES IntegerThresholdViewUnits = "MEGABYTES"
	// MILLISECONDS.
	IntegerThresholdViewUnits_MILLISECONDS IntegerThresholdViewUnits = "MILLISECONDS"
	// MINUTES.
	IntegerThresholdViewUnits_MINUTES IntegerThresholdViewUnits = "MINUTES"
	// PETABYTES.
	IntegerThresholdViewUnits_PETABYTES IntegerThresholdViewUnits = "PETABYTES"
	// RAW.
	IntegerThresholdViewUnits_RAW IntegerThresholdViewUnits = "RAW"
	// SECONDS.
	IntegerThresholdViewUnits_SECONDS IntegerThresholdViewUnits = "SECONDS"
	// TERABYTES.
	IntegerThresholdViewUnits_TERABYTES IntegerThresholdViewUnits = "TERABYTES"
)

