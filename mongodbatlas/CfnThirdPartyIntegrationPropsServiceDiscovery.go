// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Desired method to discover the Prometheus service.
type CfnThirdPartyIntegrationPropsServiceDiscovery string

const (
	// http.
	CfnThirdPartyIntegrationPropsServiceDiscovery_HTTP CfnThirdPartyIntegrationPropsServiceDiscovery = "HTTP"
	// file.
	CfnThirdPartyIntegrationPropsServiceDiscovery_FILE CfnThirdPartyIntegrationPropsServiceDiscovery = "FILE"
)

