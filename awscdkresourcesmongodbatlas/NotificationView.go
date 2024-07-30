package awscdkresourcesmongodbatlas


type NotificationView struct {
	// Slack API token or Bot token that MongoDB Cloud needs to send alert notifications via Slack.
	//
	// The resource requires this parameter when '"notifications.typeName" : "SLACK"'. If the token later becomes invalid, MongoDB Cloud sends an email to the project owners. If the token remains invalid, MongoDB Cloud removes the token.
	ApiToken *string `field:"optional" json:"apiToken" yaml:"apiToken"`
	// Name of the Slack channel to which MongoDB Cloud sends alert notifications.
	//
	// The resource requires this parameter when '"notifications.typeName" : "SLACK"'.
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// Datadog API Key that MongoDB Cloud needs to send alert notifications to Datadog.
	//
	// You can find this API key in the Datadog dashboard. The resource requires this parameter when '"notifications.typeName" : "DATADOG"'.
	DatadogApiKey *string `field:"optional" json:"datadogApiKey" yaml:"datadogApiKey"`
	// Datadog region that indicates which API Uniform Resource Locator (URL) to use.
	//
	// The resource requires this parameter when '"notifications.typeName" : "DATADOG"'.
	DatadogRegion NotificationViewDatadogRegion `field:"optional" json:"datadogRegion" yaml:"datadogRegion"`
	// Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification.
	DelayMin *float64 `field:"optional" json:"delayMin" yaml:"delayMin"`
	// Email address to which MongoDB Cloud sends alert notifications.
	//
	// The resource requires this parameter when '"notifications.typeName" : "EMAIL"'. You don't need to set this value to send emails to individual or groups of MongoDB Cloud users including:
	//
	// - specific MongoDB Cloud users ('"notifications.typeName" : "USER"')
	// - MongoDB Cloud users with specific project roles ('"notifications.typeName" : "GROUP"')
	// - MongoDB Cloud users with specific organization roles ('"notifications.typeName" : "ORG"')
	// - MongoDB Cloud teams ('"notifications.typeName" : "TEAM"')
	//
	// To send emails to one MongoDB Cloud user or grouping of users, set the **notifications.emailEnabled** parameter.
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// Flag that indicates whether MongoDB Cloud should send email notifications.
	//
	// The resource requires this parameter when one of the following values have been set:
	//
	// - '"notifications.typeName" : "ORG"'
	// - '"notifications.typeName" : "GROUP"'
	// - '"notifications.typeName" : "USER"'
	EmailEnabled *bool `field:"optional" json:"emailEnabled" yaml:"emailEnabled"`
	// Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.
	//
	// PagerDuty, VictorOps, and OpsGenie notifications don't return this element. Configure and manage the notification interval within each of those services.
	IntervalMin *float64 `field:"optional" json:"intervalMin" yaml:"intervalMin"`
	// Microsoft Teams Webhook Uniform Resource Locator (URL) that MongoDB Cloud needs to send this notification via Microsoft Teams.
	//
	// The resource requires this parameter when '"notifications.typeName" : "MICROSOFT_TEAMS"'. If the URL later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.
	MicrosoftTeamsWebhookUrl *string `field:"optional" json:"microsoftTeamsWebhookUrl" yaml:"microsoftTeamsWebhookUrl"`
	// Mobile phone number to which MongoDB Cloud sends alert notifications.
	//
	// The resource requires this parameter when '"notifications.typeName" : "SMS"'.
	MobileNumber *string `field:"optional" json:"mobileNumber" yaml:"mobileNumber"`
	// HipChat API token that MongoDB Cloud needs to send alert notifications to HipChat.
	//
	// The resource requires this parameter when '"notifications.typeName" : "HIP_CHAT"'". If the token later becomes invalid, MongoDB Cloud sends an email to the project owners. If the token remains invalid, MongoDB Cloud removes it.
	NotificationToken *string `field:"optional" json:"notificationToken" yaml:"notificationToken"`
	// API Key that MongoDB Cloud needs to send this notification via Opsgenie.
	//
	// The resource requires this parameter when '"notifications.typeName" : "OPS_GENIE"'. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.
	OpsGenieApiKey *string `field:"optional" json:"opsGenieApiKey" yaml:"opsGenieApiKey"`
	// Opsgenie region that indicates which API Uniform Resource Locator (URL) to use.
	OpsGenieRegion NotificationViewOpsGenieRegion `field:"optional" json:"opsGenieRegion" yaml:"opsGenieRegion"`
	// Flowdock organization name to which MongoDB Cloud sends alert notifications.
	//
	// This name appears after 'www.flowdock.com/app/' in the Uniform Resource Locator (URL) path. The resource requires this parameter when '"notifications.typeName" : "FLOWDOCK"'.
	OrgName *string `field:"optional" json:"orgName" yaml:"orgName"`
	// List that contains the one or more organization or project roles that receive the configured alert.
	//
	// The resource requires this parameter when '"notifications.typeName" : "GROUP"' or '"notifications.typeName" : "ORG"'. If you include this parameter, MongoDB Cloud sends alerts only to users assigned the roles you specify in the array. If you omit this parameter, MongoDB Cloud sends alerts to users assigned any role.
	Roles *[]NotificationViewRoles `field:"optional" json:"roles" yaml:"roles"`
	// HipChat API room name to which MongoDB Cloud sends alert notifications.
	//
	// The resource requires this parameter when '"notifications.typeName" : "HIP_CHAT"'".
	RoomName *string `field:"optional" json:"roomName" yaml:"roomName"`
	// PagerDuty service key that MongoDB Cloud needs to send notifications via PagerDuty.
	//
	// The resource requires this parameter when '"notifications.typeName" : "PAGER_DUTY"'. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.
	ServiceKey *string `field:"optional" json:"serviceKey" yaml:"serviceKey"`
	// Degree of seriousness given to this notification.
	Severity NotificationViewSeverity `field:"optional" json:"severity" yaml:"severity"`
	// Flag that indicates whether MongoDB Cloud should send text message notifications.
	//
	// The resource requires this parameter when one of the following values have been set:
	//
	// - '"notifications.typeName" : "ORG"'
	// - '"notifications.typeName" : "GROUP"'
	// - '"notifications.typeName" : "USER"'
	SmsEnabled *bool `field:"optional" json:"smsEnabled" yaml:"smsEnabled"`
	// Unique 24-hexadecimal digit string that identifies one MongoDB Cloud team.
	//
	// The resource requires this parameter when '"notifications.typeName" : "TEAM"'.
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
	// Name of the MongoDB Cloud team that receives this notification.
	//
	// The resource requires this parameter when '"notifications.typeName" : "TEAM"'.
	TeamName *string `field:"optional" json:"teamName" yaml:"teamName"`
	// Human-readable label that displays the alert notification type.
	TypeName NotificationViewTypeName `field:"optional" json:"typeName" yaml:"typeName"`
	// MongoDB Cloud username of the person to whom MongoDB Cloud sends notifications.
	//
	// Specify only MongoDB Cloud users who belong to the project that owns the alert configuration. The resource requires this parameter when '"notifications.typeName" : "USER"'.
	Username *string `field:"optional" json:"username" yaml:"username"`
	// API key that MongoDB Cloud needs to send alert notifications to Splunk On-Call.
	//
	// The resource requires this parameter when '"notifications.typeName" : "VICTOR_OPS"'. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.
	VictorOpsApiKey *string `field:"optional" json:"victorOpsApiKey" yaml:"victorOpsApiKey"`
	// Routing key that MongoDB Cloud needs to send alert notifications to Splunk On-Call.
	//
	// The resource requires this parameter when '"notifications.typeName" : "VICTOR_OPS"'. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.
	VictorOpsRoutingKey *string `field:"optional" json:"victorOpsRoutingKey" yaml:"victorOpsRoutingKey"`
	// An optional field for your webhook secret.
	WebhookSecret *string `field:"optional" json:"webhookSecret" yaml:"webhookSecret"`
	// Your webhook URL.
	WebhookUrl *string `field:"optional" json:"webhookUrl" yaml:"webhookUrl"`
}

