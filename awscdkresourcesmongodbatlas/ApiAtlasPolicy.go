package awscdkresourcesmongodbatlas


type ApiAtlasPolicy struct {
	// A string that defines the permissions for the policy.
	//
	// The syntax used is the Cedar Policy language.
	Body *string `field:"required" json:"body" yaml:"body"`
	// Unique 24-hexadecimal character string that identifies the policy.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

