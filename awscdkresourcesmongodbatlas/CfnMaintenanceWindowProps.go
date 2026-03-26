package awscdkresourcesmongodbatlas


// The maintenanceWindow resource provides access to retrieve or update the current Atlas project maintenance window.
type CfnMaintenanceWindowProps struct {
	// One-based integer that represents the day of the week that the maintenance window starts.
	//
	// | Value | Day of Week |
	// |---|---|
	// | `1` | Sunday |
	// | `2` | Monday |
	// | `3` | Tuesday |
	// | `4` | Wednesday |
	// | `5` | Thursday |
	// | `6` | Friday |
	// | `7` | Saturday |.
	DayOfWeek *float64 `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Zero-based integer that represents the hour of the of the day that the maintenance window starts according to a 24-hour clock.
	//
	// Use `0` for midnight and `12` for noon.
	HourOfDay *float64 `field:"required" json:"hourOfDay" yaml:"hourOfDay"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Flag that indicates whether MongoDB Cloud should automatically defer maintenance windows for one week when they occur during the defined maintenance window.
	AutoDefer *bool `field:"optional" json:"autoDefer" yaml:"autoDefer"`
	// Flag that indicates whether MongoDB Cloud should defer all maintenance windows for one week after you enable them.
	AutoDeferOnceEnabled *bool `field:"optional" json:"autoDeferOnceEnabled" yaml:"autoDeferOnceEnabled"`
	// Flag that indicates whether to defer the maintenance window.
	//
	// When set to true, the next scheduled maintenance will be deferred.
	Defer *bool `field:"optional" json:"defer" yaml:"defer"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml)
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Protected hours during which MongoDB Cloud cannot start maintenance.
	ProtectedHours *CfnMaintenanceWindowPropsProtectedHours `field:"optional" json:"protectedHours" yaml:"protectedHours"`
}

