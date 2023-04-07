// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Means by which MongoDB Cloud selects data to archive.
//
// Data can be chosen using the age of the data or a MongoDB query.
// **DATE** selects documents to archive based on a date.
// **CUSTOM** selects documents to archive based on a custom JSON query. MongoDB Cloud doesn't support **CUSTOM** when `"collectionType": "TIMESERIES"`.
type CriteriaViewType string

const (
	// DATE.
	CriteriaViewType_DATE CriteriaViewType = "DATE"
	// CUSTOM.
	CriteriaViewType_CUSTOM CriteriaViewType = "CUSTOM"
)

