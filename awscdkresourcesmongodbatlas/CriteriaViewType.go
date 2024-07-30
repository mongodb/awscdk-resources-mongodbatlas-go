package awscdkresourcesmongodbatlas


// Means by which MongoDB Cloud selects data to archive.
//
// Data can be chosen using the age of the data or a MongoDB query. **DATE** selects documents to archive based on a date. (if DATE is selected, the PartitionFields.FieldName must be completed with the Criteria.DateField value)**CUSTOM** selects documents to archive based on a custom JSON query (When selecting this option, the Query property must be inputted). MongoDB Cloud doesn't support **CUSTOM** when `collectionType: TIMESERIES`.
type CriteriaViewType string

const (
	// DATE.
	CriteriaViewType_DATE CriteriaViewType = "DATE"
	// CUSTOM.
	CriteriaViewType_CUSTOM CriteriaViewType = "CUSTOM"
)

