package awscdkresourcesmongodbatlas


type ScheduleView struct {
	// Day of the month when the scheduled archive starts.
	DayOfMonth *float64 `field:"optional" json:"dayOfMonth" yaml:"dayOfMonth"`
	// Day of the month when the scheduled archive starts.
	DayOfWeek *float64 `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Hour of the day when the scheduled window to run one online archive ends.
	EndHour *float64 `field:"optional" json:"endHour" yaml:"endHour"`
	// Minute of the hour when the scheduled window to run one online archive ends.
	EndMinute *float64 `field:"optional" json:"endMinute" yaml:"endMinute"`
	// Hour of the day when the when the scheduled window to run one online archive starts.
	StartHour *float64 `field:"optional" json:"startHour" yaml:"startHour"`
	// Minute of the hour when the scheduled window to run one online archive starts.
	StartMinute *float64 `field:"optional" json:"startMinute" yaml:"startMinute"`
	Type ScheduleViewType `field:"optional" json:"type" yaml:"type"`
}

