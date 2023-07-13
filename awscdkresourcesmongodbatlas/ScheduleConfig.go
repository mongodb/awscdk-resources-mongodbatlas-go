package awscdkresourcesmongodbatlas


type ScheduleConfig struct {
	// A [cron expression](https://www.mongodb.com/docs/atlas/app-services/triggers/scheduled-triggers/#cron-expressions) that specifies when the trigger executes.
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// If `true`, enabling the trigger after it was disabled will not invoke events that occurred while the trigger was disabled.
	SkipcatchupEvents *bool `field:"optional" json:"skipcatchupEvents" yaml:"skipcatchupEvents"`
}

