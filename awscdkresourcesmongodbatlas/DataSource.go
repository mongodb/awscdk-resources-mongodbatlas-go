package awscdkresourcesmongodbatlas


type DataSource struct {
	// Flag that validates the scheme in the specified URLs.
	//
	// If true, allows insecure HTTP scheme, doesn't verify the server's certificate chain and hostname, and accepts any certificate with any hostname presented by the server. If false, allows secure HTTPS scheme only.
	AllowInsecure *bool `field:"optional" json:"allowInsecure" yaml:"allowInsecure"`
	// Human-readable label that identifies the collection in the database.
	//
	// For creating a wildcard (*) collection, you must omit this parameter.
	Collection *string `field:"optional" json:"collection" yaml:"collection"`
	// Regex pattern to use for creating the wildcard (*) collection.
	//
	// To learn more about the regex syntax, see Go programming language.( https://pkg.go.dev/regexp ).
	CollectionRegex *string `field:"optional" json:"collectionRegex" yaml:"collectionRegex"`
	// Human-readable label that identifies the database, which contains the collection in the cluster.
	//
	// You must omit this parameter to generate wildcard (*) collections for dynamically generated databases.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Regex pattern to use for creating the wildcard (*) collection.
	//
	// To learn more about the regex syntax, see Go programming language.( https://pkg.go.dev/regexp ).
	DatabaseRegex *string `field:"optional" json:"databaseRegex" yaml:"databaseRegex"`
	// File format that MongoDB Cloud uses if it encounters a file without a file extension while searching storeName.Enum: ".avro" ".avro.bz2" ".avro.gz" ".bson" ".bson.bz2" ".bson.gz" ".bsonx" ".csv" ".csv.bz2" ".csv.gz" ".json" ".json.bz2" ".json.gz" ".orc" ".parquet" ".tsv" ".tsv.bz2" ".tsv.gz".
	DefaultFormat *string `field:"optional" json:"defaultFormat" yaml:"defaultFormat"`
	// File path that controls how MongoDB Cloud searches for and parses files in the storeName before mapping them to a collection.Specify / to capture all files and folders from the prefix path.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Name for the field that includes the provenance of the documents in the results.
	//
	// MongoDB Cloud returns different fields in the results for each supported provider.
	ProvenanceFieldName *string `field:"optional" json:"provenanceFieldName" yaml:"provenanceFieldName"`
	// Human-readable label that identifies the data store that MongoDB Cloud maps to the collection.
	StoreName *string `field:"optional" json:"storeName" yaml:"storeName"`
	// URLs of the publicly accessible data files.
	//
	// You can't specify URLs that require authentication. Atlas Data Lake creates a partition for each URL. If empty or omitted, Data Lake uses the URLs from the store specified in the dataSources.storeName parameter.
	Urls *[]*string `field:"optional" json:"urls" yaml:"urls"`
}

