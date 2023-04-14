// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type Link struct {
	// Uniform Resource Locator (URL) that points another API resource to which this response has some relationship.
	//
	// This URL often begins with `https://mms.mongodb.com`.
	Href *string `field:"optional" json:"href" yaml:"href"`
	// Uniform Resource Locator (URL) that defines the semantic relationship between this resource and another API resource.
	//
	// This URL often begins with `https://mms.mongodb.com`.
	Rel *string `field:"optional" json:"rel" yaml:"rel"`
}

