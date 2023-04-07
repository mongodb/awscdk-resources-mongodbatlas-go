// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type DataLakeDataProcessRegionView struct {
	// Name of the cloud service that hosts the data lake's data stores.
	CloudProvider DataLakeDataProcessRegionViewCloudProvider `field:"optional" json:"cloudProvider" yaml:"cloudProvider"`
	// Name of the region to which the data lake routes client connections.
	Region DataLakeDataProcessRegionViewRegion `field:"optional" json:"region" yaml:"region"`
}

