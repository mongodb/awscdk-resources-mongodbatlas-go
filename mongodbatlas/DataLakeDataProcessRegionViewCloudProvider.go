// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Name of the cloud service that hosts the data lake's data stores.
type DataLakeDataProcessRegionViewCloudProvider string

const (
	// AWS.
	DataLakeDataProcessRegionViewCloudProvider_AWS DataLakeDataProcessRegionViewCloudProvider = "AWS"
	// GCP.
	DataLakeDataProcessRegionViewCloudProvider_GCP DataLakeDataProcessRegionViewCloudProvider = "GCP"
	// AZURE.
	DataLakeDataProcessRegionViewCloudProvider_AZURE DataLakeDataProcessRegionViewCloudProvider = "AZURE"
	// TENANT.
	DataLakeDataProcessRegionViewCloudProvider_TENANT DataLakeDataProcessRegionViewCloudProvider = "TENANT"
	// SERVERLESS.
	DataLakeDataProcessRegionViewCloudProvider_SERVERLESS DataLakeDataProcessRegionViewCloudProvider = "SERVERLESS"
)

