package awscdkresourcesmongodbatlas


// Options that needs to be set to control the synchronous creation flow, this options need to be set if EnableSynchronousCreation is se to TRUE.
type SynchronousCreationOptions struct {
	// Represents the time interval, measured in seconds, for the synchronous process to wait before checking again to verify if the job has been completed.
	//
	// example: if set to 20, it will chek every 20 seconds if the resource is completed, default (30 seconds).
	CallbackDelaySeconds *float64 `field:"optional" json:"callbackDelaySeconds" yaml:"callbackDelaySeconds"`
	// if set to true, the process will return success, in the event of a timeOut, default false.
	ReturnSuccessIfTimeOut *bool `field:"optional" json:"returnSuccessIfTimeOut" yaml:"returnSuccessIfTimeOut"`
	// The amount of time the process will wait until exiting with a success, default (1200 seconds).
	TimeOutInSeconds *float64 `field:"optional" json:"timeOutInSeconds" yaml:"timeOutInSeconds"`
}

