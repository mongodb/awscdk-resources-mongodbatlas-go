package awscdkresourcesmongodbatlas


// Returns, adds, and removes Atlas Stream Processing Private Link Endpoints.
//
// This resource supports AWS only.
type CfnStreamPrivatelinkEndpointProps struct {
	// Unique 24-hexadecimal digit string that identifies your project.
	//
	// Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access. **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group or project id remains the same. The resource and corresponding endpoints use the term groups.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Provider where the endpoint is deployed.
	//
	// For CloudFormation, this is always AWS.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// Vendor that manages the endpoint.
	//
	// For AWS, valid values are: MSK, CONFLUENT, and S3.
	Vendor *string `field:"required" json:"vendor" yaml:"vendor"`
	// Amazon Resource Name (ARN).
	//
	// Required for AWS Provider and MSK vendor.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The domain hostname.
	//
	// Required for AWS provider with CONFLUENT vendor.
	DnsDomain *string `field:"optional" json:"dnsDomain" yaml:"dnsDomain"`
	// Sub-Domain name of Confluent cluster.
	//
	// These are typically your availability zones. Required for AWS Provider and CONFLUENT vendor. If your AWS CONFLUENT cluster doesn't use subdomains, you must set this to the empty array [].
	DnsSubDomain *[]*string `field:"optional" json:"dnsSubDomain" yaml:"dnsSubDomain"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// The region of the Provider's cluster.
	//
	// See [AWS](https://www.mongodb.com/docs/atlas/reference/amazon-aws/#stream-processing-instances) supported regions. When the vendor is CONFLUENT, this is the domain name of Confluent cluster. When the vendor is MSK, this is computed by the API from the provided ARN.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// For AWS CONFLUENT cluster, this is the [VPC Endpoint service name](https://docs.confluent.io/cloud/current/networking/private-links/aws-privatelink.html). For AWS S3 vendor, this should follow the format 'com.amazonaws.<region>.s3', for example 'com.amazonaws.us-east-1.s3'.
	ServiceEndpointId *string `field:"optional" json:"serviceEndpointId" yaml:"serviceEndpointId"`
}

