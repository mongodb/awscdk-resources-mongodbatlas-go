package awscdkresourcesmongodbatlas


type ServerlessInstanceProviderSettings struct {
	// Human-readable label that identifies the cloud service provider.
	ProviderName ServerlessInstanceProviderSettingsProviderName `field:"optional" json:"providerName" yaml:"providerName"`
	// Human-readable label that identifies the geographic location of your MongoDB serverless instance.
	//
	// The region you choose can affect network latency for clients accessing your databases. For a complete list of region names, see [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/#std-label-amazon-aws), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), and [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

