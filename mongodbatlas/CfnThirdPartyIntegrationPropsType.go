// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud.
//
// The value must match the third-party service integration type.
type CfnThirdPartyIntegrationPropsType string

const (
	// PAGER_DUTY.
	CfnThirdPartyIntegrationPropsType_PAGER_DUTY CfnThirdPartyIntegrationPropsType = "PAGER_DUTY"
	// MICROSOFT_TEAMS.
	CfnThirdPartyIntegrationPropsType_MICROSOFT_TEAMS CfnThirdPartyIntegrationPropsType = "MICROSOFT_TEAMS"
	// SLACK.
	CfnThirdPartyIntegrationPropsType_SLACK CfnThirdPartyIntegrationPropsType = "SLACK"
	// DATADOG.
	CfnThirdPartyIntegrationPropsType_DATADOG CfnThirdPartyIntegrationPropsType = "DATADOG"
	// OPS_GENIE.
	CfnThirdPartyIntegrationPropsType_OPS_GENIE CfnThirdPartyIntegrationPropsType = "OPS_GENIE"
	// VICTOR_OPS.
	CfnThirdPartyIntegrationPropsType_VICTOR_OPS CfnThirdPartyIntegrationPropsType = "VICTOR_OPS"
	// WEBHOOK.
	CfnThirdPartyIntegrationPropsType_WEBHOOK CfnThirdPartyIntegrationPropsType = "WEBHOOK"
	// PROMETHEUS.
	CfnThirdPartyIntegrationPropsType_PROMETHEUS CfnThirdPartyIntegrationPropsType = "PROMETHEUS"
)

