package awscdkresourcesmongodbatlas


type ApiPolicyItemView struct {
	// Desired frequency of the new backup policy item specified by frequencyType.
	FrequencyInterval *float64 `field:"optional" json:"frequencyInterval" yaml:"frequencyInterval"`
	// Frequency associated with the backup policy item.
	//
	// One of the following values: hourly, daily, weekly, monthly or yearly.
	FrequencyType *string `field:"optional" json:"frequencyType" yaml:"frequencyType"`
	// Unique identifier of the backup policy item.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Metric of duration of the backup policy item: days, weeks, months or years.
	RetentionUnit *string `field:"optional" json:"retentionUnit" yaml:"retentionUnit"`
	// Duration for which the backup is kept.
	//
	// Associated with retentionUnit.
	RetentionValue *float64 `field:"optional" json:"retentionValue" yaml:"retentionValue"`
}

