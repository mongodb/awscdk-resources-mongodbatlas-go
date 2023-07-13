package awscdkresourcesmongodbatlas


type ZoneMapping struct {
	// Code that represents a location that maps to a zone in your global cluster.
	//
	// MongoDB Cloud represents this location with a ISO 3166-2 location and subdivision codes when possible.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Human-readable label that identifies the zone in your global cluster.
	//
	// This zone maps to a location code.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

