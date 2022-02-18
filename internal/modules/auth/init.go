package auth

import (
	"os"

	"github.com/gamarcha/ent-goswagger-app/internal/goswagger/restapi/operations"
	"golang.org/x/oauth2"
)

func Init(api *operations.TutorAPI) {

	state := randStringBytes(64)
	clientID := os.Getenv("API_CLIENT_ID")
	clientSecret := os.Getenv("API_CLIENT_SECRET")
	authUrl := os.Getenv("API_AUTH_URL")
	tokenUrl := os.Getenv("API_TOKEN_URL")
	callbackUrl := os.Getenv("API_CALLBACK_URL")
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"public"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  authUrl,
			TokenURL: tokenUrl,
		},
		RedirectURL: callbackUrl,
	}

	api.OauthSecurityAuth = authenticate
	api.AuthenticationLoginHandler = &login{state: state, config: config}
	api.AuthenticationCallbackHandler = &callback{state: state, config: config}
}