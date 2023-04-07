// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type AuthConfigProviders string

const (
	// anon-user.
	AuthConfigProviders_ANON_USER AuthConfigProviders = "ANON_USER"
	// api-key.
	AuthConfigProviders_API_KEY AuthConfigProviders = "API_KEY"
	// custom-token.
	AuthConfigProviders_CUSTOM_TOKEN AuthConfigProviders = "CUSTOM_TOKEN"
	// custom-function.
	AuthConfigProviders_CUSTOM_FUNCTION AuthConfigProviders = "CUSTOM_FUNCTION"
	// local-userpass.
	AuthConfigProviders_LOCAL_USERPASS AuthConfigProviders = "LOCAL_USERPASS"
	// oauth2-apple.
	AuthConfigProviders_OAUTH2_APPLE AuthConfigProviders = "OAUTH2_APPLE"
	// oauth2-facebook.
	AuthConfigProviders_OAUTH2_FACEBOOK AuthConfigProviders = "OAUTH2_FACEBOOK"
	// oauth2-google.
	AuthConfigProviders_OAUTH2_GOOGLE AuthConfigProviders = "OAUTH2_GOOGLE"
)

