package credselect

import "github.com/1Password/shell-plugins/sdk"

const (
	Any                             = sdk.CredentialSelector("any")
	SAMLIdentityProvider            = sdk.CredentialSelector("saml")
	CanAuthenticateHTTPRequests     = sdk.CredentialSelector("http")
	CanAuthenticateToDockerRegistry = sdk.CredentialSelector("docker-registry")
)

var details = map[sdk.CredentialSelector]sdk.CredentialSelectorDetails{
	Any: {
		Description: "Any credential from any plugin",
	},
	SAMLIdentityProvider: {
		Description: "Credentials that can be used for SAML authentication",
	},
	CanAuthenticateHTTPRequests: {
		Description: "Credentials that can authenticate HTTP requests",
	},
	CanAuthenticateToDockerRegistry: {
		Description: "Credentials that can be used to authenticate to a Docker Registry",
	},
}
