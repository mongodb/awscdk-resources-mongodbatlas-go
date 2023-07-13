package awscdkresourcesmongodbatlas


// Returns and edits the conditions that trigger alerts and how MongoDB Cloud notifies users.
//
// This collection remains under revision and may change. Refer to the legacy documentation for this collection in the following link.
type CfnAlertConfigurationProps struct {
	// Date and time when MongoDB Cloud created the alert configuration.
	//
	// This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Created *string `field:"optional" json:"created" yaml:"created"`
	// Event type that triggers an alert.
	EventTypeName *string `field:"optional" json:"eventTypeName" yaml:"eventTypeName"`
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
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Limit that triggers an alert when exceeded.
	//
	// The resource returns this parameter when **eventTypeName** has not been set to 'OUTSIDE_METRIC_THRESHOLD'.
	Threshold *IntegerThresholdView `field:"optional" json:"threshold" yaml:"threshold"`
	// Human-readable label that displays the alert type.
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
}

