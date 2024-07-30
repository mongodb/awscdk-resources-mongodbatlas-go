package awscdkresourcesmongodbatlas


type ApiPolicyView struct {
	// Unique 24-hexadecimal digit string that identifies this backup policy.
	//
	// The policy id can be retrieved by running: atlas backups schedule describe "${clusterName}" --projectId "${projectId}" | jq -r '.policies[0].id'
	Id *string `field:"optional" json:"id" yaml:"id"`
	PolicyItems *[]*ApiPolicyItemView `field:"optional" json:"policyItems" yaml:"policyItems"`
}

