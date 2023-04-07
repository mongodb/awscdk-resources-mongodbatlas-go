// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// File format that MongoDB Cloud uses if it encounters a file without a file extension while searching **storeName**.
type DataLakeDatabaseDataSourceViewDefaultFormat string

const (
	// .avro.
	DataLakeDatabaseDataSourceViewDefaultFormat_AVRO DataLakeDatabaseDataSourceViewDefaultFormat = "AVRO"
	// .avro.gz.
	DataLakeDatabaseDataSourceViewDefaultFormat_AVRO_GZ DataLakeDatabaseDataSourceViewDefaultFormat = "AVRO_GZ"
	// .bson.
	DataLakeDatabaseDataSourceViewDefaultFormat_BSON DataLakeDatabaseDataSourceViewDefaultFormat = "BSON"
	// .bson.gz.
	DataLakeDatabaseDataSourceViewDefaultFormat_BSON_GZ DataLakeDatabaseDataSourceViewDefaultFormat = "BSON_GZ"
	// .csv.
	DataLakeDatabaseDataSourceViewDefaultFormat_CSV DataLakeDatabaseDataSourceViewDefaultFormat = "CSV"
	// .json.
	DataLakeDatabaseDataSourceViewDefaultFormat_JSON DataLakeDatabaseDataSourceViewDefaultFormat = "JSON"
	// .json.gz.
	DataLakeDatabaseDataSourceViewDefaultFormat_JSON_GZ DataLakeDatabaseDataSourceViewDefaultFormat = "JSON_GZ"
	// .orc.
	DataLakeDatabaseDataSourceViewDefaultFormat_ORC DataLakeDatabaseDataSourceViewDefaultFormat = "ORC"
	// .tsv.
	DataLakeDatabaseDataSourceViewDefaultFormat_TSV DataLakeDatabaseDataSourceViewDefaultFormat = "TSV"
	// .tsv.gz.
	DataLakeDatabaseDataSourceViewDefaultFormat_TSV_GZ DataLakeDatabaseDataSourceViewDefaultFormat = "TSV_GZ"
)

