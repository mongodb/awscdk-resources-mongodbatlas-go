package awscdkresourcesmongodbatlas


type ProjectSettings struct {
	// Flag that indicates whether to collect database-specific metrics for the specified project.
	IsCollectDatabaseSpecificsStatisticsEnabled *bool `field:"optional" json:"isCollectDatabaseSpecificsStatisticsEnabled" yaml:"isCollectDatabaseSpecificsStatisticsEnabled"`
	// Flag that indicates whether to enable the Data Explorer for the specified project.
	IsDataExplorerEnabled *bool `field:"optional" json:"isDataExplorerEnabled" yaml:"isDataExplorerEnabled"`
	// Flag that indicates whether to enable extended storage sizes for the specified project.
	IsExtendedStorageSizesEnabled *bool `field:"optional" json:"isExtendedStorageSizesEnabled" yaml:"isExtendedStorageSizesEnabled"`
	// Flag that indicates whether to enable the Performance Advisor and Profiler for the specified project.
	IsPerformanceAdvisorEnabled *bool `field:"optional" json:"isPerformanceAdvisorEnabled" yaml:"isPerformanceAdvisorEnabled"`
	// Flag that indicates whether to enable the Real Time Performance Panel for the specified project.
	IsRealtimePerformancePanelEnabled *bool `field:"optional" json:"isRealtimePerformancePanelEnabled" yaml:"isRealtimePerformancePanelEnabled"`
	// Flag that indicates whether to enable the Schema Advisor for the specified project.
	IsSchemaAdvisorEnabled *bool `field:"optional" json:"isSchemaAdvisorEnabled" yaml:"isSchemaAdvisorEnabled"`
}

