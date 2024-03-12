package awscdkresourcesmongodbatlas


// Type of ingestion source of this Data Lake Pipeline.
type SourceType string

const (
	// ON_DEMAND_CPS.
	SourceType_ON_DEMAND_CPS SourceType = "ON_DEMAND_CPS"
	// PERIODIC_CPS.
	SourceType_PERIODIC_CPS SourceType = "PERIODIC_CPS"
)

