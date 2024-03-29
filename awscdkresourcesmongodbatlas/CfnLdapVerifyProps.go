package awscdkresourcesmongodbatlas


// Requests a verification of an LDAP configuration over TLS for an Atlas project.
//
// Pass the requestId in the response object to the Verify |ldap| Configuration endpoint to get the status of a verification request. Atlas retains only the most recent request for each project.
type CfnLdapVerifyProps struct {
	// Password that MongoDB Cloud uses to authenticate the **bindUsername**.
	BindPassword *string `field:"required" json:"bindPassword" yaml:"bindPassword"`
	// Full Distinguished Name (DN) of the Lightweight Directory Access Protocol (LDAP) user that MongoDB Cloud uses to connect to the LDAP host.
	//
	// LDAP distinguished names must be formatted according to RFC 2253.
	BindUsername *string `field:"required" json:"bindUsername" yaml:"bindUsername"`
	// Human-readable label that identifies the hostname or Internet Protocol (IP) address of the Lightweight Directory Access Protocol (LDAP) host.
	//
	// This host must have access to the internet or have a Virtual Private Cloud (VPC) peering connection to your cluster.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Port to which the Lightweight Directory Access Protocol (LDAP) host listens for client connections.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Lightweight Directory Access Protocol (LDAP) query template that MongoDB Cloud runs to obtain the LDAP groups associated with the authenticated user.
	//
	// MongoDB Cloud uses this parameter only for user authorization. Use the `{USER}` placeholder in the Uniform Resource Locator (URL) to substitute the authenticated username. The query relates to the host specified with the hostname. Format this query according to [RFC 4515](https://tools.ietf.org/search/rfc4515) and [RFC 4516](https://datatracker.ietf.org/doc/html/rfc4516).
	AuthzQueryTemplate *string `field:"optional" json:"authzQueryTemplate" yaml:"authzQueryTemplate"`
	// Certificate Authority (CA) certificate that MongoDB Cloud uses to verify the identity of the Lightweight Directory Access Protocol (LDAP) host.
	//
	// MongoDB Cloud allows self-signed certificates. To delete an assigned value, pass an empty string: `"caCertificate": ""`
	CaCertificate *string `field:"optional" json:"caCertificate" yaml:"caCertificate"`
	// Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// List of validation messages related to the verification of the provided LDAP over TLS configuration details.
	//
	// The array contains a document for each test that Atlas runs. Atlas stops running tests after the first failure.
	Validations *[]*Validation `field:"optional" json:"validations" yaml:"validations"`
}

