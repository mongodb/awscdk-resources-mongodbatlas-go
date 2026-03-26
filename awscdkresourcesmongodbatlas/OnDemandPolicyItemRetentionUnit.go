package awscdkresourcesmongodbatlas


// Unit of time in which MongoDB Cloud measures snapshot retention.
//
// Required when OnDemandPolicyItem is provided.
type OnDemandPolicyItemRetentionUnit string

const (
	// days.
	OnDemandPolicyItemRetentionUnit_DAYS OnDemandPolicyItemRetentionUnit = "DAYS"
	// weeks.
	OnDemandPolicyItemRetentionUnit_WEEKS OnDemandPolicyItemRetentionUnit = "WEEKS"
	// months.
	OnDemandPolicyItemRetentionUnit_MONTHS OnDemandPolicyItemRetentionUnit = "MONTHS"
	// years.
	OnDemandPolicyItemRetentionUnit_YEARS OnDemandPolicyItemRetentionUnit = "YEARS"
)

