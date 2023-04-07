// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


type ServerlessInstancePrivateEndpoint struct {
	// List that contains the private endpoints through which you connect to MongoDB Cloud when you use **connectionStrings.privateEndpoint[n].srvConnectionString**.
	Endpoints *[]*ServerlessInstancePrivateEndpointEndpoint `field:"optional" json:"endpoints" yaml:"endpoints"`
	// Private endpoint-aware connection string that uses the `mongodb+srv://` protocol to connect to MongoDB Cloud through a private endpoint.
	//
	// The `mongodb+srv` protocol tells the driver to look up the seed list of hosts in the Domain Name System (DNS).
	SrvConnectionString *string `field:"optional" json:"srvConnectionString" yaml:"srvConnectionString"`
	// MongoDB process type to which your application connects.
	Type ServerlessInstancePrivateEndpointType `field:"optional" json:"type" yaml:"type"`
}

