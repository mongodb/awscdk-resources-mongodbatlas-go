// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type AlertView struct {
	// Date and time until which this alert has been acknowledged.
	//
	// This parameter expresses its value in the ISO 8601 timestamp format in UTC. The resource returns this parameter if a MongoDB User previously acknowledged this alert.
	//
	// - To acknowledge this alert forever, set the parameter value to 100 years in the future.
	//
	// - To unacknowledge a previously acknowledged alert, set the parameter value to a date in the past.
	AcknowledgedUntil *string `field:"optional" json:"acknowledgedUntil" yaml:"acknowledgedUntil"`
	// Comment that a MongoDB Cloud user submitted when acknowledging the alert.
	AcknowledgementComment *string `field:"optional" json:"acknowledgementComment" yaml:"acknowledgementComment"`
	// MongoDB Cloud username of the person who acknowledged the alert.
	//
	// The response returns this parameter if a MongoDB Cloud user previously acknowledged this alert.
	AcknowledgingUsername *string `field:"optional" json:"acknowledgingUsername" yaml:"acknowledgingUsername"`
	// Unique 24-hexadecimal digit string that identifies the alert configuration that sets this alert.
	AlertConfigId *string `field:"optional" json:"alertConfigId" yaml:"alertConfigId"`
	// Human-readable label that identifies the cluster to which this alert applies.
	//
	// This resource returns this parameter for alerts of events impacting backups, replica sets, or sharded clusters.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Date and time when MongoDB Cloud created this alert.
	//
	// This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Created *string `field:"optional" json:"created" yaml:"created"`
	// Value of the metric that triggered the alert.
	//
	// The resource returns this parameter for alerts of events impacting hosts.
	CurrentValue *CurrentValue `field:"optional" json:"currentValue" yaml:"currentValue"`
	// Incident that triggered this alert.
	EventTypeName AlertViewEventTypeName `field:"optional" json:"eventTypeName" yaml:"eventTypeName"`
	// Unique 24-hexadecimal digit string that identifies the project that owns this alert.
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Hostname and port of the host to which this alert applies.
	//
	// The resource returns this parameter for alerts of events impacting hosts or replica sets.
	HostnameAndPort *string `field:"optional" json:"hostnameAndPort" yaml:"hostnameAndPort"`
	// Unique 24-hexadecimal digit string that identifies this alert.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Date and time that any notifications were last sent for this alert.
	//
	// This parameter expresses its value in the ISO 8601 timestamp format in UTC. The resource returns this parameter if MongoDB Cloud has sent notifications for this alert.
	LastNotified *string `field:"optional" json:"lastNotified" yaml:"lastNotified"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both.
	//
	// RFC 5988 outlines these relationships.
	Links *[]*Link `field:"optional" json:"links" yaml:"links"`
	// Human-readable label that identifies the metric against which MongoDB Cloud checks the alert.
	MetricName AlertViewMetricName `field:"optional" json:"metricName" yaml:"metricName"`
	// Name of the replica set to which this alert applies.
	//
	// The response returns this parameter for alerts of events impacting backups, hosts, or replica sets.
	ReplicaSetName *string `field:"optional" json:"replicaSetName" yaml:"replicaSetName"`
	// Date and time that this alert changed to '"status" : "CLOSED"'.
	//
	// This parameter expresses its value in the ISO 8601 timestamp format in UTC. The resource returns this parameter once '"status" : "CLOSED"'.
	Resolved *string `field:"optional" json:"resolved" yaml:"resolved"`
	// State of this alert at the time you requested its details.
	Status AlertViewStatus `field:"optional" json:"status" yaml:"status"`
	// Category in which MongoDB Cloud classifies this alert.
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
	// Date and time when someone last updated this alert.
	//
	// This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Updated *string `field:"optional" json:"updated" yaml:"updated"`
}

