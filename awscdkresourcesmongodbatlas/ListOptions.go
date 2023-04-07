// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type ListOptions struct {
	// Flag that indicates whether the response returns the total number of items (totalCount) in the response.
	IncludeCount *bool `field:"optional" json:"includeCount" yaml:"includeCount"`
	// Number of items that the response returns per page.
	ItemsPerPage *float64 `field:"optional" json:"itemsPerPage" yaml:"itemsPerPage"`
	// Number of the page that displays the current set of the total objects that the response returns.
	PageNum *float64 `field:"optional" json:"pageNum" yaml:"pageNum"`
}

