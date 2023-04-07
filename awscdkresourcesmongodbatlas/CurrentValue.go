// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type CurrentValue struct {
	// Amount of the **metricName** recorded at the time of the event.
	//
	// This value triggered the alert.
	Number *float64 `field:"optional" json:"number" yaml:"number"`
	// Element used to express the quantity in **currentValue.number**. This can be an element of time, storage capacity, and the like. This metric triggered the alert.
	Units CurrentValueUnits `field:"optional" json:"units" yaml:"units"`
}

