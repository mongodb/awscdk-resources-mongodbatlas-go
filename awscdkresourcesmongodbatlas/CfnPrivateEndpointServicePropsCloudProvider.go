package awscdkresourcesmongodbatlas


// Cloud service provider that manages this private endpoint, default : AWS.
type CfnPrivateEndpointServicePropsCloudProvider string

const (
	// AWS.
	CfnPrivateEndpointServicePropsCloudProvider_AWS CfnPrivateEndpointServicePropsCloudProvider = "AWS"
	// AZURE.
	CfnPrivateEndpointServicePropsCloudProvider_AZURE CfnPrivateEndpointServicePropsCloudProvider = "AZURE"
	// GCP.
	CfnPrivateEndpointServicePropsCloudProvider_GCP CfnPrivateEndpointServicePropsCloudProvider = "GCP"
)

