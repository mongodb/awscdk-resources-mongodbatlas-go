// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Returns, adds, edits, and removes third-party service integration configurations.
//
// MongoDB Cloud sends alerts to each third-party service that you configure.
type CfnThirdPartyIntegrationProps struct {
	// Key that allows MongoDB Cloud to access your Opsgenie/Datadog account.
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Key that allows MongoDB Cloud to access your Slack account.
	ApiToken *string `field:"optional" json:"apiToken" yaml:"apiToken"`
	// Name of the Slack channel to which MongoDB Cloud sends alert notifications.
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// Flag that indicates whether someone has activated the Prometheus integration.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Combination of IPv4 address and Internet Assigned Numbers Authority (IANA) port or the IANA port alone to which Prometheus binds to ingest MongoDB metrics.
	ListenAddress *string `field:"optional" json:"listenAddress" yaml:"listenAddress"`
	// Endpoint web address of the Microsoft Teams webhook to which MongoDB Cloud sends notifications.
	MicrosoftTeamsWebhookUrl *string `field:"optional" json:"microsoftTeamsWebhookUrl" yaml:"microsoftTeamsWebhookUrl"`
	// Password required for your integration with Prometheus.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Two-letter code that indicates which regional URL MongoDB uses to access the Opsgenie/Datadog API.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Routing key associated with your Splunk On-Call account.
	RoutingKey *string `field:"optional" json:"routingKey" yaml:"routingKey"`
	// Security Scheme to apply to HyperText Transfer Protocol (HTTP) traffic between Prometheus and MongoDB Cloud.
	Scheme CfnThirdPartyIntegrationPropsScheme `field:"optional" json:"scheme" yaml:"scheme"`
	// Parameter returned if someone configure this webhook with a secret.
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
	// Desired method to discover the Prometheus service.
	ServiceDiscovery CfnThirdPartyIntegrationPropsServiceDiscovery `field:"optional" json:"serviceDiscovery" yaml:"serviceDiscovery"`
	// Service key associated with your PagerDuty account.
	ServiceKey *string `field:"optional" json:"serviceKey" yaml:"serviceKey"`
	// Human-readable label that identifies your Slack team.
	//
	// Set this parameter when you configure a legacy Slack integration.
	TeamName *string `field:"optional" json:"teamName" yaml:"teamName"`
	// Root-relative path to the Transport Layer Security (TLS) Privacy Enhanced Mail (PEM) key and certificate file on the host.
	TlsPemPath *string `field:"optional" json:"tlsPemPath" yaml:"tlsPemPath"`
	// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud.
	//
	// The value must match the third-party service integration type.
	Type CfnThirdPartyIntegrationPropsType `field:"optional" json:"type" yaml:"type"`
	// Endpoint web address to which MongoDB Cloud sends notifications.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Human-readable label that identifies your Prometheus incoming webhook.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

