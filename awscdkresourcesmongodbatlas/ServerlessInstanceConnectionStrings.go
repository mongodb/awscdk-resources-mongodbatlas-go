package awscdkresourcesmongodbatlas


type ServerlessInstanceConnectionStrings struct {
	// List of private endpoint connection strings that you can use to connect to this serverless instance through a private endpoint.
	//
	// This parameter returns only if you created a private endpoint for this serverless instance and it is AVAILABLE.
	PrivateEndpoint *[]*ServerlessInstancePrivateEndpoint `field:"optional" json:"privateEndpoint" yaml:"privateEndpoint"`
	// Public connection string that you can use to connect to this serverless instance.
	//
	// This connection string uses the `mongodb+srv://` protocol.
	StandardSrv *string `field:"optional" json:"standardSrv" yaml:"standardSrv"`
}

