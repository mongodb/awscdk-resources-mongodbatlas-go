package awscdkresourcesmongodbatlas


type CriteriaView struct {
	// Indexed database parameter that stores the date that determines when data moves to the online archive.
	//
	// MongoDB Cloud archives the data when the current date exceeds the date in this database parameter plus the number of days specified through the expireAfterDays parameter. Set this parameter when you set "criteria.type" : "DATE".
	DateField *string `field:"optional" json:"dateField" yaml:"dateField"`
	// Syntax used to write the date after which data moves to the online archive.
	//
	// Date can be expressed as ISO 8601 or Epoch timestamps. The Epoch timestamp can be expressed as nanoseconds, milliseconds, or seconds. Set this parameter when "criteria.type" : "DATE". You must set "criteria.type" : "DATE" if "collectionType": "TIMESERIES".
	DateFormat CriteriaViewDateFormat `field:"optional" json:"dateFormat" yaml:"dateFormat"`
	// Number of days after the value in the criteria.dateField when MongoDB Cloud archives data in the specified cluster. Set this parameter when you set "criteria.type" : "DATE".
	ExpireAfterDays *float64 `field:"optional" json:"expireAfterDays" yaml:"expireAfterDays"`
	// MongoDB find query that selects documents to archive.
	//
	// The specified query follows the syntax of the db.collection.find(query) command. This query can't use the empty document ({}) to return all documents. Set this parameter when "criteria.type" : "CUSTOM".
	Query *string `field:"optional" json:"query" yaml:"query"`
	// Means by which MongoDB Cloud selects data to archive.
	//
	// Data can be chosen using the age of the data or a MongoDB query.
	// **DATE** selects documents to archive based on a date.
	// **CUSTOM** selects documents to archive based on a custom JSON query. MongoDB Cloud doesn't support **CUSTOM** when `"collectionType": "TIMESERIES"`.
	Type CriteriaViewType `field:"optional" json:"type" yaml:"type"`
}

