// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


// Human-readable label that identifies the metric against which MongoDB Cloud checks the configured **metricThreshold.threshold**.
type MetricThresholdViewMetricName string

const (
	// ASSERT_MSG.
	MetricThresholdViewMetricName_ASSERT_MSG MetricThresholdViewMetricName = "ASSERT_MSG"
	// ASSERT_REGULAR.
	MetricThresholdViewMetricName_ASSERT_REGULAR MetricThresholdViewMetricName = "ASSERT_REGULAR"
	// ASSERT_USER.
	MetricThresholdViewMetricName_ASSERT_USER MetricThresholdViewMetricName = "ASSERT_USER"
	// ASSERT_WARNING.
	MetricThresholdViewMetricName_ASSERT_WARNING MetricThresholdViewMetricName = "ASSERT_WARNING"
	// AVG_COMMAND_EXECUTION_TIME.
	MetricThresholdViewMetricName_AVG_COMMAND_EXECUTION_TIME MetricThresholdViewMetricName = "AVG_COMMAND_EXECUTION_TIME"
	// AVG_READ_EXECUTION_TIME.
	MetricThresholdViewMetricName_AVG_READ_EXECUTION_TIME MetricThresholdViewMetricName = "AVG_READ_EXECUTION_TIME"
	// AVG_WRITE_EXECUTION_TIME.
	MetricThresholdViewMetricName_AVG_WRITE_EXECUTION_TIME MetricThresholdViewMetricName = "AVG_WRITE_EXECUTION_TIME"
	// BACKGROUND_FLUSH_AVG.
	MetricThresholdViewMetricName_BACKGROUND_FLUSH_AVG MetricThresholdViewMetricName = "BACKGROUND_FLUSH_AVG"
	// CACHE_BYTES_READ_INTO.
	MetricThresholdViewMetricName_CACHE_BYTES_READ_INTO MetricThresholdViewMetricName = "CACHE_BYTES_READ_INTO"
	// CACHE_BYTES_WRITTEN_FROM.
	MetricThresholdViewMetricName_CACHE_BYTES_WRITTEN_FROM MetricThresholdViewMetricName = "CACHE_BYTES_WRITTEN_FROM"
	// CACHE_USAGE_DIRTY.
	MetricThresholdViewMetricName_CACHE_USAGE_DIRTY MetricThresholdViewMetricName = "CACHE_USAGE_DIRTY"
	// CACHE_USAGE_USED.
	MetricThresholdViewMetricName_CACHE_USAGE_USED MetricThresholdViewMetricName = "CACHE_USAGE_USED"
	// COMPUTED_MEMORY.
	MetricThresholdViewMetricName_COMPUTED_MEMORY MetricThresholdViewMetricName = "COMPUTED_MEMORY"
	// CONNECTIONS.
	MetricThresholdViewMetricName_CONNECTIONS MetricThresholdViewMetricName = "CONNECTIONS"
	// CONNECTIONS_MAX.
	MetricThresholdViewMetricName_CONNECTIONS_MAX MetricThresholdViewMetricName = "CONNECTIONS_MAX"
	// CONNECTIONS_PERCENT.
	MetricThresholdViewMetricName_CONNECTIONS_PERCENT MetricThresholdViewMetricName = "CONNECTIONS_PERCENT"
	// CURSORS_TOTAL_CLIENT_CURSORS_SIZE.
	MetricThresholdViewMetricName_CURSORS_TOTAL_CLIENT_CURSORS_SIZE MetricThresholdViewMetricName = "CURSORS_TOTAL_CLIENT_CURSORS_SIZE"
	// CURSORS_TOTAL_OPEN.
	MetricThresholdViewMetricName_CURSORS_TOTAL_OPEN MetricThresholdViewMetricName = "CURSORS_TOTAL_OPEN"
	// CURSORS_TOTAL_TIMED_OUT.
	MetricThresholdViewMetricName_CURSORS_TOTAL_TIMED_OUT MetricThresholdViewMetricName = "CURSORS_TOTAL_TIMED_OUT"
	// DB_DATA_SIZE_TOTAL.
	MetricThresholdViewMetricName_DB_DATA_SIZE_TOTAL MetricThresholdViewMetricName = "DB_DATA_SIZE_TOTAL"
	// DB_INDEX_SIZE_TOTAL.
	MetricThresholdViewMetricName_DB_INDEX_SIZE_TOTAL MetricThresholdViewMetricName = "DB_INDEX_SIZE_TOTAL"
	// DB_STORAGE_TOTAL.
	MetricThresholdViewMetricName_DB_STORAGE_TOTAL MetricThresholdViewMetricName = "DB_STORAGE_TOTAL"
	// DISK_PARTITION_SPACE_USED_DATA.
	MetricThresholdViewMetricName_DISK_PARTITION_SPACE_USED_DATA MetricThresholdViewMetricName = "DISK_PARTITION_SPACE_USED_DATA"
	// DISK_PARTITION_SPACE_USED_INDEX.
	MetricThresholdViewMetricName_DISK_PARTITION_SPACE_USED_INDEX MetricThresholdViewMetricName = "DISK_PARTITION_SPACE_USED_INDEX"
	// DISK_PARTITION_SPACE_USED_JOURNAL.
	MetricThresholdViewMetricName_DISK_PARTITION_SPACE_USED_JOURNAL MetricThresholdViewMetricName = "DISK_PARTITION_SPACE_USED_JOURNAL"
	// DISK_PARTITION_UTILIZATION_DATA.
	MetricThresholdViewMetricName_DISK_PARTITION_UTILIZATION_DATA MetricThresholdViewMetricName = "DISK_PARTITION_UTILIZATION_DATA"
	// DISK_PARTITION_UTILIZATION_INDEX.
	MetricThresholdViewMetricName_DISK_PARTITION_UTILIZATION_INDEX MetricThresholdViewMetricName = "DISK_PARTITION_UTILIZATION_INDEX"
	// DISK_PARTITION_UTILIZATION_JOURNAL.
	MetricThresholdViewMetricName_DISK_PARTITION_UTILIZATION_JOURNAL MetricThresholdViewMetricName = "DISK_PARTITION_UTILIZATION_JOURNAL"
	// DOCUMENT_DELETED.
	MetricThresholdViewMetricName_DOCUMENT_DELETED MetricThresholdViewMetricName = "DOCUMENT_DELETED"
	// DOCUMENT_INSERTED.
	MetricThresholdViewMetricName_DOCUMENT_INSERTED MetricThresholdViewMetricName = "DOCUMENT_INSERTED"
	// DOCUMENT_RETURNED.
	MetricThresholdViewMetricName_DOCUMENT_RETURNED MetricThresholdViewMetricName = "DOCUMENT_RETURNED"
	// DOCUMENT_UPDATED.
	MetricThresholdViewMetricName_DOCUMENT_UPDATED MetricThresholdViewMetricName = "DOCUMENT_UPDATED"
	// EXTRA_INFO_PAGE_FAULTS.
	MetricThresholdViewMetricName_EXTRA_INFO_PAGE_FAULTS MetricThresholdViewMetricName = "EXTRA_INFO_PAGE_FAULTS"
	// FTS_DISK_UTILIZATION.
	MetricThresholdViewMetricName_FTS_DISK_UTILIZATION MetricThresholdViewMetricName = "FTS_DISK_UTILIZATION"
	// FTS_MEMORY_MAPPED.
	MetricThresholdViewMetricName_FTS_MEMORY_MAPPED MetricThresholdViewMetricName = "FTS_MEMORY_MAPPED"
	// FTS_MEMORY_RESIDENT.
	MetricThresholdViewMetricName_FTS_MEMORY_RESIDENT MetricThresholdViewMetricName = "FTS_MEMORY_RESIDENT"
	// FTS_MEMORY_VIRTUAL.
	MetricThresholdViewMetricName_FTS_MEMORY_VIRTUAL MetricThresholdViewMetricName = "FTS_MEMORY_VIRTUAL"
	// FTS_PROCESS_CPU_KERNEL.
	MetricThresholdViewMetricName_FTS_PROCESS_CPU_KERNEL MetricThresholdViewMetricName = "FTS_PROCESS_CPU_KERNEL"
	// FTS_PROCESS_CPU_USER.
	MetricThresholdViewMetricName_FTS_PROCESS_CPU_USER MetricThresholdViewMetricName = "FTS_PROCESS_CPU_USER"
	// GLOBAL_ACCESSES_NOT_IN_MEMORY.
	MetricThresholdViewMetricName_GLOBAL_ACCESSES_NOT_IN_MEMORY MetricThresholdViewMetricName = "GLOBAL_ACCESSES_NOT_IN_MEMORY"
	// GLOBAL_LOCK_CURRENT_QUEUE_READERS.
	MetricThresholdViewMetricName_GLOBAL_LOCK_CURRENT_QUEUE_READERS MetricThresholdViewMetricName = "GLOBAL_LOCK_CURRENT_QUEUE_READERS"
	// GLOBAL_LOCK_CURRENT_QUEUE_TOTAL.
	MetricThresholdViewMetricName_GLOBAL_LOCK_CURRENT_QUEUE_TOTAL MetricThresholdViewMetricName = "GLOBAL_LOCK_CURRENT_QUEUE_TOTAL"
	// GLOBAL_LOCK_CURRENT_QUEUE_WRITERS.
	MetricThresholdViewMetricName_GLOBAL_LOCK_CURRENT_QUEUE_WRITERS MetricThresholdViewMetricName = "GLOBAL_LOCK_CURRENT_QUEUE_WRITERS"
	// GLOBAL_LOCK_PERCENTAGE.
	MetricThresholdViewMetricName_GLOBAL_LOCK_PERCENTAGE MetricThresholdViewMetricName = "GLOBAL_LOCK_PERCENTAGE"
	// GLOBAL_PAGE_FAULT_EXCEPTIONS_THROWN.
	MetricThresholdViewMetricName_GLOBAL_PAGE_FAULT_EXCEPTIONS_THROWN MetricThresholdViewMetricName = "GLOBAL_PAGE_FAULT_EXCEPTIONS_THROWN"
	// INDEX_COUNTERS_BTREE_ACCESSES.
	MetricThresholdViewMetricName_INDEX_COUNTERS_BTREE_ACCESSES MetricThresholdViewMetricName = "INDEX_COUNTERS_BTREE_ACCESSES"
	// INDEX_COUNTERS_BTREE_HITS.
	MetricThresholdViewMetricName_INDEX_COUNTERS_BTREE_HITS MetricThresholdViewMetricName = "INDEX_COUNTERS_BTREE_HITS"
	// INDEX_COUNTERS_BTREE_MISS_RATIO.
	MetricThresholdViewMetricName_INDEX_COUNTERS_BTREE_MISS_RATIO MetricThresholdViewMetricName = "INDEX_COUNTERS_BTREE_MISS_RATIO"
	// INDEX_COUNTERS_BTREE_MISSES.
	MetricThresholdViewMetricName_INDEX_COUNTERS_BTREE_MISSES MetricThresholdViewMetricName = "INDEX_COUNTERS_BTREE_MISSES"
	// JOURNALING_COMMITS_IN_WRITE_LOCK.
	MetricThresholdViewMetricName_JOURNALING_COMMITS_IN_WRITE_LOCK MetricThresholdViewMetricName = "JOURNALING_COMMITS_IN_WRITE_LOCK"
	// JOURNALING_MB.
	MetricThresholdViewMetricName_JOURNALING_MB MetricThresholdViewMetricName = "JOURNALING_MB"
	// JOURNALING_WRITE_DATA_FILES_MB.
	MetricThresholdViewMetricName_JOURNALING_WRITE_DATA_FILES_MB MetricThresholdViewMetricName = "JOURNALING_WRITE_DATA_FILES_MB"
	// LOGICAL_SIZE.
	MetricThresholdViewMetricName_LOGICAL_SIZE MetricThresholdViewMetricName = "LOGICAL_SIZE"
	// MEMORY_MAPPED.
	MetricThresholdViewMetricName_MEMORY_MAPPED MetricThresholdViewMetricName = "MEMORY_MAPPED"
	// MEMORY_RESIDENT.
	MetricThresholdViewMetricName_MEMORY_RESIDENT MetricThresholdViewMetricName = "MEMORY_RESIDENT"
	// MEMORY_VIRTUAL.
	MetricThresholdViewMetricName_MEMORY_VIRTUAL MetricThresholdViewMetricName = "MEMORY_VIRTUAL"
	// MUNIN_CPU_IOWAIT.
	MetricThresholdViewMetricName_MUNIN_CPU_IOWAIT MetricThresholdViewMetricName = "MUNIN_CPU_IOWAIT"
	// MUNIN_CPU_IRQ.
	MetricThresholdViewMetricName_MUNIN_CPU_IRQ MetricThresholdViewMetricName = "MUNIN_CPU_IRQ"
	// MUNIN_CPU_NICE.
	MetricThresholdViewMetricName_MUNIN_CPU_NICE MetricThresholdViewMetricName = "MUNIN_CPU_NICE"
	// MUNIN_CPU_SOFTIRQ.
	MetricThresholdViewMetricName_MUNIN_CPU_SOFTIRQ MetricThresholdViewMetricName = "MUNIN_CPU_SOFTIRQ"
	// MUNIN_CPU_STEAL.
	MetricThresholdViewMetricName_MUNIN_CPU_STEAL MetricThresholdViewMetricName = "MUNIN_CPU_STEAL"
	// MUNIN_CPU_SYSTEM.
	MetricThresholdViewMetricName_MUNIN_CPU_SYSTEM MetricThresholdViewMetricName = "MUNIN_CPU_SYSTEM"
	// MUNIN_CPU_USER.
	MetricThresholdViewMetricName_MUNIN_CPU_USER MetricThresholdViewMetricName = "MUNIN_CPU_USER"
	// NETWORK_BYTES_IN.
	MetricThresholdViewMetricName_NETWORK_BYTES_IN MetricThresholdViewMetricName = "NETWORK_BYTES_IN"
	// NETWORK_BYTES_OUT.
	MetricThresholdViewMetricName_NETWORK_BYTES_OUT MetricThresholdViewMetricName = "NETWORK_BYTES_OUT"
	// NETWORK_NUM_REQUESTS.
	MetricThresholdViewMetricName_NETWORK_NUM_REQUESTS MetricThresholdViewMetricName = "NETWORK_NUM_REQUESTS"
	// NORMALIZED_FTS_PROCESS_CPU_KERNEL.
	MetricThresholdViewMetricName_NORMALIZED_FTS_PROCESS_CPU_KERNEL MetricThresholdViewMetricName = "NORMALIZED_FTS_PROCESS_CPU_KERNEL"
	// NORMALIZED_FTS_PROCESS_CPU_USER.
	MetricThresholdViewMetricName_NORMALIZED_FTS_PROCESS_CPU_USER MetricThresholdViewMetricName = "NORMALIZED_FTS_PROCESS_CPU_USER"
	// NORMALIZED_SYSTEM_CPU_STEAL.
	MetricThresholdViewMetricName_NORMALIZED_SYSTEM_CPU_STEAL MetricThresholdViewMetricName = "NORMALIZED_SYSTEM_CPU_STEAL"
	// NORMALIZED_SYSTEM_CPU_USER.
	MetricThresholdViewMetricName_NORMALIZED_SYSTEM_CPU_USER MetricThresholdViewMetricName = "NORMALIZED_SYSTEM_CPU_USER"
	// OPCOUNTER_CMD.
	MetricThresholdViewMetricName_OPCOUNTER_CMD MetricThresholdViewMetricName = "OPCOUNTER_CMD"
	// OPCOUNTER_DELETE.
	MetricThresholdViewMetricName_OPCOUNTER_DELETE MetricThresholdViewMetricName = "OPCOUNTER_DELETE"
	// OPCOUNTER_GETMORE.
	MetricThresholdViewMetricName_OPCOUNTER_GETMORE MetricThresholdViewMetricName = "OPCOUNTER_GETMORE"
	// OPCOUNTER_INSERT.
	MetricThresholdViewMetricName_OPCOUNTER_INSERT MetricThresholdViewMetricName = "OPCOUNTER_INSERT"
	// OPCOUNTER_QUERY.
	MetricThresholdViewMetricName_OPCOUNTER_QUERY MetricThresholdViewMetricName = "OPCOUNTER_QUERY"
	// OPCOUNTER_REPL_CMD.
	MetricThresholdViewMetricName_OPCOUNTER_REPL_CMD MetricThresholdViewMetricName = "OPCOUNTER_REPL_CMD"
	// OPCOUNTER_REPL_DELETE.
	MetricThresholdViewMetricName_OPCOUNTER_REPL_DELETE MetricThresholdViewMetricName = "OPCOUNTER_REPL_DELETE"
	// OPCOUNTER_REPL_INSERT.
	MetricThresholdViewMetricName_OPCOUNTER_REPL_INSERT MetricThresholdViewMetricName = "OPCOUNTER_REPL_INSERT"
	// OPCOUNTER_REPL_UPDATE.
	MetricThresholdViewMetricName_OPCOUNTER_REPL_UPDATE MetricThresholdViewMetricName = "OPCOUNTER_REPL_UPDATE"
	// OPCOUNTER_UPDATE.
	MetricThresholdViewMetricName_OPCOUNTER_UPDATE MetricThresholdViewMetricName = "OPCOUNTER_UPDATE"
	// OPERATIONS_SCAN_AND_ORDER.
	MetricThresholdViewMetricName_OPERATIONS_SCAN_AND_ORDER MetricThresholdViewMetricName = "OPERATIONS_SCAN_AND_ORDER"
	// OPLOG_MASTER_LAG_TIME_DIFF.
	MetricThresholdViewMetricName_OPLOG_MASTER_LAG_TIME_DIFF MetricThresholdViewMetricName = "OPLOG_MASTER_LAG_TIME_DIFF"
	// OPLOG_MASTER_TIME.
	MetricThresholdViewMetricName_OPLOG_MASTER_TIME MetricThresholdViewMetricName = "OPLOG_MASTER_TIME"
	// OPLOG_MASTER_TIME_ESTIMATED_TTL.
	MetricThresholdViewMetricName_OPLOG_MASTER_TIME_ESTIMATED_TTL MetricThresholdViewMetricName = "OPLOG_MASTER_TIME_ESTIMATED_TTL"
	// OPLOG_RATE_GB_PER_HOUR.
	MetricThresholdViewMetricName_OPLOG_RATE_GB_PER_HOUR MetricThresholdViewMetricName = "OPLOG_RATE_GB_PER_HOUR"
	// OPLOG_SLAVE_LAG_MASTER_TIME.
	MetricThresholdViewMetricName_OPLOG_SLAVE_LAG_MASTER_TIME MetricThresholdViewMetricName = "OPLOG_SLAVE_LAG_MASTER_TIME"
	// QUERY_EXECUTOR_SCANNED.
	MetricThresholdViewMetricName_QUERY_EXECUTOR_SCANNED MetricThresholdViewMetricName = "QUERY_EXECUTOR_SCANNED"
	// QUERY_EXECUTOR_SCANNED_OBJECTS.
	MetricThresholdViewMetricName_QUERY_EXECUTOR_SCANNED_OBJECTS MetricThresholdViewMetricName = "QUERY_EXECUTOR_SCANNED_OBJECTS"
	// QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED.
	MetricThresholdViewMetricName_QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED MetricThresholdViewMetricName = "QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED"
	// QUERY_TARGETING_SCANNED_PER_RETURNED.
	MetricThresholdViewMetricName_QUERY_TARGETING_SCANNED_PER_RETURNED MetricThresholdViewMetricName = "QUERY_TARGETING_SCANNED_PER_RETURNED"
	// RESTARTS_IN_LAST_HOUR.
	MetricThresholdViewMetricName_RESTARTS_IN_LAST_HOUR MetricThresholdViewMetricName = "RESTARTS_IN_LAST_HOUR"
	// SERVERLESS_CONNECTIONS.
	MetricThresholdViewMetricName_SERVERLESS_CONNECTIONS MetricThresholdViewMetricName = "SERVERLESS_CONNECTIONS"
	// SERVERLESS_CONNECTIONS_PERCENT.
	MetricThresholdViewMetricName_SERVERLESS_CONNECTIONS_PERCENT MetricThresholdViewMetricName = "SERVERLESS_CONNECTIONS_PERCENT"
	// SERVERLESS_DATA_SIZE_TOTAL.
	MetricThresholdViewMetricName_SERVERLESS_DATA_SIZE_TOTAL MetricThresholdViewMetricName = "SERVERLESS_DATA_SIZE_TOTAL"
	// SERVERLESS_NETWORK_BYTES_IN.
	MetricThresholdViewMetricName_SERVERLESS_NETWORK_BYTES_IN MetricThresholdViewMetricName = "SERVERLESS_NETWORK_BYTES_IN"
	// SERVERLESS_NETWORK_BYTES_OUT.
	MetricThresholdViewMetricName_SERVERLESS_NETWORK_BYTES_OUT MetricThresholdViewMetricName = "SERVERLESS_NETWORK_BYTES_OUT"
	// SERVERLESS_NETWORK_NUM_REQUESTS.
	MetricThresholdViewMetricName_SERVERLESS_NETWORK_NUM_REQUESTS MetricThresholdViewMetricName = "SERVERLESS_NETWORK_NUM_REQUESTS"
	// SERVERLESS_OPCOUNTER_CMD.
	MetricThresholdViewMetricName_SERVERLESS_OPCOUNTER_CMD MetricThresholdViewMetricName = "SERVERLESS_OPCOUNTER_CMD"
	// SERVERLESS_OPCOUNTER_DELETE.
	MetricThresholdViewMetricName_SERVERLESS_OPCOUNTER_DELETE MetricThresholdViewMetricName = "SERVERLESS_OPCOUNTER_DELETE"
	// SERVERLESS_OPCOUNTER_GETMORE.
	MetricThresholdViewMetricName_SERVERLESS_OPCOUNTER_GETMORE MetricThresholdViewMetricName = "SERVERLESS_OPCOUNTER_GETMORE"
	// SERVERLESS_OPCOUNTER_INSERT.
	MetricThresholdViewMetricName_SERVERLESS_OPCOUNTER_INSERT MetricThresholdViewMetricName = "SERVERLESS_OPCOUNTER_INSERT"
	// SERVERLESS_OPCOUNTER_QUERY.
	MetricThresholdViewMetricName_SERVERLESS_OPCOUNTER_QUERY MetricThresholdViewMetricName = "SERVERLESS_OPCOUNTER_QUERY"
	// SERVERLESS_OPCOUNTER_UPDATE.
	MetricThresholdViewMetricName_SERVERLESS_OPCOUNTER_UPDATE MetricThresholdViewMetricName = "SERVERLESS_OPCOUNTER_UPDATE"
	// SERVERLESS_TOTAL_READ_UNITS.
	MetricThresholdViewMetricName_SERVERLESS_TOTAL_READ_UNITS MetricThresholdViewMetricName = "SERVERLESS_TOTAL_READ_UNITS"
	// SERVERLESS_TOTAL_WRITE_UNITS.
	MetricThresholdViewMetricName_SERVERLESS_TOTAL_WRITE_UNITS MetricThresholdViewMetricName = "SERVERLESS_TOTAL_WRITE_UNITS"
	// SWAP_USAGE_FREE.
	MetricThresholdViewMetricName_SWAP_USAGE_FREE MetricThresholdViewMetricName = "SWAP_USAGE_FREE"
	// SWAP_USAGE_USED.
	MetricThresholdViewMetricName_SWAP_USAGE_USED MetricThresholdViewMetricName = "SWAP_USAGE_USED"
	// SYSTEM_MEMORY_AVAILABLE.
	MetricThresholdViewMetricName_SYSTEM_MEMORY_AVAILABLE MetricThresholdViewMetricName = "SYSTEM_MEMORY_AVAILABLE"
	// SYSTEM_MEMORY_FREE.
	MetricThresholdViewMetricName_SYSTEM_MEMORY_FREE MetricThresholdViewMetricName = "SYSTEM_MEMORY_FREE"
	// SYSTEM_MEMORY_USED.
	MetricThresholdViewMetricName_SYSTEM_MEMORY_USED MetricThresholdViewMetricName = "SYSTEM_MEMORY_USED"
	// SYSTEM_NETWORK_IN.
	MetricThresholdViewMetricName_SYSTEM_NETWORK_IN MetricThresholdViewMetricName = "SYSTEM_NETWORK_IN"
	// SYSTEM_NETWORK_OUT.
	MetricThresholdViewMetricName_SYSTEM_NETWORK_OUT MetricThresholdViewMetricName = "SYSTEM_NETWORK_OUT"
	// TICKETS_AVAILABLE_READS.
	MetricThresholdViewMetricName_TICKETS_AVAILABLE_READS MetricThresholdViewMetricName = "TICKETS_AVAILABLE_READS"
	// TICKETS_AVAILABLE_WRITES.
	MetricThresholdViewMetricName_TICKETS_AVAILABLE_WRITES MetricThresholdViewMetricName = "TICKETS_AVAILABLE_WRITES"
)

