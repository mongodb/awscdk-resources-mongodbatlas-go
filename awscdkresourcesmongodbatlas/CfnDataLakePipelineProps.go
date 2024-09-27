package awscdkresourcesmongodbatlas


// Data Lake is deprecated.
//
// As of September 2024, Data Lake is deprecated. If you use Data Lake, you should migrate to alternative solutions before the service is removed. To learn more, see <https://dochub.mongodb.org/core/data-lake-deprecation>. This resource returns, adds, edits, and removes data lake pipelines.
type CfnDataLakePipelineProps struct {
	// Name of this Data Lake Pipeline.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	Sink *Sink `field:"required" json:"sink" yaml:"sink"`
	// Ingestion destination of a Data Lake Pipeline.
	Transformations *[]*Transformations `field:"required" json:"transformations" yaml:"transformations"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	Source *Source `field:"optional" json:"source" yaml:"source"`
	// State of the Data Lake Pipeline.
	State *string `field:"optional" json:"state" yaml:"state"`
}

