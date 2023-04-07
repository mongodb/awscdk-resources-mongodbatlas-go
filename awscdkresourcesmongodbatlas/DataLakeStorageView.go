// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type DataLakeStorageView struct {
	// Array that contains the queryable databases and collections for this data lake.
	Databases *[]*DataLakeDatabaseView `field:"optional" json:"databases" yaml:"databases"`
	// Array that contains the data stores for the data lake.
	Stores *[]*StoreDetail `field:"optional" json:"stores" yaml:"stores"`
}

