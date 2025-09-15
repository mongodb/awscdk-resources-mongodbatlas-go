package awscdkresourcesmongodbatlas


// Group of cloud provider settings that configure the provisioned MongoDB flex cluster.
type ProviderSettings struct {
	// Cloud service provider on which MongoDB Cloud provisioned the flex cluster.
	BackingProviderName *string `field:"required" json:"backingProviderName" yaml:"backingProviderName"`
	// Human-readable label that identifies the geographic location of your MongoDB flex cluster.
	//
	// The region you choose can affect network latency for clients accessing your databases.
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
	// Storage capacity available to the flex cluster expressed in gigabytes.
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Human-readable label that identifies the cloud service provider.
	ProviderName *string `field:"optional" json:"providerName" yaml:"providerName"`
}

