package github

import (
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"

	"github.com/koderover/zadig/pkg/microservice/systemconfig/core/codehost/internal/oauth"
)

type oAuth struct {
	oauth2Config *oauth2.Config
}

func New(callbackURL, clientID, clientSecret, address string) oauth.Oauth {
	endpoint := github.Endpoint
	if address != "" {
		endpoint = oauth2.Endpoint{
			AuthURL:  address + "/login/oauth/authorize",
			TokenURL: address + "/login/oauth/access_token",
		}
	}
	return &oAuth{
		oauth2Config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Endpoint:     endpoint,
			RedirectURL:  callbackURL,
			Scopes:       []string{"repo", "user"},
		},
	}
}

func (o *oAuth) LoginURL(state string) string {
	return o.oauth2Config.AuthCodeURL(state)
}

func (o *oAuth) HandleCallback(r *http.Request) (*oauth2.Token, error) {
	q := r.URL.Query()
	if errType := q.Get("error"); errType != "" {
		return nil, &oauth.OAuth2Error{errType, q.Get("error_description")}
	}
	return o.oauth2Config.Exchange(r.Context(), q.Get("code"))
}
