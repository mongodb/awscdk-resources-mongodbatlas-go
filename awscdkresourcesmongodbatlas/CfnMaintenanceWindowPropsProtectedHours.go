package awscdkresourcesmongodbatlas


// Protected hours during which MongoDB Cloud cannot start maintenance.
type CfnMaintenanceWindowPropsProtectedHours struct {
	// Hour of the day when protected hours end (0-23).
	EndHourOfDay *float64 `field:"optional" json:"endHourOfDay" yaml:"endHourOfDay"`
	// Hour of the day when protected hours start (0-23).
	StartHourOfDay *float64 `field:"optional" json:"startHourOfDay" yaml:"startHourOfDay"`
}

