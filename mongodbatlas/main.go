package mongodbatlas

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AccessListDefinition",
		reflect.TypeOf((*AccessListDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Action",
		reflect.TypeOf((*Action)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AdvancedAutoScaling",
		reflect.TypeOf((*AdvancedAutoScaling)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AdvancedRegionConfig",
		reflect.TypeOf((*AdvancedRegionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AdvancedReplicationSpec",
		reflect.TypeOf((*AdvancedReplicationSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AlertView",
		reflect.TypeOf((*AlertView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.AlertViewEventTypeName",
		reflect.TypeOf((*AlertViewEventTypeName)(nil)).Elem(),
		map[string]interface{}{
			"AWS_ENCRYPTION_KEY_NEEDS_ROTATION": AlertViewEventTypeName_AWS_ENCRYPTION_KEY_NEEDS_ROTATION,
			"AZURE_ENCRYPTION_KEY_NEEDS_ROTATION": AlertViewEventTypeName_AZURE_ENCRYPTION_KEY_NEEDS_ROTATION,
			"CLUSTER_MONGOS_IS_MISSING": AlertViewEventTypeName_CLUSTER_MONGOS_IS_MISSING,
			"CPS_RESTORE_FAILED": AlertViewEventTypeName_CPS_RESTORE_FAILED,
			"CPS_RESTORE_SUCCESSFUL": AlertViewEventTypeName_CPS_RESTORE_SUCCESSFUL,
			"CPS_SNAPSHOT_BEHIND": AlertViewEventTypeName_CPS_SNAPSHOT_BEHIND,
			"CPS_SNAPSHOT_DOWNLOAD_REQUEST_FAILED": AlertViewEventTypeName_CPS_SNAPSHOT_DOWNLOAD_REQUEST_FAILED,
			"CPS_SNAPSHOT_FALLBACK_FAILED": AlertViewEventTypeName_CPS_SNAPSHOT_FALLBACK_FAILED,
			"CPS_SNAPSHOT_FALLBACK_SUCCESSFUL": AlertViewEventTypeName_CPS_SNAPSHOT_FALLBACK_SUCCESSFUL,
			"CPS_SNAPSHOT_SUCCESSFUL": AlertViewEventTypeName_CPS_SNAPSHOT_SUCCESSFUL,
			"CREDIT_CARD_ABOUT_TO_EXPIRE": AlertViewEventTypeName_CREDIT_CARD_ABOUT_TO_EXPIRE,
			"DAILY_BILL_OVER_THRESHOLD": AlertViewEventTypeName_DAILY_BILL_OVER_THRESHOLD,
			"GCP_ENCRYPTION_KEY_NEEDS_ROTATION": AlertViewEventTypeName_GCP_ENCRYPTION_KEY_NEEDS_ROTATION,
			"HOST_DOWN": AlertViewEventTypeName_HOST_DOWN,
			"JOINED_GROUP": AlertViewEventTypeName_JOINED_GROUP,
			"NDS_X509_USER_AUTHENTICATION_CUSTOMER_CA_EXPIRATION_CHECK": AlertViewEventTypeName_NDS_X509_USER_AUTHENTICATION_CUSTOMER_CA_EXPIRATION_CHECK,
			"NDS_X509_USER_AUTHENTICATION_MANAGED_USER_CERTS_EXPIRATION_CHECK": AlertViewEventTypeName_NDS_X509_USER_AUTHENTICATION_MANAGED_USER_CERTS_EXPIRATION_CHECK,
			"NO_PRIMARY": AlertViewEventTypeName_NO_PRIMARY,
			"OUTSIDE_METRIC_THRESHOLD": AlertViewEventTypeName_OUTSIDE_METRIC_THRESHOLD,
			"OUTSIDE_SERVERLESS_METRIC_THRESHOLD": AlertViewEventTypeName_OUTSIDE_SERVERLESS_METRIC_THRESHOLD,
			"PENDING_INVOICE_OVER_THRESHOLD": AlertViewEventTypeName_PENDING_INVOICE_OVER_THRESHOLD,
			"PRIMARY_ELECTED": AlertViewEventTypeName_PRIMARY_ELECTED,
			"REMOVED_FROM_GROUP": AlertViewEventTypeName_REMOVED_FROM_GROUP,
			"REPLICATION_OPLOG_WINDOW_RUNNING_OUT": AlertViewEventTypeName_REPLICATION_OPLOG_WINDOW_RUNNING_OUT,
			"TOO_MANY_ELECTIONS": AlertViewEventTypeName_TOO_MANY_ELECTIONS,
			"USERS_WITHOUT_MULTIFACTOR_AUTH": AlertViewEventTypeName_USERS_WITHOUT_MULTIFACTOR_AUTH,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.AlertViewMetricName",
		reflect.TypeOf((*AlertViewMetricName)(nil)).Elem(),
		map[string]interface{}{
			"ASSERT_MSG": AlertViewMetricName_ASSERT_MSG,
			"ASSERT_REGULAR": AlertViewMetricName_ASSERT_REGULAR,
			"ASSERT_USER": AlertViewMetricName_ASSERT_USER,
			"ASSERT_WARNING": AlertViewMetricName_ASSERT_WARNING,
			"AVG_COMMAND_EXECUTION_TIME": AlertViewMetricName_AVG_COMMAND_EXECUTION_TIME,
			"AVG_READ_EXECUTION_TIME": AlertViewMetricName_AVG_READ_EXECUTION_TIME,
			"AVG_WRITE_EXECUTION_TIME": AlertViewMetricName_AVG_WRITE_EXECUTION_TIME,
			"CACHE_BYTES_READ_INTO": AlertViewMetricName_CACHE_BYTES_READ_INTO,
			"CACHE_BYTES_WRITTEN_FROM": AlertViewMetricName_CACHE_BYTES_WRITTEN_FROM,
			"CACHE_DIRTY_BYTES": AlertViewMetricName_CACHE_DIRTY_BYTES,
			"CACHE_USED_BYTES": AlertViewMetricName_CACHE_USED_BYTES,
			"COMPUTED_MEMORY": AlertViewMetricName_COMPUTED_MEMORY,
			"CONNECTIONS": AlertViewMetricName_CONNECTIONS,
			"CONNECTIONS_PERCENT": AlertViewMetricName_CONNECTIONS_PERCENT,
			"CURSORS_TOTAL_OPEN": AlertViewMetricName_CURSORS_TOTAL_OPEN,
			"CURSORS_TOTAL_TIMED_OUT": AlertViewMetricName_CURSORS_TOTAL_TIMED_OUT,
			"DB_DATA_SIZE_TOTAL": AlertViewMetricName_DB_DATA_SIZE_TOTAL,
			"DB_INDEX_SIZE_TOTAL": AlertViewMetricName_DB_INDEX_SIZE_TOTAL,
			"DB_STORAGE_TOTAL": AlertViewMetricName_DB_STORAGE_TOTAL,
			"DISK_PARTITION_SPACE_USED_DATA": AlertViewMetricName_DISK_PARTITION_SPACE_USED_DATA,
			"DISK_PARTITION_SPACE_USED_INDEX": AlertViewMetricName_DISK_PARTITION_SPACE_USED_INDEX,
			"DISK_PARTITION_SPACE_USED_JOURNAL": AlertViewMetricName_DISK_PARTITION_SPACE_USED_JOURNAL,
			"DISK_PARTITION_UTILIZATION_DATA": AlertViewMetricName_DISK_PARTITION_UTILIZATION_DATA,
			"DISK_PARTITION_UTILIZATION_INDEX": AlertViewMetricName_DISK_PARTITION_UTILIZATION_INDEX,
			"DISK_PARTITION_UTILIZATION_JOURNAL": AlertViewMetricName_DISK_PARTITION_UTILIZATION_JOURNAL,
			"DOCUMENT_DELETED": AlertViewMetricName_DOCUMENT_DELETED,
			"DOCUMENT_INSERTED": AlertViewMetricName_DOCUMENT_INSERTED,
			"DOCUMENT_RETURNED": AlertViewMetricName_DOCUMENT_RETURNED,
			"DOCUMENT_UPDATED": AlertViewMetricName_DOCUMENT_UPDATED,
			"EXTRA_INFO_PAGE_FAULTS": AlertViewMetricName_EXTRA_INFO_PAGE_FAULTS,
			"FTS_MEMORY_RESIDENT": AlertViewMetricName_FTS_MEMORY_RESIDENT,
			"FTS_MEMORY_SHARED": AlertViewMetricName_FTS_MEMORY_SHARED,
			"FTS_MEMORY_VIRTUAL": AlertViewMetricName_FTS_MEMORY_VIRTUAL,
			"FTS_PROCESS_CPU_KERNEL": AlertViewMetricName_FTS_PROCESS_CPU_KERNEL,
			"FTS_PROCESS_CPU_USER": AlertViewMetricName_FTS_PROCESS_CPU_USER,
			"FTS_PROCESS_DISK": AlertViewMetricName_FTS_PROCESS_DISK,
			"GLOBAL_LOCK_CURRENT_QUEUE_READERS": AlertViewMetricName_GLOBAL_LOCK_CURRENT_QUEUE_READERS,
			"GLOBAL_LOCK_CURRENT_QUEUE_TOTAL": AlertViewMetricName_GLOBAL_LOCK_CURRENT_QUEUE_TOTAL,
			"GLOBAL_LOCK_CURRENT_QUEUE_WRITERS": AlertViewMetricName_GLOBAL_LOCK_CURRENT_QUEUE_WRITERS,
			"LOGICAL_SIZE": AlertViewMetricName_LOGICAL_SIZE,
			"MEMORY_RESIDENT": AlertViewMetricName_MEMORY_RESIDENT,
			"MEMORY_VIRTUAL": AlertViewMetricName_MEMORY_VIRTUAL,
			"NETWORK_BYTES_IN": AlertViewMetricName_NETWORK_BYTES_IN,
			"NETWORK_BYTES_OUT": AlertViewMetricName_NETWORK_BYTES_OUT,
			"NETWORK_NUM_REQUESTS": AlertViewMetricName_NETWORK_NUM_REQUESTS,
			"NORMALIZED_FTS_PROCESS_CPU_KERNEL": AlertViewMetricName_NORMALIZED_FTS_PROCESS_CPU_KERNEL,
			"NORMALIZED_FTS_PROCESS_CPU_USER": AlertViewMetricName_NORMALIZED_FTS_PROCESS_CPU_USER,
			"NORMALIZED_SYSTEM_CPU_STEAL": AlertViewMetricName_NORMALIZED_SYSTEM_CPU_STEAL,
			"NORMALIZED_SYSTEM_CPU_USER": AlertViewMetricName_NORMALIZED_SYSTEM_CPU_USER,
			"OPCOUNTER_CMD": AlertViewMetricName_OPCOUNTER_CMD,
			"OPCOUNTER_DELETE": AlertViewMetricName_OPCOUNTER_DELETE,
			"OPCOUNTER_GETMORE": AlertViewMetricName_OPCOUNTER_GETMORE,
			"OPCOUNTER_INSERT": AlertViewMetricName_OPCOUNTER_INSERT,
			"OPCOUNTER_QUERY": AlertViewMetricName_OPCOUNTER_QUERY,
			"OPCOUNTER_REPL_CMD": AlertViewMetricName_OPCOUNTER_REPL_CMD,
			"OPCOUNTER_REPL_DELETE": AlertViewMetricName_OPCOUNTER_REPL_DELETE,
			"OPCOUNTER_REPL_INSERT": AlertViewMetricName_OPCOUNTER_REPL_INSERT,
			"OPCOUNTER_REPL_UPDATE": AlertViewMetricName_OPCOUNTER_REPL_UPDATE,
			"OPCOUNTER_UPDATE": AlertViewMetricName_OPCOUNTER_UPDATE,
			"OPERATIONS_SCAN_AND_ORDER": AlertViewMetricName_OPERATIONS_SCAN_AND_ORDER,
			"OPLOG_MASTER_LAG_TIME_DIFF": AlertViewMetricName_OPLOG_MASTER_LAG_TIME_DIFF,
			"OPLOG_MASTER_TIME": AlertViewMetricName_OPLOG_MASTER_TIME,
			"OPLOG_RATE_GB_PER_HOUR": AlertViewMetricName_OPLOG_RATE_GB_PER_HOUR,
			"OPLOG_SLAVE_LAG_MASTER_TIME": AlertViewMetricName_OPLOG_SLAVE_LAG_MASTER_TIME,
			"QUERY_EXECUTOR_SCANNED": AlertViewMetricName_QUERY_EXECUTOR_SCANNED,
			"QUERY_EXECUTOR_SCANNED_OBJECTS": AlertViewMetricName_QUERY_EXECUTOR_SCANNED_OBJECTS,
			"QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED": AlertViewMetricName_QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED,
			"QUERY_TARGETING_SCANNED_PER_RETURNED": AlertViewMetricName_QUERY_TARGETING_SCANNED_PER_RETURNED,
			"RESTARTS_IN_LAST_HOUR": AlertViewMetricName_RESTARTS_IN_LAST_HOUR,
			"SERVERLESS_CONNECTIONS": AlertViewMetricName_SERVERLESS_CONNECTIONS,
			"SERVERLESS_CONNECTIONS_PERCENT": AlertViewMetricName_SERVERLESS_CONNECTIONS_PERCENT,
			"SERVERLESS_DATA_SIZE_TOTAL": AlertViewMetricName_SERVERLESS_DATA_SIZE_TOTAL,
			"SERVERLESS_NETWORK_BYTES_IN": AlertViewMetricName_SERVERLESS_NETWORK_BYTES_IN,
			"SERVERLESS_NETWORK_BYTES_OUT": AlertViewMetricName_SERVERLESS_NETWORK_BYTES_OUT,
			"SERVERLESS_NETWORK_NUM_REQUESTS": AlertViewMetricName_SERVERLESS_NETWORK_NUM_REQUESTS,
			"SERVERLESS_OPCOUNTER_CMD": AlertViewMetricName_SERVERLESS_OPCOUNTER_CMD,
			"SERVERLESS_OPCOUNTER_DELETE": AlertViewMetricName_SERVERLESS_OPCOUNTER_DELETE,
			"SERVERLESS_OPCOUNTER_GETMORE": AlertViewMetricName_SERVERLESS_OPCOUNTER_GETMORE,
			"SERVERLESS_OPCOUNTER_INSERT": AlertViewMetricName_SERVERLESS_OPCOUNTER_INSERT,
			"SERVERLESS_OPCOUNTER_QUERY": AlertViewMetricName_SERVERLESS_OPCOUNTER_QUERY,
			"SERVERLESS_OPCOUNTER_UPDATE": AlertViewMetricName_SERVERLESS_OPCOUNTER_UPDATE,
			"SERVERLESS_TOTAL_READ_UNITS": AlertViewMetricName_SERVERLESS_TOTAL_READ_UNITS,
			"SERVERLESS_TOTAL_WRITE_UNITS": AlertViewMetricName_SERVERLESS_TOTAL_WRITE_UNITS,
			"TICKETS_AVAILABLE_READS": AlertViewMetricName_TICKETS_AVAILABLE_READS,
			"TICKETS_AVAILABLE_WRITES": AlertViewMetricName_TICKETS_AVAILABLE_WRITES,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.AlertViewStatus",
		reflect.TypeOf((*AlertViewStatus)(nil)).Elem(),
		map[string]interface{}{
			"CANCELLED": AlertViewStatus_CANCELLED,
			"CLOSED": AlertViewStatus_CLOSED,
			"OPEN": AlertViewStatus_OPEN,
			"TRACKING": AlertViewStatus_TRACKING,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiAtlasDiskBackupCopySettingView",
		reflect.TypeOf((*ApiAtlasDiskBackupCopySettingView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiAtlasDiskBackupShardedClusterSnapshotMemberView",
		reflect.TypeOf((*ApiAtlasDiskBackupShardedClusterSnapshotMemberView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider",
		reflect.TypeOf((*ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider)(nil)).Elem(),
		map[string]interface{}{
			"AWS": ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider_AWS,
			"AZURE": ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider_AZURE,
			"GCP": ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider_GCP,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiAtlasDiskBackupShardedClusterSnapshotView",
		reflect.TypeOf((*ApiAtlasDiskBackupShardedClusterSnapshotView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType",
		reflect.TypeOf((*ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType)(nil)).Elem(),
		map[string]interface{}{
			"HOURLY": ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType_HOURLY,
			"DAILY": ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType_DAILY,
			"WEEKLY": ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType_WEEKLY,
			"MONTHLY": ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType_MONTHLY,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType",
		reflect.TypeOf((*ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType)(nil)).Elem(),
		map[string]interface{}{
			"ON_DEMAND": ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType_ON_DEMAND,
			"SCHEDULED": ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType_SCHEDULED,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.ApiAtlasDiskBackupShardedClusterSnapshotViewStatus",
		reflect.TypeOf((*ApiAtlasDiskBackupShardedClusterSnapshotViewStatus)(nil)).Elem(),
		map[string]interface{}{
			"QUEUED": ApiAtlasDiskBackupShardedClusterSnapshotViewStatus_QUEUED,
			"IN_PROGRESS": ApiAtlasDiskBackupShardedClusterSnapshotViewStatus_IN_PROGRESS,
			"COMPLETED": ApiAtlasDiskBackupShardedClusterSnapshotViewStatus_COMPLETED,
			"FAILED": ApiAtlasDiskBackupShardedClusterSnapshotViewStatus_FAILED,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.ApiAtlasDiskBackupShardedClusterSnapshotViewType",
		reflect.TypeOf((*ApiAtlasDiskBackupShardedClusterSnapshotViewType)(nil)).Elem(),
		map[string]interface{}{
			"REPLICA_SET": ApiAtlasDiskBackupShardedClusterSnapshotViewType_REPLICA_SET,
			"SHARDED_CLUSTER": ApiAtlasDiskBackupShardedClusterSnapshotViewType_SHARDED_CLUSTER,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiAtlasFtsAnalyzersViewManual",
		reflect.TypeOf((*ApiAtlasFtsAnalyzersViewManual)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiAtlasFtsMappingsViewManual",
		reflect.TypeOf((*ApiAtlasFtsMappingsViewManual)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiAtlasFtsSynonymMappingDefinitionView",
		reflect.TypeOf((*ApiAtlasFtsSynonymMappingDefinitionView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiAtlasNdsUserToDnMappingView",
		reflect.TypeOf((*ApiAtlasNdsUserToDnMappingView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiDeleteCopiedBackupsView",
		reflect.TypeOf((*ApiDeleteCopiedBackupsView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiKeyDefinition",
		reflect.TypeOf((*ApiKeyDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiPolicyItemView",
		reflect.TypeOf((*ApiPolicyItemView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiPolicyView",
		reflect.TypeOf((*ApiPolicyView)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.AtlasBasic",
		reflect.TypeOf((*AtlasBasic)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ipAccessList", GoGetter: "IpAccessList"},
			_jsii_.MemberProperty{JsiiProperty: "mCluster", GoGetter: "MCluster"},
			_jsii_.MemberProperty{JsiiProperty: "mDBUser", GoGetter: "MDBUser"},
			_jsii_.MemberProperty{JsiiProperty: "mProject", GoGetter: "MProject"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AtlasBasic{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.AtlasBasicPrivateEndpoint",
		reflect.TypeOf((*AtlasBasicPrivateEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "atlasBasic", GoGetter: "AtlasBasic"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "privateEndpoint", GoGetter: "PrivateEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AtlasBasicPrivateEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AtlasBasicPrivateEndpointProps",
		reflect.TypeOf((*AtlasBasicPrivateEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AtlasBasicProps",
		reflect.TypeOf((*AtlasBasicProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.AtlasEncryptionAtRest",
		reflect.TypeOf((*AtlasEncryptionAtRest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cfnEncryptionAtRest", GoGetter: "CfnEncryptionAtRest"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AtlasEncryptionAtRest{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.AtlasEncryptionAtRestExpress",
		reflect.TypeOf((*AtlasEncryptionAtRestExpress)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessList", GoGetter: "AccessList"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "databaseUser", GoGetter: "DatabaseUser"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionAtRest", GoGetter: "EncryptionAtRest"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AtlasEncryptionAtRestExpress{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AtlasEncryptionAtRestExpressProps",
		reflect.TypeOf((*AtlasEncryptionAtRestExpressProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AtlasEncryptionAtRestProps",
		reflect.TypeOf((*AtlasEncryptionAtRestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AuthConfig",
		reflect.TypeOf((*AuthConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.AuthConfigOperationType",
		reflect.TypeOf((*AuthConfigOperationType)(nil)).Elem(),
		map[string]interface{}{
			"LOGIN": AuthConfigOperationType_LOGIN,
			"CREATE": AuthConfigOperationType_CREATE,
			"DELETE": AuthConfigOperationType_DELETE,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.AuthConfigProviders",
		reflect.TypeOf((*AuthConfigProviders)(nil)).Elem(),
		map[string]interface{}{
			"ANON_USER": AuthConfigProviders_ANON_USER,
			"API_KEY": AuthConfigProviders_API_KEY,
			"CUSTOM_TOKEN": AuthConfigProviders_CUSTOM_TOKEN,
			"CUSTOM_FUNCTION": AuthConfigProviders_CUSTOM_FUNCTION,
			"LOCAL_USERPASS": AuthConfigProviders_LOCAL_USERPASS,
			"OAUTH2_APPLE": AuthConfigProviders_OAUTH2_APPLE,
			"OAUTH2_FACEBOOK": AuthConfigProviders_OAUTH2_FACEBOOK,
			"OAUTH2_GOOGLE": AuthConfigProviders_OAUTH2_GOOGLE,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AwsKmsConfiguration",
		reflect.TypeOf((*AwsKmsConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnAlertConfiguration",
		reflect.TypeOf((*CfnAlertConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreated", GoGetter: "AttrCreated"},
			_jsii_.MemberProperty{JsiiProperty: "attrEnabled", GoGetter: "AttrEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "attrGroupId", GoGetter: "AttrGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrTotalCount", GoGetter: "AttrTotalCount"},
			_jsii_.MemberProperty{JsiiProperty: "attrTypeName", GoGetter: "AttrTypeName"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdated", GoGetter: "AttrUpdated"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAlertConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnAlertConfigurationProps",
		reflect.TypeOf((*CfnAlertConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnAlertConfigurationPropsEventTypeName",
		reflect.TypeOf((*CfnAlertConfigurationPropsEventTypeName)(nil)).Elem(),
		map[string]interface{}{
			"AWS_ENCRYPTION_KEY_NEEDS_ROTATION": CfnAlertConfigurationPropsEventTypeName_AWS_ENCRYPTION_KEY_NEEDS_ROTATION,
			"AZURE_ENCRYPTION_KEY_NEEDS_ROTATION": CfnAlertConfigurationPropsEventTypeName_AZURE_ENCRYPTION_KEY_NEEDS_ROTATION,
			"CLUSTER_MONGOS_IS_MISSING": CfnAlertConfigurationPropsEventTypeName_CLUSTER_MONGOS_IS_MISSING,
			"CPS_RESTORE_FAILED": CfnAlertConfigurationPropsEventTypeName_CPS_RESTORE_FAILED,
			"CPS_RESTORE_SUCCESSFUL": CfnAlertConfigurationPropsEventTypeName_CPS_RESTORE_SUCCESSFUL,
			"CPS_SNAPSHOT_BEHIND": CfnAlertConfigurationPropsEventTypeName_CPS_SNAPSHOT_BEHIND,
			"CPS_SNAPSHOT_DOWNLOAD_REQUEST_FAILED": CfnAlertConfigurationPropsEventTypeName_CPS_SNAPSHOT_DOWNLOAD_REQUEST_FAILED,
			"CPS_SNAPSHOT_FALLBACK_FAILED": CfnAlertConfigurationPropsEventTypeName_CPS_SNAPSHOT_FALLBACK_FAILED,
			"CPS_SNAPSHOT_FALLBACK_SUCCESSFUL": CfnAlertConfigurationPropsEventTypeName_CPS_SNAPSHOT_FALLBACK_SUCCESSFUL,
			"CPS_SNAPSHOT_SUCCESSFUL": CfnAlertConfigurationPropsEventTypeName_CPS_SNAPSHOT_SUCCESSFUL,
			"CREDIT_CARD_ABOUT_TO_EXPIRE": CfnAlertConfigurationPropsEventTypeName_CREDIT_CARD_ABOUT_TO_EXPIRE,
			"DAILY_BILL_OVER_THRESHOLD": CfnAlertConfigurationPropsEventTypeName_DAILY_BILL_OVER_THRESHOLD,
			"GCP_ENCRYPTION_KEY_NEEDS_ROTATION": CfnAlertConfigurationPropsEventTypeName_GCP_ENCRYPTION_KEY_NEEDS_ROTATION,
			"HOST_DOWN": CfnAlertConfigurationPropsEventTypeName_HOST_DOWN,
			"JOINED_GROUP": CfnAlertConfigurationPropsEventTypeName_JOINED_GROUP,
			"NDS_X509_USER_AUTHENTICATION_CUSTOMER_CA_EXPIRATION_CHECK": CfnAlertConfigurationPropsEventTypeName_NDS_X509_USER_AUTHENTICATION_CUSTOMER_CA_EXPIRATION_CHECK,
			"NDS_X509_USER_AUTHENTICATION_MANAGED_USER_CERTS_EXPIRATION_CHECK": CfnAlertConfigurationPropsEventTypeName_NDS_X509_USER_AUTHENTICATION_MANAGED_USER_CERTS_EXPIRATION_CHECK,
			"NO_PRIMARY": CfnAlertConfigurationPropsEventTypeName_NO_PRIMARY,
			"OUTSIDE_METRIC_THRESHOLD": CfnAlertConfigurationPropsEventTypeName_OUTSIDE_METRIC_THRESHOLD,
			"OUTSIDE_SERVERLESS_METRIC_THRESHOLD": CfnAlertConfigurationPropsEventTypeName_OUTSIDE_SERVERLESS_METRIC_THRESHOLD,
			"OUTSIDE_REALM_METRIC_THRESHOLD": CfnAlertConfigurationPropsEventTypeName_OUTSIDE_REALM_METRIC_THRESHOLD,
			"PENDING_INVOICE_OVER_THRESHOLD": CfnAlertConfigurationPropsEventTypeName_PENDING_INVOICE_OVER_THRESHOLD,
			"PRIMARY_ELECTED": CfnAlertConfigurationPropsEventTypeName_PRIMARY_ELECTED,
			"REMOVED_FROM_GROUP": CfnAlertConfigurationPropsEventTypeName_REMOVED_FROM_GROUP,
			"REPLICATION_OPLOG_WINDOW_RUNNING_OUT": CfnAlertConfigurationPropsEventTypeName_REPLICATION_OPLOG_WINDOW_RUNNING_OUT,
			"TOO_MANY_ELECTIONS": CfnAlertConfigurationPropsEventTypeName_TOO_MANY_ELECTIONS,
			"USER_ROLES_CHANGED_AUDIT": CfnAlertConfigurationPropsEventTypeName_USER_ROLES_CHANGED_AUDIT,
			"USERS_WITHOUT_MULTIFACTOR_AUTH": CfnAlertConfigurationPropsEventTypeName_USERS_WITHOUT_MULTIFACTOR_AUTH,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnAuditing",
		reflect.TypeOf((*CfnAuditing)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAuditAuthorizationSuccess", GoGetter: "AttrAuditAuthorizationSuccess"},
			_jsii_.MemberProperty{JsiiProperty: "attrAuditFilter", GoGetter: "AttrAuditFilter"},
			_jsii_.MemberProperty{JsiiProperty: "attrConfigurationType", GoGetter: "AttrConfigurationType"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAuditing{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnAuditingProps",
		reflect.TypeOf((*CfnAuditingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnCloudBackUpRestoreJobs",
		reflect.TypeOf((*CfnCloudBackUpRestoreJobs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrDeliveryUrl", GoGetter: "AttrDeliveryUrl"},
			_jsii_.MemberProperty{JsiiProperty: "attrExpiresAt", GoGetter: "AttrExpiresAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrFinishedAt", GoGetter: "AttrFinishedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrTimestamp", GoGetter: "AttrTimestamp"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCloudBackUpRestoreJobs{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnCloudBackUpRestoreJobsProps",
		reflect.TypeOf((*CfnCloudBackUpRestoreJobsProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnCloudBackUpRestoreJobsPropsDeliveryType",
		reflect.TypeOf((*CfnCloudBackUpRestoreJobsPropsDeliveryType)(nil)).Elem(),
		map[string]interface{}{
			"DOWNLOAD": CfnCloudBackUpRestoreJobsPropsDeliveryType_DOWNLOAD,
			"AUTOMATED": CfnCloudBackUpRestoreJobsPropsDeliveryType_AUTOMATED,
			"POINT_IN_TIME": CfnCloudBackUpRestoreJobsPropsDeliveryType_POINT_IN_TIME,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnCloudBackupSchedule",
		reflect.TypeOf((*CfnCloudBackupSchedule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrClusterId", GoGetter: "AttrClusterId"},
			_jsii_.MemberProperty{JsiiProperty: "attrNextSnapshot", GoGetter: "AttrNextSnapshot"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCloudBackupSchedule{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnCloudBackupScheduleProps",
		reflect.TypeOf((*CfnCloudBackupScheduleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnCloudBackupSnapshot",
		reflect.TypeOf((*CfnCloudBackupSnapshot)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCloudProvider", GoGetter: "AttrCloudProvider"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrExpiresAt", GoGetter: "AttrExpiresAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrMasterKeyUUID", GoGetter: "AttrMasterKeyUUID"},
			_jsii_.MemberProperty{JsiiProperty: "attrMongodVersion", GoGetter: "AttrMongodVersion"},
			_jsii_.MemberProperty{JsiiProperty: "attrReplicaSetName", GoGetter: "AttrReplicaSetName"},
			_jsii_.MemberProperty{JsiiProperty: "attrSnapshotId", GoGetter: "AttrSnapshotId"},
			_jsii_.MemberProperty{JsiiProperty: "attrSnapshotIds", GoGetter: "AttrSnapshotIds"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrStorageSizeBytes", GoGetter: "AttrStorageSizeBytes"},
			_jsii_.MemberProperty{JsiiProperty: "attrType", GoGetter: "AttrType"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCloudBackupSnapshot{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnCloudBackupSnapshotExportBucket",
		reflect.TypeOf((*CfnCloudBackupSnapshotExportBucket)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCloudBackupSnapshotExportBucket{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnCloudBackupSnapshotExportBucketProps",
		reflect.TypeOf((*CfnCloudBackupSnapshotExportBucketProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnCloudBackupSnapshotProps",
		reflect.TypeOf((*CfnCloudBackupSnapshotProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnCloudBackupSnapshotPropsFrequencyType",
		reflect.TypeOf((*CfnCloudBackupSnapshotPropsFrequencyType)(nil)).Elem(),
		map[string]interface{}{
			"HOURLY": CfnCloudBackupSnapshotPropsFrequencyType_HOURLY,
			"DAILY": CfnCloudBackupSnapshotPropsFrequencyType_DAILY,
			"WEEKLY": CfnCloudBackupSnapshotPropsFrequencyType_WEEKLY,
			"MONTHLY": CfnCloudBackupSnapshotPropsFrequencyType_MONTHLY,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnCloudBackupSnapshotPropsSnapshotType",
		reflect.TypeOf((*CfnCloudBackupSnapshotPropsSnapshotType)(nil)).Elem(),
		map[string]interface{}{
			"ON_DEMAND": CfnCloudBackupSnapshotPropsSnapshotType_ON_DEMAND,
			"SCHEDULED": CfnCloudBackupSnapshotPropsSnapshotType_SCHEDULED,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnCluster",
		reflect.TypeOf((*CfnCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedDate", GoGetter: "AttrCreatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrMongoDBVersion", GoGetter: "AttrMongoDBVersion"},
			_jsii_.MemberProperty{JsiiProperty: "attrStateName", GoGetter: "AttrStateName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCluster{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnClusterProps",
		reflect.TypeOf((*CfnClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnClusterPropsBiConnector",
		reflect.TypeOf((*CfnClusterPropsBiConnector)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnClusterPropsEncryptionAtRestProvider",
		reflect.TypeOf((*CfnClusterPropsEncryptionAtRestProvider)(nil)).Elem(),
		map[string]interface{}{
			"AWS": CfnClusterPropsEncryptionAtRestProvider_AWS,
			"GCP": CfnClusterPropsEncryptionAtRestProvider_GCP,
			"AZURE": CfnClusterPropsEncryptionAtRestProvider_AZURE,
			"NONE": CfnClusterPropsEncryptionAtRestProvider_NONE,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnClusterPropsLabels",
		reflect.TypeOf((*CfnClusterPropsLabels)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnCustomDbRole",
		reflect.TypeOf((*CfnCustomDbRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCustomDbRole{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnCustomDbRoleProps",
		reflect.TypeOf((*CfnCustomDbRoleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnCustomDnsConfigurationClusterAws",
		reflect.TypeOf((*CfnCustomDnsConfigurationClusterAws)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCustomDnsConfigurationClusterAws{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnCustomDnsConfigurationClusterAwsProps",
		reflect.TypeOf((*CfnCustomDnsConfigurationClusterAwsProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnDataLakes",
		reflect.TypeOf((*CfnDataLakes)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrHostnames", GoGetter: "AttrHostnames"},
			_jsii_.MemberProperty{JsiiProperty: "attrStartDate", GoGetter: "AttrStartDate"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataLakes{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnDataLakesProps",
		reflect.TypeOf((*CfnDataLakesProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnDatabaseUser",
		reflect.TypeOf((*CfnDatabaseUser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrUserCFNIdentifier", GoGetter: "AttrUserCFNIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDatabaseUser{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnDatabaseUserProps",
		reflect.TypeOf((*CfnDatabaseUserProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnDatabaseUserPropsAwsiamType",
		reflect.TypeOf((*CfnDatabaseUserPropsAwsiamType)(nil)).Elem(),
		map[string]interface{}{
			"NONE": CfnDatabaseUserPropsAwsiamType_NONE,
			"USER": CfnDatabaseUserPropsAwsiamType_USER,
			"ROLE": CfnDatabaseUserPropsAwsiamType_ROLE,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnDatabaseUserPropsLdapAuthType",
		reflect.TypeOf((*CfnDatabaseUserPropsLdapAuthType)(nil)).Elem(),
		map[string]interface{}{
			"NONE": CfnDatabaseUserPropsLdapAuthType_NONE,
			"USER": CfnDatabaseUserPropsLdapAuthType_USER,
			"GROUP": CfnDatabaseUserPropsLdapAuthType_GROUP,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnDatabaseUserPropsX509Type",
		reflect.TypeOf((*CfnDatabaseUserPropsX509Type)(nil)).Elem(),
		map[string]interface{}{
			"NONE": CfnDatabaseUserPropsX509Type_NONE,
			"MANAGED": CfnDatabaseUserPropsX509Type_MANAGED,
			"CUSTOMER": CfnDatabaseUserPropsX509Type_CUSTOMER,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnEncryptionAtRest",
		reflect.TypeOf((*CfnEncryptionAtRest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEncryptionAtRest{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnEncryptionAtRestProps",
		reflect.TypeOf((*CfnEncryptionAtRestProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnFederatedSettingsOrgRoleMapping",
		reflect.TypeOf((*CfnFederatedSettingsOrgRoleMapping)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFederatedSettingsOrgRoleMapping{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnFederatedSettingsOrgRoleMappingProps",
		reflect.TypeOf((*CfnFederatedSettingsOrgRoleMappingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnGlobalClusterConfig",
		reflect.TypeOf((*CfnGlobalClusterConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrRemoveAllZoneMapping", GoGetter: "AttrRemoveAllZoneMapping"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGlobalClusterConfig{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnGlobalClusterConfigProps",
		reflect.TypeOf((*CfnGlobalClusterConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnLdapConfiguration",
		reflect.TypeOf((*CfnLdapConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLdapConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnLdapConfigurationProps",
		reflect.TypeOf((*CfnLdapConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnLdapVerify",
		reflect.TypeOf((*CfnLdapVerify)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrRequestId", GoGetter: "AttrRequestId"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLdapVerify{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnLdapVerifyProps",
		reflect.TypeOf((*CfnLdapVerifyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnMaintenanceWindow",
		reflect.TypeOf((*CfnMaintenanceWindow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMaintenanceWindow{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnMaintenanceWindowProps",
		reflect.TypeOf((*CfnMaintenanceWindowProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnNetworkContainer",
		reflect.TypeOf((*CfnNetworkContainer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNetworkContainer{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnNetworkContainerProps",
		reflect.TypeOf((*CfnNetworkContainerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnNetworkPeering",
		reflect.TypeOf((*CfnNetworkPeering)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrConnectionId", GoGetter: "AttrConnectionId"},
			_jsii_.MemberProperty{JsiiProperty: "attrErrorStateName", GoGetter: "AttrErrorStateName"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatusName", GoGetter: "AttrStatusName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNetworkPeering{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnNetworkPeeringProps",
		reflect.TypeOf((*CfnNetworkPeeringProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnOnlineArchive",
		reflect.TypeOf((*CfnOnlineArchive)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArchiveId", GoGetter: "AttrArchiveId"},
			_jsii_.MemberProperty{JsiiProperty: "attrProjectId", GoGetter: "AttrProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "attrTotalCount", GoGetter: "AttrTotalCount"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOnlineArchive{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnOnlineArchiveProps",
		reflect.TypeOf((*CfnOnlineArchiveProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnOnlineArchivePropsCollectionType",
		reflect.TypeOf((*CfnOnlineArchivePropsCollectionType)(nil)).Elem(),
		map[string]interface{}{
			"STANDARD": CfnOnlineArchivePropsCollectionType_STANDARD,
			"TIMESERIES": CfnOnlineArchivePropsCollectionType_TIMESERIES,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnOrgInvitation",
		reflect.TypeOf((*CfnOrgInvitation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrExpiresAt", GoGetter: "AttrExpiresAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrInviterUsername", GoGetter: "AttrInviterUsername"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOrgInvitation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnOrgInvitationProps",
		reflect.TypeOf((*CfnOrgInvitationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnOrgInvitationPropsRoles",
		reflect.TypeOf((*CfnOrgInvitationPropsRoles)(nil)).Elem(),
		map[string]interface{}{
			"ORG_OWNER": CfnOrgInvitationPropsRoles_ORG_OWNER,
			"ORG_MEMBER": CfnOrgInvitationPropsRoles_ORG_MEMBER,
			"ORG_GROUP_CREATOR": CfnOrgInvitationPropsRoles_ORG_GROUP_CREATOR,
			"ORG_BILLING_ADMIN": CfnOrgInvitationPropsRoles_ORG_BILLING_ADMIN,
			"ORG_READ_ONLY": CfnOrgInvitationPropsRoles_ORG_READ_ONLY,
			"GROUP_CLUSTER_MANAGER": CfnOrgInvitationPropsRoles_GROUP_CLUSTER_MANAGER,
			"GROUP_DATA_ACCESS_ADMIN": CfnOrgInvitationPropsRoles_GROUP_DATA_ACCESS_ADMIN,
			"GROUP_DATA_ACCESS_READ_ONLY": CfnOrgInvitationPropsRoles_GROUP_DATA_ACCESS_READ_ONLY,
			"GROUP_DATA_ACCESS_READ_WRITE": CfnOrgInvitationPropsRoles_GROUP_DATA_ACCESS_READ_WRITE,
			"GROUP_OWNER": CfnOrgInvitationPropsRoles_GROUP_OWNER,
			"GROUP_READ_ONLY": CfnOrgInvitationPropsRoles_GROUP_READ_ONLY,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnPrivateEndPointRegionalMode",
		reflect.TypeOf((*CfnPrivateEndPointRegionalMode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPrivateEndPointRegionalMode{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnPrivateEndPointRegionalModeProps",
		reflect.TypeOf((*CfnPrivateEndPointRegionalModeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnPrivateEndpoint",
		reflect.TypeOf((*CfnPrivateEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrInterfaceEndpoints", GoGetter: "AttrInterfaceEndpoints"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPrivateEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnPrivateEndpointAdl",
		reflect.TypeOf((*CfnPrivateEndpointAdl)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPrivateEndpointAdl{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnPrivateEndpointAdlProps",
		reflect.TypeOf((*CfnPrivateEndpointAdlProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnPrivateEndpointProps",
		reflect.TypeOf((*CfnPrivateEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnProject",
		reflect.TypeOf((*CfnProject)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrClusterCount", GoGetter: "AttrClusterCount"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreated", GoGetter: "AttrCreated"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrProjectOwnerId", GoGetter: "AttrProjectOwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProject{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnProjectInvitation",
		reflect.TypeOf((*CfnProjectInvitation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrExpiresAt", GoGetter: "AttrExpiresAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrInviterUsername", GoGetter: "AttrInviterUsername"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectInvitation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnProjectInvitationProps",
		reflect.TypeOf((*CfnProjectInvitationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnProjectInvitationPropsRoles",
		reflect.TypeOf((*CfnProjectInvitationPropsRoles)(nil)).Elem(),
		map[string]interface{}{
			"GROUP_CLUSTER_MANAGER": CfnProjectInvitationPropsRoles_GROUP_CLUSTER_MANAGER,
			"GROUP_DATA_ACCESS_ADMIN": CfnProjectInvitationPropsRoles_GROUP_DATA_ACCESS_ADMIN,
			"GROUP_DATA_ACCESS_READ_ONLY": CfnProjectInvitationPropsRoles_GROUP_DATA_ACCESS_READ_ONLY,
			"GROUP_DATA_ACCESS_READ_WRITE": CfnProjectInvitationPropsRoles_GROUP_DATA_ACCESS_READ_WRITE,
			"GROUP_OWNER": CfnProjectInvitationPropsRoles_GROUP_OWNER,
			"GROUP_READ_ONLY": CfnProjectInvitationPropsRoles_GROUP_READ_ONLY,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnProjectIpAccessList",
		reflect.TypeOf((*CfnProjectIpAccessList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectIpAccessList{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnProjectIpAccessListProps",
		reflect.TypeOf((*CfnProjectIpAccessListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnProjectProps",
		reflect.TypeOf((*CfnProjectProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnSearchIndex",
		reflect.TypeOf((*CfnSearchIndex)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrIndexId", GoGetter: "AttrIndexId"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSearchIndex{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnSearchIndexProps",
		reflect.TypeOf((*CfnSearchIndexProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnServerlessInstance",
		reflect.TypeOf((*CfnServerlessInstance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreateDate", GoGetter: "AttrCreateDate"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrMongoDBVersion", GoGetter: "AttrMongoDBVersion"},
			_jsii_.MemberProperty{JsiiProperty: "attrStateName", GoGetter: "AttrStateName"},
			_jsii_.MemberProperty{JsiiProperty: "attrTotalCount", GoGetter: "AttrTotalCount"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServerlessInstance{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnServerlessInstanceProps",
		reflect.TypeOf((*CfnServerlessInstanceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnTeams",
		reflect.TypeOf((*CfnTeams)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrTeamId", GoGetter: "AttrTeamId"},
			_jsii_.MemberProperty{JsiiProperty: "attrUsers", GoGetter: "AttrUsers"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTeams{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnTeamsProps",
		reflect.TypeOf((*CfnTeamsProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnTeamsPropsRoleNames",
		reflect.TypeOf((*CfnTeamsPropsRoleNames)(nil)).Elem(),
		map[string]interface{}{
			"GROUP_CLUSTER_MANAGER": CfnTeamsPropsRoleNames_GROUP_CLUSTER_MANAGER,
			"GROUP_DATA_ACCESS_ADMIN": CfnTeamsPropsRoleNames_GROUP_DATA_ACCESS_ADMIN,
			"GROUP_DATA_ACCESS_READ_ONLY": CfnTeamsPropsRoleNames_GROUP_DATA_ACCESS_READ_ONLY,
			"GROUP_DATA_ACCESS_READ_WRITE": CfnTeamsPropsRoleNames_GROUP_DATA_ACCESS_READ_WRITE,
			"GROUP_OWNER": CfnTeamsPropsRoleNames_GROUP_OWNER,
			"GROUP_READ_ONLY": CfnTeamsPropsRoleNames_GROUP_READ_ONLY,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnThirdPartyIntegration",
		reflect.TypeOf((*CfnThirdPartyIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnThirdPartyIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnThirdPartyIntegrationProps",
		reflect.TypeOf((*CfnThirdPartyIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnThirdPartyIntegrationPropsScheme",
		reflect.TypeOf((*CfnThirdPartyIntegrationPropsScheme)(nil)).Elem(),
		map[string]interface{}{
			"HTTP": CfnThirdPartyIntegrationPropsScheme_HTTP,
			"HTTPS": CfnThirdPartyIntegrationPropsScheme_HTTPS,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnThirdPartyIntegrationPropsServiceDiscovery",
		reflect.TypeOf((*CfnThirdPartyIntegrationPropsServiceDiscovery)(nil)).Elem(),
		map[string]interface{}{
			"HTTP": CfnThirdPartyIntegrationPropsServiceDiscovery_HTTP,
			"FILE": CfnThirdPartyIntegrationPropsServiceDiscovery_FILE,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnThirdPartyIntegrationPropsType",
		reflect.TypeOf((*CfnThirdPartyIntegrationPropsType)(nil)).Elem(),
		map[string]interface{}{
			"PAGER_DUTY": CfnThirdPartyIntegrationPropsType_PAGER_DUTY,
			"MICROSOFT_TEAMS": CfnThirdPartyIntegrationPropsType_MICROSOFT_TEAMS,
			"SLACK": CfnThirdPartyIntegrationPropsType_SLACK,
			"DATADOG": CfnThirdPartyIntegrationPropsType_DATADOG,
			"OPS_GENIE": CfnThirdPartyIntegrationPropsType_OPS_GENIE,
			"VICTOR_OPS": CfnThirdPartyIntegrationPropsType_VICTOR_OPS,
			"WEBHOOK": CfnThirdPartyIntegrationPropsType_WEBHOOK,
			"PROMETHEUS": CfnThirdPartyIntegrationPropsType_PROMETHEUS,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnTrigger",
		reflect.TypeOf((*CfnTrigger)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrigger{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnTriggerProps",
		reflect.TypeOf((*CfnTriggerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnX509AuthenticationDatabaseUser",
		reflect.TypeOf((*CfnX509AuthenticationDatabaseUser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrLinks", GoGetter: "AttrLinks"},
			_jsii_.MemberProperty{JsiiProperty: "attrResults", GoGetter: "AttrResults"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnX509AuthenticationDatabaseUser{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnX509AuthenticationDatabaseUserProps",
		reflect.TypeOf((*CfnX509AuthenticationDatabaseUserProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ClusterProps",
		reflect.TypeOf((*ClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Compute",
		reflect.TypeOf((*Compute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ConnectionStrings",
		reflect.TypeOf((*ConnectionStrings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CriteriaView",
		reflect.TypeOf((*CriteriaView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CriteriaViewDateFormat",
		reflect.TypeOf((*CriteriaViewDateFormat)(nil)).Elem(),
		map[string]interface{}{
			"ISODATE": CriteriaViewDateFormat_ISODATE,
			"EPOCH_SECONDS": CriteriaViewDateFormat_EPOCH_SECONDS,
			"EPOCH_MILLIS": CriteriaViewDateFormat_EPOCH_MILLIS,
			"EPOCH_NANOSECONDS": CriteriaViewDateFormat_EPOCH_NANOSECONDS,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CriteriaViewType",
		reflect.TypeOf((*CriteriaViewType)(nil)).Elem(),
		map[string]interface{}{
			"DATE": CriteriaViewType_DATE,
			"CUSTOM": CriteriaViewType_CUSTOM,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CurrentValue",
		reflect.TypeOf((*CurrentValue)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CurrentValueUnits",
		reflect.TypeOf((*CurrentValueUnits)(nil)).Elem(),
		map[string]interface{}{
			"BITS": CurrentValueUnits_BITS,
			"BYTES": CurrentValueUnits_BYTES,
			"DAYS": CurrentValueUnits_DAYS,
			"GIGABITS": CurrentValueUnits_GIGABITS,
			"GIGABYTES": CurrentValueUnits_GIGABYTES,
			"HOURS": CurrentValueUnits_HOURS,
			"KILOBITS": CurrentValueUnits_KILOBITS,
			"KILOBYTES": CurrentValueUnits_KILOBYTES,
			"MEGABITS": CurrentValueUnits_MEGABITS,
			"MEGABYTES": CurrentValueUnits_MEGABYTES,
			"MILLISECONDS": CurrentValueUnits_MILLISECONDS,
			"MINUTES": CurrentValueUnits_MINUTES,
			"PETABYTES": CurrentValueUnits_PETABYTES,
			"RAW": CurrentValueUnits_RAW,
			"SECONDS": CurrentValueUnits_SECONDS,
			"TERABYTES": CurrentValueUnits_TERABYTES,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CustomerX509",
		reflect.TypeOf((*CustomerX509)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DataLakeAwsCloudProviderConfigView",
		reflect.TypeOf((*DataLakeAwsCloudProviderConfigView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DataLakeCloudProviderConfigView",
		reflect.TypeOf((*DataLakeCloudProviderConfigView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DataLakeDataProcessRegionView",
		reflect.TypeOf((*DataLakeDataProcessRegionView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.DataLakeDataProcessRegionViewCloudProvider",
		reflect.TypeOf((*DataLakeDataProcessRegionViewCloudProvider)(nil)).Elem(),
		map[string]interface{}{
			"AWS": DataLakeDataProcessRegionViewCloudProvider_AWS,
			"GCP": DataLakeDataProcessRegionViewCloudProvider_GCP,
			"AZURE": DataLakeDataProcessRegionViewCloudProvider_AZURE,
			"TENANT": DataLakeDataProcessRegionViewCloudProvider_TENANT,
			"SERVERLESS": DataLakeDataProcessRegionViewCloudProvider_SERVERLESS,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.DataLakeDataProcessRegionViewRegion",
		reflect.TypeOf((*DataLakeDataProcessRegionViewRegion)(nil)).Elem(),
		map[string]interface{}{
			"DUBLIN_IRL": DataLakeDataProcessRegionViewRegion_DUBLIN_IRL,
			"FRANKFURT_DEU": DataLakeDataProcessRegionViewRegion_FRANKFURT_DEU,
			"LONDON_GBR": DataLakeDataProcessRegionViewRegion_LONDON_GBR,
			"MUMBAI_IND": DataLakeDataProcessRegionViewRegion_MUMBAI_IND,
			"OREGON_USA": DataLakeDataProcessRegionViewRegion_OREGON_USA,
			"SYDNEY_AUS": DataLakeDataProcessRegionViewRegion_SYDNEY_AUS,
			"VIRGINIA_USA": DataLakeDataProcessRegionViewRegion_VIRGINIA_USA,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DataLakeDatabaseCollectionView",
		reflect.TypeOf((*DataLakeDatabaseCollectionView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DataLakeDatabaseDataSourceView",
		reflect.TypeOf((*DataLakeDatabaseDataSourceView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.DataLakeDatabaseDataSourceViewDefaultFormat",
		reflect.TypeOf((*DataLakeDatabaseDataSourceViewDefaultFormat)(nil)).Elem(),
		map[string]interface{}{
			"AVRO": DataLakeDatabaseDataSourceViewDefaultFormat_AVRO,
			"AVRO_GZ": DataLakeDatabaseDataSourceViewDefaultFormat_AVRO_GZ,
			"BSON": DataLakeDatabaseDataSourceViewDefaultFormat_BSON,
			"BSON_GZ": DataLakeDatabaseDataSourceViewDefaultFormat_BSON_GZ,
			"CSV": DataLakeDatabaseDataSourceViewDefaultFormat_CSV,
			"JSON": DataLakeDatabaseDataSourceViewDefaultFormat_JSON,
			"JSON_GZ": DataLakeDatabaseDataSourceViewDefaultFormat_JSON_GZ,
			"ORC": DataLakeDatabaseDataSourceViewDefaultFormat_ORC,
			"TSV": DataLakeDatabaseDataSourceViewDefaultFormat_TSV,
			"TSV_GZ": DataLakeDatabaseDataSourceViewDefaultFormat_TSV_GZ,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DataLakeDatabaseView",
		reflect.TypeOf((*DataLakeDatabaseView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DataLakeStorageView",
		reflect.TypeOf((*DataLakeStorageView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DataLakeViewView",
		reflect.TypeOf((*DataLakeViewView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DatabaseConfig",
		reflect.TypeOf((*DatabaseConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.DatabaseConfigOperationTypes",
		reflect.TypeOf((*DatabaseConfigOperationTypes)(nil)).Elem(),
		map[string]interface{}{
			"INSERT": DatabaseConfigOperationTypes_INSERT,
			"UPDATE": DatabaseConfigOperationTypes_UPDATE,
			"REPLACE": DatabaseConfigOperationTypes_REPLACE,
			"DELETE": DatabaseConfigOperationTypes_DELETE,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DatabaseUserProps",
		reflect.TypeOf((*DatabaseUserProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.DatadogIntegration",
		reflect.TypeOf((*DatadogIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cfnThirdPartyIntegration", GoGetter: "CfnThirdPartyIntegration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DatadogIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DatadogIntegrationProps",
		reflect.TypeOf((*DatadogIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.DatadogRegion",
		reflect.TypeOf((*DatadogRegion)(nil)).Elem(),
		map[string]interface{}{
			"US": DatadogRegion_US,
			"EU": DatadogRegion_EU,
			"US3": DatadogRegion_US3,
			"US5": DatadogRegion_US5,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DiskGb",
		reflect.TypeOf((*DiskGb)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.EncryptionAtRestProps",
		reflect.TypeOf((*EncryptionAtRestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Endpoint",
		reflect.TypeOf((*Endpoint)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Event",
		reflect.TypeOf((*Event)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.EventAwseventbridge",
		reflect.TypeOf((*EventAwseventbridge)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.EventAwseventbridgeAwsConfig",
		reflect.TypeOf((*EventAwseventbridgeAwsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.EventFunction",
		reflect.TypeOf((*EventFunction)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.EventFunctionFuncConfig",
		reflect.TypeOf((*EventFunctionFuncConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Export",
		reflect.TypeOf((*Export)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.InheritedRole",
		reflect.TypeOf((*InheritedRole)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.IntegerThresholdView",
		reflect.TypeOf((*IntegerThresholdView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.IntegerThresholdViewOperator",
		reflect.TypeOf((*IntegerThresholdViewOperator)(nil)).Elem(),
		map[string]interface{}{
			"GREATER_THAN": IntegerThresholdViewOperator_GREATER_THAN,
			"LESS_THAN": IntegerThresholdViewOperator_LESS_THAN,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.IntegerThresholdViewUnits",
		reflect.TypeOf((*IntegerThresholdViewUnits)(nil)).Elem(),
		map[string]interface{}{
			"BITS": IntegerThresholdViewUnits_BITS,
			"BYTES": IntegerThresholdViewUnits_BYTES,
			"DAYS": IntegerThresholdViewUnits_DAYS,
			"GIGABITS": IntegerThresholdViewUnits_GIGABITS,
			"GIGABYTES": IntegerThresholdViewUnits_GIGABYTES,
			"HOURS": IntegerThresholdViewUnits_HOURS,
			"KILOBITS": IntegerThresholdViewUnits_KILOBITS,
			"KILOBYTES": IntegerThresholdViewUnits_KILOBYTES,
			"MEGABITS": IntegerThresholdViewUnits_MEGABITS,
			"MEGABYTES": IntegerThresholdViewUnits_MEGABYTES,
			"MILLISECONDS": IntegerThresholdViewUnits_MILLISECONDS,
			"MINUTES": IntegerThresholdViewUnits_MINUTES,
			"PETABYTES": IntegerThresholdViewUnits_PETABYTES,
			"RAW": IntegerThresholdViewUnits_RAW,
			"SECONDS": IntegerThresholdViewUnits_SECONDS,
			"TERABYTES": IntegerThresholdViewUnits_TERABYTES,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.IpAccessListProps",
		reflect.TypeOf((*IpAccessListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.LabelDefinition",
		reflect.TypeOf((*LabelDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Link",
		reflect.TypeOf((*Link)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ListOptions",
		reflect.TypeOf((*ListOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ManagedNamespace",
		reflect.TypeOf((*ManagedNamespace)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Matcher",
		reflect.TypeOf((*Matcher)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.MatcherFieldName",
		reflect.TypeOf((*MatcherFieldName)(nil)).Elem(),
		map[string]interface{}{
			"CLUSTER_NAME": MatcherFieldName_CLUSTER_NAME,
			"HOSTNAME": MatcherFieldName_HOSTNAME,
			"HOSTNAME_AND_PORT": MatcherFieldName_HOSTNAME_AND_PORT,
			"PORT": MatcherFieldName_PORT,
			"REPLICA_SET_NAME": MatcherFieldName_REPLICA_SET_NAME,
			"SHARD_NAME": MatcherFieldName_SHARD_NAME,
			"TYPE_NAME": MatcherFieldName_TYPE_NAME,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.MatcherOperator",
		reflect.TypeOf((*MatcherOperator)(nil)).Elem(),
		map[string]interface{}{
			"EQUALS": MatcherOperator_EQUALS,
			"CONTAINS": MatcherOperator_CONTAINS,
			"STARTS_WITH": MatcherOperator_STARTS_WITH,
			"ENDS_WITH": MatcherOperator_ENDS_WITH,
			"NOT_EQUALS": MatcherOperator_NOT_EQUALS,
			"NOT_CONTAINS": MatcherOperator_NOT_CONTAINS,
			"REGEX": MatcherOperator_REGEX,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.MetricThresholdView",
		reflect.TypeOf((*MetricThresholdView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.MetricThresholdViewMetricName",
		reflect.TypeOf((*MetricThresholdViewMetricName)(nil)).Elem(),
		map[string]interface{}{
			"ASSERT_MSG": MetricThresholdViewMetricName_ASSERT_MSG,
			"ASSERT_REGULAR": MetricThresholdViewMetricName_ASSERT_REGULAR,
			"ASSERT_USER": MetricThresholdViewMetricName_ASSERT_USER,
			"ASSERT_WARNING": MetricThresholdViewMetricName_ASSERT_WARNING,
			"AVG_COMMAND_EXECUTION_TIME": MetricThresholdViewMetricName_AVG_COMMAND_EXECUTION_TIME,
			"AVG_READ_EXECUTION_TIME": MetricThresholdViewMetricName_AVG_READ_EXECUTION_TIME,
			"AVG_WRITE_EXECUTION_TIME": MetricThresholdViewMetricName_AVG_WRITE_EXECUTION_TIME,
			"BACKGROUND_FLUSH_AVG": MetricThresholdViewMetricName_BACKGROUND_FLUSH_AVG,
			"CACHE_BYTES_READ_INTO": MetricThresholdViewMetricName_CACHE_BYTES_READ_INTO,
			"CACHE_BYTES_WRITTEN_FROM": MetricThresholdViewMetricName_CACHE_BYTES_WRITTEN_FROM,
			"CACHE_USAGE_DIRTY": MetricThresholdViewMetricName_CACHE_USAGE_DIRTY,
			"CACHE_USAGE_USED": MetricThresholdViewMetricName_CACHE_USAGE_USED,
			"COMPUTED_MEMORY": MetricThresholdViewMetricName_COMPUTED_MEMORY,
			"CONNECTIONS": MetricThresholdViewMetricName_CONNECTIONS,
			"CONNECTIONS_MAX": MetricThresholdViewMetricName_CONNECTIONS_MAX,
			"CONNECTIONS_PERCENT": MetricThresholdViewMetricName_CONNECTIONS_PERCENT,
			"CURSORS_TOTAL_CLIENT_CURSORS_SIZE": MetricThresholdViewMetricName_CURSORS_TOTAL_CLIENT_CURSORS_SIZE,
			"CURSORS_TOTAL_OPEN": MetricThresholdViewMetricName_CURSORS_TOTAL_OPEN,
			"CURSORS_TOTAL_TIMED_OUT": MetricThresholdViewMetricName_CURSORS_TOTAL_TIMED_OUT,
			"DB_DATA_SIZE_TOTAL": MetricThresholdViewMetricName_DB_DATA_SIZE_TOTAL,
			"DB_INDEX_SIZE_TOTAL": MetricThresholdViewMetricName_DB_INDEX_SIZE_TOTAL,
			"DB_STORAGE_TOTAL": MetricThresholdViewMetricName_DB_STORAGE_TOTAL,
			"DISK_PARTITION_SPACE_USED_DATA": MetricThresholdViewMetricName_DISK_PARTITION_SPACE_USED_DATA,
			"DISK_PARTITION_SPACE_USED_INDEX": MetricThresholdViewMetricName_DISK_PARTITION_SPACE_USED_INDEX,
			"DISK_PARTITION_SPACE_USED_JOURNAL": MetricThresholdViewMetricName_DISK_PARTITION_SPACE_USED_JOURNAL,
			"DISK_PARTITION_UTILIZATION_DATA": MetricThresholdViewMetricName_DISK_PARTITION_UTILIZATION_DATA,
			"DISK_PARTITION_UTILIZATION_INDEX": MetricThresholdViewMetricName_DISK_PARTITION_UTILIZATION_INDEX,
			"DISK_PARTITION_UTILIZATION_JOURNAL": MetricThresholdViewMetricName_DISK_PARTITION_UTILIZATION_JOURNAL,
			"DOCUMENT_DELETED": MetricThresholdViewMetricName_DOCUMENT_DELETED,
			"DOCUMENT_INSERTED": MetricThresholdViewMetricName_DOCUMENT_INSERTED,
			"DOCUMENT_RETURNED": MetricThresholdViewMetricName_DOCUMENT_RETURNED,
			"DOCUMENT_UPDATED": MetricThresholdViewMetricName_DOCUMENT_UPDATED,
			"EXTRA_INFO_PAGE_FAULTS": MetricThresholdViewMetricName_EXTRA_INFO_PAGE_FAULTS,
			"FTS_DISK_UTILIZATION": MetricThresholdViewMetricName_FTS_DISK_UTILIZATION,
			"FTS_MEMORY_MAPPED": MetricThresholdViewMetricName_FTS_MEMORY_MAPPED,
			"FTS_MEMORY_RESIDENT": MetricThresholdViewMetricName_FTS_MEMORY_RESIDENT,
			"FTS_MEMORY_VIRTUAL": MetricThresholdViewMetricName_FTS_MEMORY_VIRTUAL,
			"FTS_PROCESS_CPU_KERNEL": MetricThresholdViewMetricName_FTS_PROCESS_CPU_KERNEL,
			"FTS_PROCESS_CPU_USER": MetricThresholdViewMetricName_FTS_PROCESS_CPU_USER,
			"GLOBAL_ACCESSES_NOT_IN_MEMORY": MetricThresholdViewMetricName_GLOBAL_ACCESSES_NOT_IN_MEMORY,
			"GLOBAL_LOCK_CURRENT_QUEUE_READERS": MetricThresholdViewMetricName_GLOBAL_LOCK_CURRENT_QUEUE_READERS,
			"GLOBAL_LOCK_CURRENT_QUEUE_TOTAL": MetricThresholdViewMetricName_GLOBAL_LOCK_CURRENT_QUEUE_TOTAL,
			"GLOBAL_LOCK_CURRENT_QUEUE_WRITERS": MetricThresholdViewMetricName_GLOBAL_LOCK_CURRENT_QUEUE_WRITERS,
			"GLOBAL_LOCK_PERCENTAGE": MetricThresholdViewMetricName_GLOBAL_LOCK_PERCENTAGE,
			"GLOBAL_PAGE_FAULT_EXCEPTIONS_THROWN": MetricThresholdViewMetricName_GLOBAL_PAGE_FAULT_EXCEPTIONS_THROWN,
			"INDEX_COUNTERS_BTREE_ACCESSES": MetricThresholdViewMetricName_INDEX_COUNTERS_BTREE_ACCESSES,
			"INDEX_COUNTERS_BTREE_HITS": MetricThresholdViewMetricName_INDEX_COUNTERS_BTREE_HITS,
			"INDEX_COUNTERS_BTREE_MISS_RATIO": MetricThresholdViewMetricName_INDEX_COUNTERS_BTREE_MISS_RATIO,
			"INDEX_COUNTERS_BTREE_MISSES": MetricThresholdViewMetricName_INDEX_COUNTERS_BTREE_MISSES,
			"JOURNALING_COMMITS_IN_WRITE_LOCK": MetricThresholdViewMetricName_JOURNALING_COMMITS_IN_WRITE_LOCK,
			"JOURNALING_MB": MetricThresholdViewMetricName_JOURNALING_MB,
			"JOURNALING_WRITE_DATA_FILES_MB": MetricThresholdViewMetricName_JOURNALING_WRITE_DATA_FILES_MB,
			"LOGICAL_SIZE": MetricThresholdViewMetricName_LOGICAL_SIZE,
			"MEMORY_MAPPED": MetricThresholdViewMetricName_MEMORY_MAPPED,
			"MEMORY_RESIDENT": MetricThresholdViewMetricName_MEMORY_RESIDENT,
			"MEMORY_VIRTUAL": MetricThresholdViewMetricName_MEMORY_VIRTUAL,
			"MUNIN_CPU_IOWAIT": MetricThresholdViewMetricName_MUNIN_CPU_IOWAIT,
			"MUNIN_CPU_IRQ": MetricThresholdViewMetricName_MUNIN_CPU_IRQ,
			"MUNIN_CPU_NICE": MetricThresholdViewMetricName_MUNIN_CPU_NICE,
			"MUNIN_CPU_SOFTIRQ": MetricThresholdViewMetricName_MUNIN_CPU_SOFTIRQ,
			"MUNIN_CPU_STEAL": MetricThresholdViewMetricName_MUNIN_CPU_STEAL,
			"MUNIN_CPU_SYSTEM": MetricThresholdViewMetricName_MUNIN_CPU_SYSTEM,
			"MUNIN_CPU_USER": MetricThresholdViewMetricName_MUNIN_CPU_USER,
			"NETWORK_BYTES_IN": MetricThresholdViewMetricName_NETWORK_BYTES_IN,
			"NETWORK_BYTES_OUT": MetricThresholdViewMetricName_NETWORK_BYTES_OUT,
			"NETWORK_NUM_REQUESTS": MetricThresholdViewMetricName_NETWORK_NUM_REQUESTS,
			"NORMALIZED_FTS_PROCESS_CPU_KERNEL": MetricThresholdViewMetricName_NORMALIZED_FTS_PROCESS_CPU_KERNEL,
			"NORMALIZED_FTS_PROCESS_CPU_USER": MetricThresholdViewMetricName_NORMALIZED_FTS_PROCESS_CPU_USER,
			"NORMALIZED_SYSTEM_CPU_STEAL": MetricThresholdViewMetricName_NORMALIZED_SYSTEM_CPU_STEAL,
			"NORMALIZED_SYSTEM_CPU_USER": MetricThresholdViewMetricName_NORMALIZED_SYSTEM_CPU_USER,
			"OPCOUNTER_CMD": MetricThresholdViewMetricName_OPCOUNTER_CMD,
			"OPCOUNTER_DELETE": MetricThresholdViewMetricName_OPCOUNTER_DELETE,
			"OPCOUNTER_GETMORE": MetricThresholdViewMetricName_OPCOUNTER_GETMORE,
			"OPCOUNTER_INSERT": MetricThresholdViewMetricName_OPCOUNTER_INSERT,
			"OPCOUNTER_QUERY": MetricThresholdViewMetricName_OPCOUNTER_QUERY,
			"OPCOUNTER_REPL_CMD": MetricThresholdViewMetricName_OPCOUNTER_REPL_CMD,
			"OPCOUNTER_REPL_DELETE": MetricThresholdViewMetricName_OPCOUNTER_REPL_DELETE,
			"OPCOUNTER_REPL_INSERT": MetricThresholdViewMetricName_OPCOUNTER_REPL_INSERT,
			"OPCOUNTER_REPL_UPDATE": MetricThresholdViewMetricName_OPCOUNTER_REPL_UPDATE,
			"OPCOUNTER_UPDATE": MetricThresholdViewMetricName_OPCOUNTER_UPDATE,
			"OPERATIONS_SCAN_AND_ORDER": MetricThresholdViewMetricName_OPERATIONS_SCAN_AND_ORDER,
			"OPLOG_MASTER_LAG_TIME_DIFF": MetricThresholdViewMetricName_OPLOG_MASTER_LAG_TIME_DIFF,
			"OPLOG_MASTER_TIME": MetricThresholdViewMetricName_OPLOG_MASTER_TIME,
			"OPLOG_MASTER_TIME_ESTIMATED_TTL": MetricThresholdViewMetricName_OPLOG_MASTER_TIME_ESTIMATED_TTL,
			"OPLOG_RATE_GB_PER_HOUR": MetricThresholdViewMetricName_OPLOG_RATE_GB_PER_HOUR,
			"OPLOG_SLAVE_LAG_MASTER_TIME": MetricThresholdViewMetricName_OPLOG_SLAVE_LAG_MASTER_TIME,
			"QUERY_EXECUTOR_SCANNED": MetricThresholdViewMetricName_QUERY_EXECUTOR_SCANNED,
			"QUERY_EXECUTOR_SCANNED_OBJECTS": MetricThresholdViewMetricName_QUERY_EXECUTOR_SCANNED_OBJECTS,
			"QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED": MetricThresholdViewMetricName_QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED,
			"QUERY_TARGETING_SCANNED_PER_RETURNED": MetricThresholdViewMetricName_QUERY_TARGETING_SCANNED_PER_RETURNED,
			"RESTARTS_IN_LAST_HOUR": MetricThresholdViewMetricName_RESTARTS_IN_LAST_HOUR,
			"SERVERLESS_CONNECTIONS": MetricThresholdViewMetricName_SERVERLESS_CONNECTIONS,
			"SERVERLESS_CONNECTIONS_PERCENT": MetricThresholdViewMetricName_SERVERLESS_CONNECTIONS_PERCENT,
			"SERVERLESS_DATA_SIZE_TOTAL": MetricThresholdViewMetricName_SERVERLESS_DATA_SIZE_TOTAL,
			"SERVERLESS_NETWORK_BYTES_IN": MetricThresholdViewMetricName_SERVERLESS_NETWORK_BYTES_IN,
			"SERVERLESS_NETWORK_BYTES_OUT": MetricThresholdViewMetricName_SERVERLESS_NETWORK_BYTES_OUT,
			"SERVERLESS_NETWORK_NUM_REQUESTS": MetricThresholdViewMetricName_SERVERLESS_NETWORK_NUM_REQUESTS,
			"SERVERLESS_OPCOUNTER_CMD": MetricThresholdViewMetricName_SERVERLESS_OPCOUNTER_CMD,
			"SERVERLESS_OPCOUNTER_DELETE": MetricThresholdViewMetricName_SERVERLESS_OPCOUNTER_DELETE,
			"SERVERLESS_OPCOUNTER_GETMORE": MetricThresholdViewMetricName_SERVERLESS_OPCOUNTER_GETMORE,
			"SERVERLESS_OPCOUNTER_INSERT": MetricThresholdViewMetricName_SERVERLESS_OPCOUNTER_INSERT,
			"SERVERLESS_OPCOUNTER_QUERY": MetricThresholdViewMetricName_SERVERLESS_OPCOUNTER_QUERY,
			"SERVERLESS_OPCOUNTER_UPDATE": MetricThresholdViewMetricName_SERVERLESS_OPCOUNTER_UPDATE,
			"SERVERLESS_TOTAL_READ_UNITS": MetricThresholdViewMetricName_SERVERLESS_TOTAL_READ_UNITS,
			"SERVERLESS_TOTAL_WRITE_UNITS": MetricThresholdViewMetricName_SERVERLESS_TOTAL_WRITE_UNITS,
			"SWAP_USAGE_FREE": MetricThresholdViewMetricName_SWAP_USAGE_FREE,
			"SWAP_USAGE_USED": MetricThresholdViewMetricName_SWAP_USAGE_USED,
			"SYSTEM_MEMORY_AVAILABLE": MetricThresholdViewMetricName_SYSTEM_MEMORY_AVAILABLE,
			"SYSTEM_MEMORY_FREE": MetricThresholdViewMetricName_SYSTEM_MEMORY_FREE,
			"SYSTEM_MEMORY_USED": MetricThresholdViewMetricName_SYSTEM_MEMORY_USED,
			"SYSTEM_NETWORK_IN": MetricThresholdViewMetricName_SYSTEM_NETWORK_IN,
			"SYSTEM_NETWORK_OUT": MetricThresholdViewMetricName_SYSTEM_NETWORK_OUT,
			"TICKETS_AVAILABLE_READS": MetricThresholdViewMetricName_TICKETS_AVAILABLE_READS,
			"TICKETS_AVAILABLE_WRITES": MetricThresholdViewMetricName_TICKETS_AVAILABLE_WRITES,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.MetricThresholdViewMode",
		reflect.TypeOf((*MetricThresholdViewMode)(nil)).Elem(),
		map[string]interface{}{
			"AVERAGE": MetricThresholdViewMode_AVERAGE,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.MetricThresholdViewOperator",
		reflect.TypeOf((*MetricThresholdViewOperator)(nil)).Elem(),
		map[string]interface{}{
			"GREATER_THAN": MetricThresholdViewOperator_GREATER_THAN,
			"LESS_THAN": MetricThresholdViewOperator_LESS_THAN,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.MetricThresholdViewUnits",
		reflect.TypeOf((*MetricThresholdViewUnits)(nil)).Elem(),
		map[string]interface{}{
			"BITS": MetricThresholdViewUnits_BITS,
			"BYTES": MetricThresholdViewUnits_BYTES,
			"DAYS": MetricThresholdViewUnits_DAYS,
			"GIGABITS": MetricThresholdViewUnits_GIGABITS,
			"GIGABYTES": MetricThresholdViewUnits_GIGABYTES,
			"HOURS": MetricThresholdViewUnits_HOURS,
			"KILOBITS": MetricThresholdViewUnits_KILOBITS,
			"KILOBYTES": MetricThresholdViewUnits_KILOBYTES,
			"MEGABITS": MetricThresholdViewUnits_MEGABITS,
			"MEGABYTES": MetricThresholdViewUnits_MEGABYTES,
			"MILLISECONDS": MetricThresholdViewUnits_MILLISECONDS,
			"MINUTES": MetricThresholdViewUnits_MINUTES,
			"PETABYTES": MetricThresholdViewUnits_PETABYTES,
			"RAW": MetricThresholdViewUnits_RAW,
			"SECONDS": MetricThresholdViewUnits_SECONDS,
			"TERABYTES": MetricThresholdViewUnits_TERABYTES,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.MicrosoftTeamsIntegration",
		reflect.TypeOf((*MicrosoftTeamsIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cfnThirdPartyIntegration", GoGetter: "CfnThirdPartyIntegration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MicrosoftTeamsIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.MicrosoftTeamsIntegrationProps",
		reflect.TypeOf((*MicrosoftTeamsIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.NotificationView",
		reflect.TypeOf((*NotificationView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.NotificationViewDatadogRegion",
		reflect.TypeOf((*NotificationViewDatadogRegion)(nil)).Elem(),
		map[string]interface{}{
			"EU": NotificationViewDatadogRegion_EU,
			"US": NotificationViewDatadogRegion_US,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.NotificationViewOpsGenieRegion",
		reflect.TypeOf((*NotificationViewOpsGenieRegion)(nil)).Elem(),
		map[string]interface{}{
			"EU": NotificationViewOpsGenieRegion_EU,
			"US": NotificationViewOpsGenieRegion_US,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.NotificationViewRoles",
		reflect.TypeOf((*NotificationViewRoles)(nil)).Elem(),
		map[string]interface{}{
			"GROUP_CLUSTER_MANAGER": NotificationViewRoles_GROUP_CLUSTER_MANAGER,
			"GROUP_DATA_ACCESS_ADMIN": NotificationViewRoles_GROUP_DATA_ACCESS_ADMIN,
			"GROUP_DATA_ACCESS_READ_ONLY": NotificationViewRoles_GROUP_DATA_ACCESS_READ_ONLY,
			"GROUP_DATA_ACCESS_READ_WRITE": NotificationViewRoles_GROUP_DATA_ACCESS_READ_WRITE,
			"GROUP_OWNER": NotificationViewRoles_GROUP_OWNER,
			"GROUP_READ_WRITE": NotificationViewRoles_GROUP_READ_WRITE,
			"ORG_OWNER": NotificationViewRoles_ORG_OWNER,
			"ORG_MEMBER": NotificationViewRoles_ORG_MEMBER,
			"ORG_GROUP_CREATOR": NotificationViewRoles_ORG_GROUP_CREATOR,
			"ORG_BILLING_ADMIN": NotificationViewRoles_ORG_BILLING_ADMIN,
			"ORG_READ_ONLY": NotificationViewRoles_ORG_READ_ONLY,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.NotificationViewSeverity",
		reflect.TypeOf((*NotificationViewSeverity)(nil)).Elem(),
		map[string]interface{}{
			"CRITICAL": NotificationViewSeverity_CRITICAL,
			"ERROR": NotificationViewSeverity_ERROR,
			"WARNING": NotificationViewSeverity_WARNING,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.NotificationViewTypeName",
		reflect.TypeOf((*NotificationViewTypeName)(nil)).Elem(),
		map[string]interface{}{
			"DATADOG": NotificationViewTypeName_DATADOG,
			"EMAIL": NotificationViewTypeName_EMAIL,
			"FLOWDOCK": NotificationViewTypeName_FLOWDOCK,
			"GROUP": NotificationViewTypeName_GROUP,
			"MICROSOFT_TEAMS": NotificationViewTypeName_MICROSOFT_TEAMS,
			"OPS_GENIE": NotificationViewTypeName_OPS_GENIE,
			"ORG": NotificationViewTypeName_ORG,
			"PAGER_DUTY": NotificationViewTypeName_PAGER_DUTY,
			"PROMETHEUS": NotificationViewTypeName_PROMETHEUS,
			"SLACK": NotificationViewTypeName_SLACK,
			"SMS": NotificationViewTypeName_SMS,
			"TEAM": NotificationViewTypeName_TEAM,
			"USER": NotificationViewTypeName_USER,
			"VICTOR_OPS": NotificationViewTypeName_VICTOR_OPS,
			"WEBHOOK": NotificationViewTypeName_WEBHOOK,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.PagerDutyIntegration",
		reflect.TypeOf((*PagerDutyIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cfnThirdPartyIntegration", GoGetter: "CfnThirdPartyIntegration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PagerDutyIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.PagerDutyIntegrationProps",
		reflect.TypeOf((*PagerDutyIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.PagerDutyRegion",
		reflect.TypeOf((*PagerDutyRegion)(nil)).Elem(),
		map[string]interface{}{
			"US": PagerDutyRegion_US,
			"EU": PagerDutyRegion_EU,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.PartitionFieldView",
		reflect.TypeOf((*PartitionFieldView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.PartitionFieldViewFieldType",
		reflect.TypeOf((*PartitionFieldViewFieldType)(nil)).Elem(),
		map[string]interface{}{
			"DATE": PartitionFieldViewFieldType_DATE,
			"INT": PartitionFieldViewFieldType_INT,
			"LONG": PartitionFieldViewFieldType_LONG,
			"OBJECT_ID": PartitionFieldViewFieldType_OBJECT_ID,
			"STRING": PartitionFieldViewFieldType_STRING,
			"UUID": PartitionFieldViewFieldType_UUID,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.PrivateEndpoint",
		reflect.TypeOf((*PrivateEndpoint)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.PrivateEndpointProps",
		reflect.TypeOf((*PrivateEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ProcessArgs",
		reflect.TypeOf((*ProcessArgs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ProjectApiKey",
		reflect.TypeOf((*ProjectApiKey)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ProjectProps",
		reflect.TypeOf((*ProjectProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ProjectSettings",
		reflect.TypeOf((*ProjectSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ProjectTeam",
		reflect.TypeOf((*ProjectTeam)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Resource",
		reflect.TypeOf((*Resource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.RoleAssignment",
		reflect.TypeOf((*RoleAssignment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.RoleDefinition",
		reflect.TypeOf((*RoleDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ScheduleConfig",
		reflect.TypeOf((*ScheduleConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ScheduleView",
		reflect.TypeOf((*ScheduleView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.ScheduleViewType",
		reflect.TypeOf((*ScheduleViewType)(nil)).Elem(),
		map[string]interface{}{
			"DAILY": ScheduleViewType_DAILY,
			"MONTHLY": ScheduleViewType_MONTHLY,
			"DEFAULT": ScheduleViewType_DEFAULT,
			"WEEKLY": ScheduleViewType_WEEKLY,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ScopeDefinition",
		reflect.TypeOf((*ScopeDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.ScopeDefinitionType",
		reflect.TypeOf((*ScopeDefinitionType)(nil)).Elem(),
		map[string]interface{}{
			"CLUSTER": ScopeDefinitionType_CLUSTER,
			"DATA_LAKE": ScopeDefinitionType_DATA_LAKE,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ServerlessInstanceConnectionStrings",
		reflect.TypeOf((*ServerlessInstanceConnectionStrings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ServerlessInstancePrivateEndpoint",
		reflect.TypeOf((*ServerlessInstancePrivateEndpoint)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ServerlessInstancePrivateEndpointEndpoint",
		reflect.TypeOf((*ServerlessInstancePrivateEndpointEndpoint)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.ServerlessInstancePrivateEndpointType",
		reflect.TypeOf((*ServerlessInstancePrivateEndpointType)(nil)).Elem(),
		map[string]interface{}{
			"MONGOS": ServerlessInstancePrivateEndpointType_MONGOS,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ServerlessInstanceProviderSettings",
		reflect.TypeOf((*ServerlessInstanceProviderSettings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.ServerlessInstanceProviderSettingsProviderName",
		reflect.TypeOf((*ServerlessInstanceProviderSettingsProviderName)(nil)).Elem(),
		map[string]interface{}{
			"SERVERLESS": ServerlessInstanceProviderSettingsProviderName_SERVERLESS,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Specs",
		reflect.TypeOf((*Specs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.StoreDetail",
		reflect.TypeOf((*StoreDetail)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.SynonymSource",
		reflect.TypeOf((*SynonymSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ThirdPartyIntegrationProps",
		reflect.TypeOf((*ThirdPartyIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Validation",
		reflect.TypeOf((*Validation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ZoneMapping",
		reflect.TypeOf((*ZoneMapping)(nil)).Elem(),
	)
}
