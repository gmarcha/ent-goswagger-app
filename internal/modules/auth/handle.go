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

func (this *login) Handle(params authentication.LoginParams) middleware.Responder {
	return middleware.ResponderFunc(
		func(w http.ResponseWriter, pr runtime.Producer) {
			http.Redirect(w, params.HTTPRequest, this.config.AuthCodeURL(this.state), http.StatusFound)
		})
}

type callback struct {
	state  string
	config *oauth2.Config
	// user   *user.Service
}

func (this *callback) Handle(params authentication.CallbackParams) middleware.Responder {

	r := params.HTTPRequest

	if r.URL.Query().Get("state") != this.state {
		return middleware.Error(401, fmt.Sprintln("invalid state"))
	}
	code := r.URL.Query().Get("code")

	client := &http.Client{}
	ctx := oidc.ClientContext(context.Background(), client)
	token, err := this.config.Exchange(ctx, code)
	if err != nil {
		return middleware.Error(500, fmt.Sprintln("failed to fetch token"))
	}

	// With that token, make a /me on 42 API.
	// --> handle back id and login;
	// --> it exists in our user, cool, create a JWT;
	// --> it doesn't exist throw CallBackNotOK.

	return authentication.NewCallbackOK().WithPayload(token.AccessToken)
}
