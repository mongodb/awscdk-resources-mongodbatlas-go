// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Element used to express the quantity in **currentValue.number**. This can be an element of time, storage capacity, and the like. This metric triggered the alert.
type CurrentValueUnits string

const (
	// BITS.
	CurrentValueUnits_BITS CurrentValueUnits = "BITS"
	// BYTES.
	CurrentValueUnits_BYTES CurrentValueUnits = "BYTES"
	// DAYS.
	CurrentValueUnits_DAYS CurrentValueUnits = "DAYS"
	// GIGABITS.
	CurrentValueUnits_GIGABITS CurrentValueUnits = "GIGABITS"
	// GIGABYTES.
	CurrentValueUnits_GIGABYTES CurrentValueUnits = "GIGABYTES"
	// HOURS.
	CurrentValueUnits_HOURS CurrentValueUnits = "HOURS"
	// KILOBITS.
	CurrentValueUnits_KILOBITS CurrentValueUnits = "KILOBITS"
	// KILOBYTES.
	CurrentValueUnits_KILOBYTES CurrentValueUnits = "KILOBYTES"
	// MEGABITS.
	CurrentValueUnits_MEGABITS CurrentValueUnits = "MEGABITS"
	// MEGABYTES.
	CurrentValueUnits_MEGABYTES CurrentValueUnits = "MEGABYTES"
	// MILLISECONDS.
	CurrentValueUnits_MILLISECONDS CurrentValueUnits = "MILLISECONDS"
	// MINUTES.
	CurrentValueUnits_MINUTES CurrentValueUnits = "MINUTES"
	// PETABYTES.
	CurrentValueUnits_PETABYTES CurrentValueUnits = "PETABYTES"
	// RAW.
	CurrentValueUnits_RAW CurrentValueUnits = "RAW"
	// SECONDS.
	CurrentValueUnits_SECONDS CurrentValueUnits = "SECONDS"
	// TERABYTES.
	CurrentValueUnits_TERABYTES CurrentValueUnits = "TERABYTES"
)

