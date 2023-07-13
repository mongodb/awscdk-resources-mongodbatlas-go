package awscdkresourcesmongodbatlas


type ApiAtlasNdsUserToDnMappingView struct {
	// Lightweight Directory Access Protocol (LDAP) query template that inserts the LDAP name that the regular expression matches into an LDAP query Uniform Resource Identifier (URI).
	//
	// The formatting for the query must conform to [RFC 4515](https://datatracker.ietf.org/doc/html/rfc4515) and [RFC 4516](https://datatracker.ietf.org/doc/html/rfc4516).
	LdapQuery *string `field:"optional" json:"ldapQuery" yaml:"ldapQuery"`
	// Regular expression that MongoDB Cloud uses to match against the provided Lightweight Directory Access Protocol (LDAP) username.
	//
	// Each parenthesis-enclosed section represents a regular expression capture group that the substitution or `ldapQuery` template uses.
	Match *string `field:"optional" json:"match" yaml:"match"`
	// Lightweight Directory Access Protocol (LDAP) Distinguished Name (DN) template that converts the LDAP username that matches regular expression in the *match* parameter into an LDAP Distinguished Name (DN).
	Substitution *string `field:"optional" json:"substitution" yaml:"substitution"`
}

