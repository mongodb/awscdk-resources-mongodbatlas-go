package awscdkresourcesmongodbatlas


// Scheduled backup policy item configuration (hourly, daily, weekly, monthly, or yearly).
//
// When provided, FrequencyInterval, RetentionUnit, and RetentionValue are required.
type ScheduledPolicyItem struct {
	// Desired frequency of the new backup policy item specified by frequencyType.
	//
	// Required when the policy item is provided.
	FrequencyInterval *float64 `field:"optional" json:"frequencyInterval" yaml:"frequencyInterval"`
	// Frequency associated with the backup policy item.
	//
	// One of the following values: hourly, daily, weekly, monthly, or yearly. This is a read-only value.
	FrequencyType *string `field:"optional" json:"frequencyType" yaml:"frequencyType"`
	// Unique identifier of the backup policy item.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Unit of time in which MongoDB Cloud measures snapshot retention.
	//
	// Required when the policy item is provided.
	RetentionUnit ScheduledPolicyItemRetentionUnit `field:"optional" json:"retentionUnit" yaml:"retentionUnit"`
	// Duration in days, weeks, months, or years that MongoDB Cloud retains the snapshot.
	//
	// Required when the policy item is provided. For less frequent policy items, MongoDB Cloud requires that you specify a value greater than or equal to the value specified for more frequent policy items.
	RetentionValue *float64 `field:"optional" json:"retentionValue" yaml:"retentionValue"`
}

