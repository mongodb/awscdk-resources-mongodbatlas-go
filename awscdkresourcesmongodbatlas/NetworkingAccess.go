package awscdkresourcesmongodbatlas


// Network access configuration.
type NetworkingAccess struct {
	// Unique identifier of the connection.
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// Reserved.
	//
	// Will be used by PRIVATE_LINK connection type.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Reserved.
	//
	// Will be used by TRANSIT_GATEWAY connection type.
	TgwRouteId *string `field:"optional" json:"tgwRouteId" yaml:"tgwRouteId"`
	// Type of network access.
	//
	// Can be PUBLIC, VPC, PRIVATE_LINK, or TRANSIT_GATEWAY.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

