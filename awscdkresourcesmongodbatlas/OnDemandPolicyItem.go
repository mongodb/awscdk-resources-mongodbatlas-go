package awscdkresourcesmongodbatlas


// On-demand backup policy item configuration.
//
// When provided, FrequencyInterval, RetentionUnit, and RetentionValue are required.
type OnDemandPolicyItem struct {
	// Number that indicates the frequency interval for a set of snapshots.
	//
	// Required when OnDemandPolicyItem is provided.
	FrequencyInterval *float64 `field:"optional" json:"frequencyInterval" yaml:"frequencyInterval"`
	// Frequency associated with the backup policy item.
	//
	// For on-demand policies, the frequency type is 'ondemand'. This is a read-only value.
	FrequencyType *string `field:"optional" json:"frequencyType" yaml:"frequencyType"`
	// Unique identifier of the backup policy item.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Unit of time in which MongoDB Cloud measures snapshot retention.
	//
	// Required when OnDemandPolicyItem is provided.
	RetentionUnit OnDemandPolicyItemRetentionUnit `field:"optional" json:"retentionUnit" yaml:"retentionUnit"`
	// Duration in days, weeks, months, or years that MongoDB Cloud retains the snapshot.
	//
	// Required when OnDemandPolicyItem is provided.
	RetentionValue *float64 `field:"optional" json:"retentionValue" yaml:"retentionValue"`
}

