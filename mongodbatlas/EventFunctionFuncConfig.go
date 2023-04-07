// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type EventFunctionFuncConfig struct {
	// The ID of the function that the trigger calls when it fires.
	//
	// This value is the same as the root-level `function_id`.
	// You can either define the value here or in `function_id`.
	// The App Services backend duplicates the value to the configuration location where you did not define it.
	//
	// For example, if you define `event_processors.FUNCTION.function_id`, the backend duplicates it to `function_id`.
	FunctionId *string `field:"optional" json:"functionId" yaml:"functionId"`
	// The name of the function that the trigger calls when it fires, i.e. the function described by `function_id`.
	//
	// This value is the same as the root-level `function_name`.
	// You can either define the value here or in `function_name`.
	// The App Services backend duplicates the value to the configuration location where you did not define it.
	//
	// For example, if you define `event_processors.FUNCTION.function_name`, the backend duplicates it to `function_name`.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
}

