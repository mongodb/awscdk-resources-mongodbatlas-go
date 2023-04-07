// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Classification of MongoDB database collection that you want to return.
//
// If you set this parameter to `TIMESERIES`, set `"criteria.type" : "date"` and `"criteria.dateFormat" : "ISODATE"`.
type CfnOnlineArchivePropsCollectionType string

const (
	// STANDARD.
	CfnOnlineArchivePropsCollectionType_STANDARD CfnOnlineArchivePropsCollectionType = "STANDARD"
	// TIMESERIES.
	CfnOnlineArchivePropsCollectionType_TIMESERIES CfnOnlineArchivePropsCollectionType = "TIMESERIES"
)

