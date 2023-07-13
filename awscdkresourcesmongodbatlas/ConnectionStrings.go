package awscdkresourcesmongodbatlas


// Collection of Uniform Resource Locators that point to the MongoDB database.
type ConnectionStrings struct {
	// Private endpoint-aware connection strings that use AWS-hosted clusters with Amazon Web Services (AWS) PrivateLink.
	//
	// Each key identifies an Amazon Web Services (AWS) interface endpoint. Each value identifies the related mongodb:// connection string that you use to connect to MongoDB Cloud through the interface endpoint that the key names.
	AwsPrivateLink *string `field:"optional" json:"awsPrivateLink" yaml:"awsPrivateLink"`
	// Private endpoint-aware connection strings that use AWS-hosted clusters with Amazon Web Services (AWS) PrivateLink.
	//
	// Each key identifies an Amazon Web Services (AWS) interface endpoint. Each value identifies the related mongodb:// connection string that you use to connect to Atlas through the interface endpoint that the key names.
	AwsPrivateLinkSrv *string `field:"optional" json:"awsPrivateLinkSrv" yaml:"awsPrivateLinkSrv"`
	// Network peering connection strings for each interface Virtual Private Cloud (VPC) endpoint that you configured to connect to this cluster.
	//
	// This connection string uses the mongodb+srv:// protocol. The resource returns this parameter once someone creates a network peering connection to this cluster. This protocol tells the application to look up the host seed list in the Domain Name System (DNS). This list synchronizes with the nodes in a cluster. If the connection string uses this Uniform Resource Identifier (URI) format, you don't need to append the seed list or change the URI if the nodes change. Use this URI format if your driver supports it. If it doesn't, use connectionStrings.private. For Amazon Web Services (AWS) clusters, this resource returns this parameter only if you enable custom DNS.
	Private *string `field:"optional" json:"private" yaml:"private"`
	// List of private endpoint connection strings that you can use to connect to this cluster through a private endpoint.
	//
	// This parameter returns only if you deployed a private endpoint to all regions to which you deployed this clusters' nodes.
	PrivateEndpoint *[]*PrivateEndpoint `field:"optional" json:"privateEndpoint" yaml:"privateEndpoint"`
	// Network peering connection strings for each interface Virtual Private Cloud (VPC) endpoint that you configured to connect to this cluster.
	//
	// This connection string uses the mongodb+srv:// protocol. The resource returns this parameter when someone creates a network peering connection to this cluster. This protocol tells the application to look up the host seed list in the Domain Name System (DNS). This list synchronizes with the nodes in a cluster. If the connection string uses this Uniform Resource Identifier (URI) format, you don't need to append the seed list or change the Uniform Resource Identifier (URI) if the nodes change. Use this Uniform Resource Identifier (URI) format if your driver supports it. If it doesn't, use connectionStrings.private. For Amazon Web Services (AWS) clusters, this parameter returns only if you enable custom DNS.
	PrivateSrv *string `field:"optional" json:"privateSrv" yaml:"privateSrv"`
	// Public connection string that you can use to connect to this cluster.
	//
	// This connection string uses the mongodb:// protocol.
	Standard *string `field:"optional" json:"standard" yaml:"standard"`
	// Public connection string that you can use to connect to this cluster.
	//
	// This connection string uses the mongodb+srv:// protocol.
	StandardSrv *string `field:"optional" json:"standardSrv" yaml:"standardSrv"`
}

