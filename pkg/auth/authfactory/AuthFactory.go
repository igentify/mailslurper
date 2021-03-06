package authfactory

import (
	"github.com/mailslurper/mailslurper/pkg/auth/auth"
	"github.com/mailslurper/mailslurper/pkg/auth/authscheme"
	"github.com/mailslurper/mailslurper/pkg/auth/basicauth"
	"github.com/mailslurper/mailslurper/pkg/mailslurper"
)

/*
AuthFactory returns an Authorization Provider based on
the provided configuration
*/
type AuthFactory struct {
	Config *mailslurper.Configuration
}

/*
Get returns an auth provider
*/
func (f *AuthFactory) Get() auth.IAuthProvider {
	switch f.Config.AuthenticationScheme {
	case authscheme.BASIC:
		return &basicauth.BasicAuthProvider{
			CredentialMap:   f.Config.Credentials,
			PasswordService: &basicauth.PasswordService{},
		}

	default:
		return nil
	}
}
