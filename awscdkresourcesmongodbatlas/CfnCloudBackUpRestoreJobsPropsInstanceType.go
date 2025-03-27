package awscdkresourcesmongodbatlas


// Type of instance specified on the Instance Name serverless or cluster.
//
// **WARNING:** `serverless` instance type is deprecated and will be removed in January 2026. For more details, see [Migrate your programmatic tools from M2, M5, or Serverless Instances to Flex Clusters](https://www.mongodb.com/docs/atlas/flex-migration/).
type CfnCloudBackUpRestoreJobsPropsInstanceType string

const (
	// serverless.
	CfnCloudBackUpRestoreJobsPropsInstanceType_SERVERLESS CfnCloudBackUpRestoreJobsPropsInstanceType = "SERVERLESS"
	// cluster.
	CfnCloudBackUpRestoreJobsPropsInstanceType_CLUSTER CfnCloudBackUpRestoreJobsPropsInstanceType = "CLUSTER"
)

