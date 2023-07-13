package awscdkresourcesmongodbatlas


// Type of restore job to create.The value can be any one of download,automated or point_in_time.
type CfnCloudBackUpRestoreJobsPropsDeliveryType string

const (
	// download.
	CfnCloudBackUpRestoreJobsPropsDeliveryType_DOWNLOAD CfnCloudBackUpRestoreJobsPropsDeliveryType = "DOWNLOAD"
	// automated.
	CfnCloudBackUpRestoreJobsPropsDeliveryType_AUTOMATED CfnCloudBackUpRestoreJobsPropsDeliveryType = "AUTOMATED"
	// pointInTime.
	CfnCloudBackUpRestoreJobsPropsDeliveryType_POINT_IN_TIME CfnCloudBackUpRestoreJobsPropsDeliveryType = "POINT_IN_TIME"
)

