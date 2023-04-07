// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Returns, adds, edits, and removes serverless instances.
type CfnServerlessInstanceProps struct {
	// Collection of Uniform Resource Locators that point to the MongoDB database.
	ConnectionStrings *ServerlessInstanceConnectionStrings `field:"optional" json:"connectionStrings" yaml:"connectionStrings"`
	// Flag that indicates whether the serverless instances uses Serverless Continuous Backup.
	//
	// If this parameter is false, the serverless instance uses Basic Backup. | Option | Description | |---|---| | Serverless Continuous Backup | Atlas takes incremental snapshots of the data in your serverless instance every six hours and lets you restore the data from a selected point in time within the last 72 hours. Atlas also takes daily snapshots and retains these daily snapshots for 35 days. To learn more, see Serverless Instance Costs. | | Basic Backup | Atlas takes incremental snapshots of the data in your serverless instance every six hours and retains only the two most recent snapshots. You can use this option for free.
	ContinuousBackupEnabled *bool `field:"optional" json:"continuousBackupEnabled" yaml:"continuousBackupEnabled"`
	// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
	IncludeCount *bool `field:"optional" json:"includeCount" yaml:"includeCount"`
	// Number of items that the response returns per page.
	ItemsPerPage *float64 `field:"optional" json:"itemsPerPage" yaml:"itemsPerPage"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both.
	//
	// RFC 5988 outlines these relationships.
	Links *[]*Link `field:"optional" json:"links" yaml:"links"`
	// Human-readable label that identifies the serverless instance.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Number of the page that displays the current set of the total objects that the response returns.
	PageNum *float64 `field:"optional" json:"pageNum" yaml:"pageNum"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Group of settings that configure the provisioned MongoDB serverless instance.
	//
	// The options available relate to the cloud service provider.
	ProviderSettings *ServerlessInstanceProviderSettings `field:"optional" json:"providerSettings" yaml:"providerSettings"`
	// Flag that indicates whether termination protection is enabled on the serverless instance.
	//
	// If set to true, MongoDB Cloud won't delete the serverless instance. If set to false, MongoDB cloud will delete the serverless instance."
	TerminationProtectionEnabled *bool `field:"optional" json:"terminationProtectionEnabled" yaml:"terminationProtectionEnabled"`
}

