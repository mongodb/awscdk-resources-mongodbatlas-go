package awscdkresourcesmongodbatlas


type DatabaseConfig struct {
	// The name of a collection in the specified database.
	//
	// The
	// trigger listens to events from this collection.
	Collection *string `field:"optional" json:"collection" yaml:"collection"`
	// The name of a database in the linked data source.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// If `true`, indicates that `UPDATE` change events should include the most current [majority-committed](https://www.mongodb.com/docs/manual/reference/read-concern-majority/) version of the modified document in the `fullDocument` field.
	FullDocument *bool `field:"optional" json:"fullDocument" yaml:"fullDocument"`
	// If true, indicates that `UPDATE` change events should include a snapshot of the modified document from immediately before the update was applied.
	//
	// You must enable [document
	// preimages](https://www.mongodb.com/docs/atlas/app-services/mongodb/preimages/)
	// for your cluster to include these snapshots.
	FullDocumentBeforeChange *bool `field:"optional" json:"fullDocumentBeforeChange" yaml:"fullDocumentBeforeChange"`
	// stringify version of a [$match](https://www.mongodb.com/docs/manual/reference/operator/aggregation/match) expression filters change events. The trigger will only fire if the expression evaluates to true for a given change event.
	Match *string `field:"optional" json:"match" yaml:"match"`
	// The type(s) of MongoDB change event that the trigger listens for.
	OperationTypes *[]DatabaseConfigOperationTypes `field:"optional" json:"operationTypes" yaml:"operationTypes"`
	// stringify version of a [$project](https://www.mongodb.com/docs/manual/reference/operator/aggregation/project/) expressions to limit the data included in each event.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The _id value of a linked MongoDB data source.
	//
	// See [Get a Data Source](#operation/adminGetService).
	ServiceId *string `field:"optional" json:"serviceId" yaml:"serviceId"`
	// If `true`, enabling the Trigger after it was disabled will not invoke events that occurred while the Trigger was disabled.
	SkipCatchupEvents *bool `field:"optional" json:"skipCatchupEvents" yaml:"skipCatchupEvents"`
	// If `true`, when this Trigger's resume token cannot be found in the cluster's oplog, the Trigger automatically resumes processing events at the next relevant change stream event.
	//
	// All change stream events from when the Trigger was suspended until the Trigger
	// resumes execution do not have the Trigger fire for them.
	TolerateResumeErrors *bool `field:"optional" json:"tolerateResumeErrors" yaml:"tolerateResumeErrors"`
	// If `true`, event ordering is disabled and this Trigger can process events in parallel.
	//
	// If `false`, event
	// ordering is enabled and the Trigger executes events
	// serially.
	Unordered *bool `field:"optional" json:"unordered" yaml:"unordered"`
}

