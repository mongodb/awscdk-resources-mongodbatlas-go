package awscdkresourcesmongodbatlas


type ApiSearchDeploymentSpec struct {
	// Hardware specification for the search node instance sizes.
	//
	// The [MongoDB Atlas API](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Atlas-Search/operation/createAtlasSearchDeployment) describes the valid values. More details can also be found in the [Search Node Documentation](https://www.mongodb.com/docs/atlas/cluster-config/multi-cloud-distribution/#search-tier).
	InstanceSize *string `field:"required" json:"instanceSize" yaml:"instanceSize"`
	// Number of search nodes in the cluster.
	NodeCount *float64 `field:"required" json:"nodeCount" yaml:"nodeCount"`
}

