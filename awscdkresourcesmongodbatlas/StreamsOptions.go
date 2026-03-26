package awscdkresourcesmongodbatlas


// Optional configuration for the stream processor.
type StreamsOptions struct {
	Dlq *StreamsDlq `field:"required" json:"dlq" yaml:"dlq"`
}

