// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


// Syntax used to write the date after which data moves to the online archive.
//
// Date can be expressed as ISO 8601 or Epoch timestamps. The Epoch timestamp can be expressed as nanoseconds, milliseconds, or seconds. Set this parameter when "criteria.type" : "DATE". You must set "criteria.type" : "DATE" if "collectionType": "TIMESERIES".
type CriteriaViewDateFormat string

const (
	// ISODATE.
	CriteriaViewDateFormat_ISODATE CriteriaViewDateFormat = "ISODATE"
	// EPOCH_SECONDS.
	CriteriaViewDateFormat_EPOCH_SECONDS CriteriaViewDateFormat = "EPOCH_SECONDS"
	// EPOCH_MILLIS.
	CriteriaViewDateFormat_EPOCH_MILLIS CriteriaViewDateFormat = "EPOCH_MILLIS"
	// EPOCH_NANOSECONDS.
	CriteriaViewDateFormat_EPOCH_NANOSECONDS CriteriaViewDateFormat = "EPOCH_NANOSECONDS"
)

