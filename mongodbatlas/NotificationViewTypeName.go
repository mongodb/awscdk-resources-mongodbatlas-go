// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Human-readable label that displays the alert notification type.
type NotificationViewTypeName string

const (
	// DATADOG.
	NotificationViewTypeName_DATADOG NotificationViewTypeName = "DATADOG"
	// EMAIL.
	NotificationViewTypeName_EMAIL NotificationViewTypeName = "EMAIL"
	// FLOWDOCK.
	NotificationViewTypeName_FLOWDOCK NotificationViewTypeName = "FLOWDOCK"
	// GROUP.
	NotificationViewTypeName_GROUP NotificationViewTypeName = "GROUP"
	// MICROSOFT_TEAMS.
	NotificationViewTypeName_MICROSOFT_TEAMS NotificationViewTypeName = "MICROSOFT_TEAMS"
	// OPS_GENIE.
	NotificationViewTypeName_OPS_GENIE NotificationViewTypeName = "OPS_GENIE"
	// ORG.
	NotificationViewTypeName_ORG NotificationViewTypeName = "ORG"
	// PAGER_DUTY.
	NotificationViewTypeName_PAGER_DUTY NotificationViewTypeName = "PAGER_DUTY"
	// PROMETHEUS.
	NotificationViewTypeName_PROMETHEUS NotificationViewTypeName = "PROMETHEUS"
	// SLACK.
	NotificationViewTypeName_SLACK NotificationViewTypeName = "SLACK"
	// SMS.
	NotificationViewTypeName_SMS NotificationViewTypeName = "SMS"
	// TEAM.
	NotificationViewTypeName_TEAM NotificationViewTypeName = "TEAM"
	// USER.
	NotificationViewTypeName_USER NotificationViewTypeName = "USER"
	// VICTOR_OPS.
	NotificationViewTypeName_VICTOR_OPS NotificationViewTypeName = "VICTOR_OPS"
	// WEBHOOK.
	NotificationViewTypeName_WEBHOOK NotificationViewTypeName = "WEBHOOK"
)

