package auth

import (
	"context"
	"fmt"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/restapi/operations/authentication"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/oauth2"
)

type login struct {
	state  string
	config *oauth2.Config
}

func (l *login) Handle(params authentication.LoginParams) middleware.Responder {
	return middleware.ResponderFunc(
		func(w http.ResponseWriter, pr runtime.Producer) {
			http.Redirect(w, params.HTTPRequest, l.config.AuthCodeURL(l.state), http.StatusFound)
		})
}

type callback struct {
	state  string
	config *oauth2.Config
	// user   *user.Service
}

func (c *callback) Handle(params authentication.CallbackParams) middleware.Responder {

	r := params.HTTPRequest

	// Read parameters in query.
	if r.URL.Query().Get("state") != c.state {
		return middleware.Error(401, fmt.Sprintln("invalid state"))
	}
	authCode := r.URL.Query().Get("code")

	// Proceed to OAuth 2.0 handshake (sending authorisation code and receiving token).
	client := &http.Client{}
	ctx := oidc.ClientContext(context.Background(), client)
	token, err := c.config.Exchange(ctx, authCode)
	if err != nil {
		return middleware.Error(500, fmt.Sprintln("failed to fetch token"))
	}

	// With that token, make a /me on 42 API.
	// --> handle back id and login;
	// --> it exists in our user, cool, create a JWT;
	// --> it doesn't exist throw CallBackNotOK.

	return authentication.NewCallbackOK().WithPayload(token.AccessToken)
}
