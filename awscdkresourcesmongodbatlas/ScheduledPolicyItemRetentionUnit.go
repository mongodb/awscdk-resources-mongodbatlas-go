package awscdkresourcesmongodbatlas


// Unit of time in which MongoDB Cloud measures snapshot retention.
//
// Required when the policy item is provided.
type ScheduledPolicyItemRetentionUnit string

const (
	// days.
	ScheduledPolicyItemRetentionUnit_DAYS ScheduledPolicyItemRetentionUnit = "DAYS"
	// weeks.
	ScheduledPolicyItemRetentionUnit_WEEKS ScheduledPolicyItemRetentionUnit = "WEEKS"
	// months.
	ScheduledPolicyItemRetentionUnit_MONTHS ScheduledPolicyItemRetentionUnit = "MONTHS"
	// years.
	ScheduledPolicyItemRetentionUnit_YEARS ScheduledPolicyItemRetentionUnit = "YEARS"
)

