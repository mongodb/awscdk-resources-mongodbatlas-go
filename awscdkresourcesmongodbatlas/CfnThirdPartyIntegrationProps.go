package awscdkresourcesmongodbatlas


// Returns, adds, edits, and removes third-party service integration configurations.
//
// MongoDB Cloud sends alerts to each third-party service that you configure.
type CfnThirdPartyIntegrationProps struct {
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud.
	//
	// The value must match the third-party service integration type.
	Type CfnThirdPartyIntegrationPropsType `field:"required" json:"type" yaml:"type"`
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
	// Two-letter code that indicates which regional URL MongoDB uses to access the Opsgenie/Datadog API.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Routing key associated with your Splunk On-Call account.
	RoutingKey *string `field:"optional" json:"routingKey" yaml:"routingKey"`
	// Parameter returned if someone configure this webhook with a secret.
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
	// Flag that indicates whether to send collection latency metrics to Datadog, including database names, collection names, and latency metrics on reads, writes, commands, and transactions.
	SendCollectionLatencyMetrics *bool `field:"optional" json:"sendCollectionLatencyMetrics" yaml:"sendCollectionLatencyMetrics"`
	// Flag that indicates whether to send database metrics to Datadog, including database names and metrics on the number of collections, storage size, and index size.
	SendDatabaseMetrics *bool `field:"optional" json:"sendDatabaseMetrics" yaml:"sendDatabaseMetrics"`
	// Flag that indicates whether to include user-defined resource tags when sending metrics and alerts to third-party services.
	SendUserProvidedResourceTags *bool `field:"optional" json:"sendUserProvidedResourceTags" yaml:"sendUserProvidedResourceTags"`
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
	// Endpoint web address to which MongoDB Cloud sends notifications.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Human-readable label that identifies your Prometheus incoming webhook.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

