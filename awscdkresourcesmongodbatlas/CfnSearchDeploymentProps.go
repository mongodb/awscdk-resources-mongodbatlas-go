package awscdkresourcesmongodbatlas


// The resource lets you create, edit and delete dedicated search nodes in a cluster.
//
// For details on supported cloud providers and existing limitations you can visit the Search Node Documentation: https://www.mongodb.com/docs/atlas/cluster-config/multi-cloud-distribution/#search-nodes-for-workload-isolation. Only a single search deployment resource can be defined for each cluster.
type CfnSearchDeploymentProps struct {
	// Label that identifies the cluster to return the search nodes for.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Unique 24-hexadecimal character string that identifies the project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// List of settings that configure the search nodes for your cluster.
	//
	// This list is currently limited to defining a single element.
	Specs *[]*ApiSearchDeploymentSpec `field:"required" json:"specs" yaml:"specs"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

