// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


// Generates one X.509 certificate for the specified MongoDB user. Atlas manages the certificate and MongoDB user that belong to one project. To use this resource, the requesting API Key must have the Project Atlas Admin role. This resource doesn't require the API Key to have an Access List.
//
// To get MongoDB Cloud to generate a managed certificate for a database user, set `"x509Type" : "MANAGED"` on the desired MongoDB Database User.
//
// If you are managing your own Certificate Authority (CA) in Self-Managed X.509 mode, you must generate certificates for database users using your own CA.
type CfnX509AuthenticationDatabaseUserProps struct {
	// CustomerX509 represents Customer-managed X.509 configuration for an Atlas project.
	CustomerX509 *CustomerX509 `field:"optional" json:"customerX509" yaml:"customerX509"`
	// A number of months that the created certificate is valid for before expiry, up to 24 months.default 3.
	MonthsUntilExpiration *float64 `field:"optional" json:"monthsUntilExpiration" yaml:"monthsUntilExpiration"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// The unique identifier for the project .
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Total number of unexpired certificates returned in this response.
	TotalCount *float64 `field:"optional" json:"totalCount" yaml:"totalCount"`
	// Username of the database user to create a certificate for.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

