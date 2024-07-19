package awscdkresourcesmongodbatlas


// Name of the parameter in the target object that MongoDB Cloud checks.
//
// The parameter must match all rules for MongoDB Cloud to check for alert configurations.
type MatcherFieldName string

const (
	// CLUSTER_NAME.
	MatcherFieldName_CLUSTER_NAME MatcherFieldName = "CLUSTER_NAME"
	// HOSTNAME.
	MatcherFieldName_HOSTNAME MatcherFieldName = "HOSTNAME"
	// HOSTNAME_AND_PORT.
	MatcherFieldName_HOSTNAME_AND_PORT MatcherFieldName = "HOSTNAME_AND_PORT"
	// PORT.
	MatcherFieldName_PORT MatcherFieldName = "PORT"
	// REPLICA_SET_NAME.
	MatcherFieldName_REPLICA_SET_NAME MatcherFieldName = "REPLICA_SET_NAME"
	// SHARD_NAME.
	MatcherFieldName_SHARD_NAME MatcherFieldName = "SHARD_NAME"
	// TYPE_NAME.
	MatcherFieldName_TYPE_NAME MatcherFieldName = "TYPE_NAME"
	// APPLICATION_ID.
	MatcherFieldName_APPLICATION_ID MatcherFieldName = "APPLICATION_ID"
	// INSTANCE_NAME.
	MatcherFieldName_INSTANCE_NAME MatcherFieldName = "INSTANCE_NAME"
	// PROCESSOR_NAME.
	MatcherFieldName_PROCESSOR_NAME MatcherFieldName = "PROCESSOR_NAME"
)

