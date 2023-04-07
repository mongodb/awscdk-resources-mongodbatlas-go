// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


// View and manage your application's [triggers](https://www.mongodb.com/docs/atlas/app-services/triggers/overview/).
type CfnTriggerProps struct {
	// App Services Application ID.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Project Id for application services.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	AuthTrigger *AuthConfig `field:"optional" json:"authTrigger" yaml:"authTrigger"`
	DatabaseTrigger *DatabaseConfig `field:"optional" json:"databaseTrigger" yaml:"databaseTrigger"`
	// If `true`, the trigger is disabled and does not listen for events or execute.
	Disabled *bool `field:"optional" json:"disabled" yaml:"disabled"`
	// An object where each field name is an event processor ID and each value is an object that configures its corresponding event processor.
	//
	// For an example configuration object, see
	// [Send Trigger Events to AWS
	// EventBridge](https://www.mongodb.com/docs/realm/triggers/examples/send-events-aws-eventbridge#std-label-event_processor_example).
	EventProcessors *Event `field:"optional" json:"eventProcessors" yaml:"eventProcessors"`
	// The ID of the function that the trigger calls when it fires.
	//
	// This value is the same as `event_processors.FUNCTION.function_id`.
	// You can either define the value here or in `event_processors.FUNCTION.function_id`.
	// The App Services backend duplicates the value to the configuration location where you did not define it.
	//
	// For example, if you define `function_id`, the backend duplicates it to `event_processors.FUNCTION.function_id`.
	FunctionId *string `field:"optional" json:"functionId" yaml:"functionId"`
	// The name of the function that the trigger calls when it fires, i.e. the function described by `function_id`.
	//
	// This value is the same as `event_processors.FUNCTION.function_name`.
	// You can either define the value here or in `event_processors.FUNCTION.function_name`.
	// The App Services backend duplicates the value to the configuration location where you did not define it.
	//
	// For example, if you define `function_name`, the backend duplicates it to `event_processors.FUNCTION.function_name`.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// The trigger's name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	ScheduleTrigger *ScheduleConfig `field:"optional" json:"scheduleTrigger" yaml:"scheduleTrigger"`
	// The trigger's type.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

