package awscdkresourcesmongodbatlas


// Label that identifies the cloud service provider where MongoDB Cloud performs stream processing.
//
// Currently, this parameter supports AWS only.
type StreamsDataProcessRegionCloudProvider string

const (
	// AWS.
	StreamsDataProcessRegionCloudProvider_AWS StreamsDataProcessRegionCloudProvider = "AWS"
	// GCP.
	StreamsDataProcessRegionCloudProvider_GCP StreamsDataProcessRegionCloudProvider = "GCP"
	// AZURE.
	StreamsDataProcessRegionCloudProvider_AZURE StreamsDataProcessRegionCloudProvider = "AZURE"
	// TENANT.
	StreamsDataProcessRegionCloudProvider_TENANT StreamsDataProcessRegionCloudProvider = "TENANT"
	// SERVERLESS.
	StreamsDataProcessRegionCloudProvider_SERVERLESS StreamsDataProcessRegionCloudProvider = "SERVERLESS"
)

