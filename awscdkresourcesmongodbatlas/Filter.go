package awscdkresourcesmongodbatlas


type Filter struct {
	CloudProvider FilterCloudProvider `field:"optional" json:"cloudProvider" yaml:"cloudProvider"`
	Region *string `field:"optional" json:"region" yaml:"region"`
	Type *string `field:"optional" json:"type" yaml:"type"`
}

