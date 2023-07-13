package awscdkresourcesmongodbatlas


// Datadog region that indicates which API Uniform Resource Locator (URL) to use.
//
// The resource requires this parameter when '"notifications.typeName" : "DATADOG"'.
type NotificationViewDatadogRegion string

const (
	// EU.
	NotificationViewDatadogRegion_EU NotificationViewDatadogRegion = "EU"
	// US.
	NotificationViewDatadogRegion_US NotificationViewDatadogRegion = "US"
)

