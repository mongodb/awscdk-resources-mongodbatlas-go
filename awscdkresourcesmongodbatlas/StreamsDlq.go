package awscdkresourcesmongodbatlas


// Dead letter queue for the stream processor.
//
// Refer to the MongoDB Atlas Docs for more information.
type StreamsDlq struct {
	// Name of the collection to use for the DLQ.
	Coll *string `field:"required" json:"coll" yaml:"coll"`
	// Name of the connection to write DLQ messages to.
	//
	// Must be an Atlas connection.
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
	// Name of the database to use for the DLQ.
	Db *string `field:"required" json:"db" yaml:"db"`
}

