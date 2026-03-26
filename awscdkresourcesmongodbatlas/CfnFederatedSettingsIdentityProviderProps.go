package awscdkresourcesmongodbatlas


// Resource for managing MongoDB Atlas Federated Settings Identity Providers (SAML and OIDC) within an Atlas federation.
type CfnFederatedSettingsIdentityProviderProps struct {
	// Unique 24-hexadecimal digit string that identifies your federation.
	FederationSettingsId *string `field:"required" json:"federationSettingsId" yaml:"federationSettingsId"`
	// Issuer URI of the identity provider.
	IssuerUri *string `field:"required" json:"issuerUri" yaml:"issuerUri"`
	// Human-readable name (display name) of the identity provider.
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of associated domains for this identity provider.
	AssociatedDomains *[]*string `field:"optional" json:"associatedDomains" yaml:"associatedDomains"`
	// OIDC audience.
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// OIDC authorization type.
	AuthorizationType *string `field:"optional" json:"authorizationType" yaml:"authorizationType"`
	// OIDC client ID.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Description of the identity provider.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// OIDC groups claim.
	GroupsClaim *string `field:"optional" json:"groupsClaim" yaml:"groupsClaim"`
	// Identity provider type (for OIDC).
	IdpType *string `field:"optional" json:"idpType" yaml:"idpType"`
	// The profile is defined in AWS Secrets Manager.
	//
	// See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Identity provider protocol.
	Protocol CfnFederatedSettingsIdentityProviderPropsProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// SAML request binding.
	RequestBinding *string `field:"optional" json:"requestBinding" yaml:"requestBinding"`
	// OIDC requested scopes.
	RequestedScopes *[]*string `field:"optional" json:"requestedScopes" yaml:"requestedScopes"`
	// SAML response signature algorithm.
	ResponseSignatureAlgorithm *string `field:"optional" json:"responseSignatureAlgorithm" yaml:"responseSignatureAlgorithm"`
	// Flag that indicates whether to enable SSO debug.
	SsoDebugEnabled *bool `field:"optional" json:"ssoDebugEnabled" yaml:"ssoDebugEnabled"`
	// SSO URL.
	SsoUrl *string `field:"optional" json:"ssoUrl" yaml:"ssoUrl"`
	// Identity provider status.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// OIDC user claim.
	UserClaim *string `field:"optional" json:"userClaim" yaml:"userClaim"`
}

