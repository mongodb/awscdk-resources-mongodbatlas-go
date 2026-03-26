package awscdkresourcesmongodbatlas


// Resource for managing MongoDB Atlas Backup Compliance Policy.
//
// Backup Compliance Policy prevents any user, regardless of role, from modifying or deleting specific cluster settings, backups, and backup configurations. When enabled, the Backup Compliance Policy will be applied as the minimum policy for all clusters and backups in the project.
type CfnBackupCompliancePolicyProps struct {
	// Email address of the user authorized to update the Backup Compliance Policy settings.
	AuthorizedEmail *string `field:"required" json:"authorizedEmail" yaml:"authorizedEmail"`
	// First name of the user authorized to update the Backup Compliance Policy settings.
	AuthorizedUserFirstName *string `field:"required" json:"authorizedUserFirstName" yaml:"authorizedUserFirstName"`
	// Last name of the user authorized to update the Backup Compliance Policy settings.
	AuthorizedUserLastName *string `field:"required" json:"authorizedUserLastName" yaml:"authorizedUserLastName"`
	// The unique identifier of the project for the Backup Compliance Policy.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Flag that indicates whether to enable additional copy protection for the cluster.
	//
	// If enabled, cloud backup snapshots cannot be deleted until the retention period expires. Defaults to false if not specified.
	// Default: false if not specified.
	//
	CopyProtectionEnabled *bool `field:"optional" json:"copyProtectionEnabled" yaml:"copyProtectionEnabled"`
	// Flag that indicates whether Encryption at Rest using Customer Key Management is required for all clusters with a Backup Compliance Policy.
	//
	// Defaults to false if not specified.
	// Default: false if not specified.
	//
	EncryptionAtRestEnabled *bool `field:"optional" json:"encryptionAtRestEnabled" yaml:"encryptionAtRestEnabled"`
	// On-demand backup policy item configuration.
	OnDemandPolicyItem *OnDemandPolicyItem `field:"optional" json:"onDemandPolicyItem" yaml:"onDemandPolicyItem"`
	// Flag that indicates whether the cluster uses Continuous Cloud Backup.
	//
	// If enabled, cloud backup snapshots are taken every 6 hours. Defaults to false if not specified.
	// Default: false if not specified.
	//
	PitEnabled *bool `field:"optional" json:"pitEnabled" yaml:"pitEnabled"`
	// Daily backup policy item configuration.
	//
	// Only one daily policy item is allowed.
	PolicyItemDaily *ScheduledPolicyItem `field:"optional" json:"policyItemDaily" yaml:"policyItemDaily"`
	// Hourly backup policy item configuration.
	//
	// Only one hourly policy item is allowed.
	PolicyItemHourly *ScheduledPolicyItem `field:"optional" json:"policyItemHourly" yaml:"policyItemHourly"`
	// Monthly backup policy item configuration.
	//
	// Multiple monthly policy items are allowed.
	PolicyItemMonthly *[]*ScheduledPolicyItem `field:"optional" json:"policyItemMonthly" yaml:"policyItemMonthly"`
	// Weekly backup policy item configuration.
	//
	// Multiple weekly policy items are allowed.
	PolicyItemWeekly *[]*ScheduledPolicyItem `field:"optional" json:"policyItemWeekly" yaml:"policyItemWeekly"`
	// Yearly backup policy item configuration.
	//
	// Multiple yearly policy items are allowed.
	PolicyItemYearly *[]*ScheduledPolicyItem `field:"optional" json:"policyItemYearly" yaml:"policyItemYearly"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Number of days back in time you can restore to with Continuous Cloud Backup accuracy.
	//
	// Must be a positive, non-zero integer. This field is optional and computed from the API.
	RestoreWindowDays *float64 `field:"optional" json:"restoreWindowDays" yaml:"restoreWindowDays"`
}

