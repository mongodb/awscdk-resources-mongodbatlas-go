// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


// Returns, adds, edits, or removes an online archive.
type CfnOnlineArchiveProps struct {
	// Human-readable label that identifies the cluster that contains the collection from which you want to remove an online archive.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Rules by which MongoDB MongoDB Cloud archives data.
	//
	// Use the **criteria.type** field to choose how MongoDB Cloud selects data to archive. Choose data using the age of the data or a MongoDB query.
	// **"criteria.type": "DATE"** selects documents to archive based on a date.
	// **"criteria.type": "CUSTOM"** selects documents to archive based on a custom JSON query. MongoDB Cloud doesn't support **"criteria.type": "CUSTOM"** when **"collectionType": "TIMESERIES"**.
	Criteria *CriteriaView `field:"required" json:"criteria" yaml:"criteria"`
	// Classification of MongoDB database collection that you want to return.
	//
	// If you set this parameter to `TIMESERIES`, set `"criteria.type" : "date"` and `"criteria.dateFormat" : "ISODATE"`.
	CollectionType CfnOnlineArchivePropsCollectionType `field:"optional" json:"collectionType" yaml:"collectionType"`
	// Human-readable label that identifies the collection for which you created the online archive.
	CollName *string `field:"optional" json:"collName" yaml:"collName"`
	// Human-readable label of the database that contains the collection that contains the online archive.
	DbName *string `field:"optional" json:"dbName" yaml:"dbName"`
	// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
	IncludeCount *bool `field:"optional" json:"includeCount" yaml:"includeCount"`
	// Number of items that the response returns per page.
	ItemsPerPage *float64 `field:"optional" json:"itemsPerPage" yaml:"itemsPerPage"`
	// Number of the page that displays the current set of the total objects that the response returns.
	PageNum *float64 `field:"optional" json:"pageNum" yaml:"pageNum"`
	// List that contains document parameters to use to logically divide data within a collection.
	//
	// Partitions provide a coarse level of filtering of the underlying collection data. To divide your data, specify up to two parameters that you frequently query. Any queries that don't use these parameters result in a full collection scan of all archived documents. This takes more time and increase your costs.
	PartitionFields *[]*PartitionFieldView `field:"optional" json:"partitionFields" yaml:"partitionFields"`
	// The profile is defined in AWS Secret manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Regular frequency and duration when archiving process occurs.
	Schedule *ScheduleView `field:"optional" json:"schedule" yaml:"schedule"`
}

