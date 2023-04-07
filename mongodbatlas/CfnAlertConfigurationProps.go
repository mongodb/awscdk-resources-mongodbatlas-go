// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Returns and edits the conditions that trigger alerts and how MongoDB Cloud notifies users.
//
// This collection remains under revision and may change. Refer to the legacy documentation for this collection in the following link.
type CfnAlertConfigurationProps struct {
	// Event type that triggers an alert.
	EventTypeName CfnAlertConfigurationPropsEventTypeName `field:"optional" json:"eventTypeName" yaml:"eventTypeName"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both.
	//
	// RFC 5988 outlines these relationships.
	Links *[]*Link `field:"optional" json:"links" yaml:"links"`
	// List of rules that determine whether MongoDB Cloud checks an object for the alert configuration.
	//
	// You can filter using the matchers array if the **eventTypeName** specifies an event for a host, replica set, or sharded cluster.
	Matchers *[]*Matcher `field:"optional" json:"matchers" yaml:"matchers"`
	// Threshold for the metric that, when exceeded, triggers an alert.
	//
	// The resource returns this parameter when '"eventTypeName" : "OUTSIDE_METRIC_THRESHOLD"'.
	MetricThreshold *MetricThresholdView `field:"optional" json:"metricThreshold" yaml:"metricThreshold"`
	// List that contains the targets that MongoDB Cloud sends notifications.
	Notifications *[]*NotificationView `field:"optional" json:"notifications" yaml:"notifications"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// List of returned documents that MongoDB Cloud provides when completing this request.
	Results *[]*AlertView `field:"optional" json:"results" yaml:"results"`
	// Limit that triggers an alert when exceeded.
	//
	// The resource returns this parameter when **eventTypeName** has not been set to 'OUTSIDE_METRIC_THRESHOLD'.
	Threshold *IntegerThresholdView `field:"optional" json:"threshold" yaml:"threshold"`
}

