// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


// The maintenanceWindow resource provides access to retrieve or update the current Atlas project maintenance window.
type CfnMaintenanceWindowProps struct {
	// Zero-based integer that represents the hour of the of the day that the maintenance window starts according to a 24-hour clock.
	//
	// Use `0` for midnight and `12` for noon.
	HourOfDay *float64 `field:"required" json:"hourOfDay" yaml:"hourOfDay"`
	// Flag that indicates whether MongoDB Cloud should defer all maintenance windows for one week after you enable them.
	AutoDeferOnceEnabled *bool `field:"optional" json:"autoDeferOnceEnabled" yaml:"autoDeferOnceEnabled"`
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
	DayOfWeek *float64 `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml)
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Flag that indicates whether MongoDB Cloud starts the maintenance window immediately upon receiving this request.
	//
	// To start the maintenance window immediately for your project, MongoDB Cloud must have maintenance scheduled and you must set a maintenance window. This flag resets to `false` after MongoDB Cloud completes maintenance.
	StartAsap *bool `field:"optional" json:"startAsap" yaml:"startAsap"`
}

