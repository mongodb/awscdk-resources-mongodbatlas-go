package awscdkresourcesmongodbatlas


type SecretDefinition struct {
	// Date and time that the secret was created on.
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Date and time when the secret expires.
	ExpiresAt *string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// Unique identifier of the secret.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Date and time when the secret was last used.
	LastUsedAt *string `field:"optional" json:"lastUsedAt" yaml:"lastUsedAt"`
	// Masked value of the secret.
	MaskedSecretValue *string `field:"optional" json:"maskedSecretValue" yaml:"maskedSecretValue"`
	// The secret value.
	//
	// Only returned on create.
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
}

