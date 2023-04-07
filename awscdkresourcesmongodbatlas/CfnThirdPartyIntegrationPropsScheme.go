// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


// Security Scheme to apply to HyperText Transfer Protocol (HTTP) traffic between Prometheus and MongoDB Cloud.
type CfnThirdPartyIntegrationPropsScheme string

const (
	// http.
	CfnThirdPartyIntegrationPropsScheme_HTTP CfnThirdPartyIntegrationPropsScheme = "HTTP"
	// https.
	CfnThirdPartyIntegrationPropsScheme_HTTPS CfnThirdPartyIntegrationPropsScheme = "HTTPS"
)

