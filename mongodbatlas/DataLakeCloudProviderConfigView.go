// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type DataLakeCloudProviderConfigView struct {
	// Name of the cloud service that hosts the data lake's data stores.
	Aws *DataLakeAwsCloudProviderConfigView `field:"optional" json:"aws" yaml:"aws"`
}

