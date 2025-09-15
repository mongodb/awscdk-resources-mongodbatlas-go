package awscdkresourcesmongodbatlas


// Collection of Uniform Resource Locators that point to the MongoDB database.
type ConnectionStrings struct {
	// Public connection string that you can use to connect to this cluster.
	//
	// This connection string uses the mongodb:// protocol.
	Standard *string `field:"optional" json:"standard" yaml:"standard"`
	// Public connection string that you can use to connect to this flex cluster.
	//
	// This connection string uses the mongodb+srv:// protocol.
	StandardSrv *string `field:"optional" json:"standardSrv" yaml:"standardSrv"`
}

