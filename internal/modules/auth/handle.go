package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	oidc "github.com/coreos/go-oidc"
	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/restapi/operations/authentication"
	"github.com/gmarcha/ent-goswagger-app/internal/modules/user"
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
	user   *user.Service
}

func (c *callback) Handle(params authentication.CallbackParams) middleware.Responder {

	r := params.HTTPRequest

	// Read parameters in query.
	if r.URL.Query().Get("state") != c.state {
		return middleware.Error(401, fmt.Sprintln("invalid state"))
	}
	authCode := r.URL.Query().Get("code")

	ctx := context.Background()
	client := &http.Client{}

	// Proceed to OAuth 2.0 handshake (sending authorisation code and receiving token).
	clientCtx := oidc.ClientContext(ctx, client)
	token, err := c.config.Exchange(clientCtx, authCode)
	if err != nil {
		return middleware.Error(500, fmt.Sprintln("failed to fetch token"))
	}

	fmt.Println(token.AccessToken)
	fmt.Println(token.RefreshToken)
	fmt.Println(token.TokenType)
	fmt.Println(token.Expiry)

	userInfoUrl := os.Getenv("API_USERINFO_URL")
	req, err := http.NewRequest("GET", userInfoUrl, nil)
	if err != nil {
		return middleware.Error(500, fmt.Sprintln("failed to create request"))
	}
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)

	resp, err := client.Do(req)
	if err != nil {
		return middleware.Error(500, fmt.Sprintln("failed to fetch user info"))
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return middleware.Error(401, fmt.Sprintln("invalid token"))
	}

	var info struct {
		Login     string `json:"login"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Image     string `json:"image_url"`
		Staff     bool   `json:"staff?"`
		Groups    []struct {
			Tutor string `json:"name"` // field value: TUTEUR
		} `json:"groups"`
	}

	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return middleware.Error(500, fmt.Sprintln("failed to parse user info"))
	}

	userInfo := &ent.User{
		Login:     info.Login,
		FirstName: info.FirstName,
		LastName:  info.LastName,
		ImagePath: info.Image,
	}

	u, err := c.user.ReadMe(ctx, userInfo.Login)
	if err != nil {
		if !ent.IsNotFound(err) {
			return middleware.Error(500, fmt.Sprintln("failed to fetch internal user info"))
		}
		_, err = c.user.CreateUser(ctx, userInfo)
		if err != nil {
			return middleware.Error(500, fmt.Sprintln("failed to create user"))
		}
	} else {
		_, err = u.Update().
			SetFirstName(userInfo.FirstName).
			SetLastName(userInfo.LastName).
			SetImagePath(userInfo.ImagePath).
			Save(ctx)
		if err != nil {
			return middleware.Error(500, fmt.Sprintln("failed to update user"))
		}
	}

	return authentication.NewCallbackOK().WithPayload(token.AccessToken)
}

type tokenInfo struct {
	user *user.Service
}

func (t *tokenInfo) Handle(params authentication.TokenInfoParams) middleware.Responder {

	return authentication.NewTokenInfoOK().WithPayload(&ent.User{})
}

type tokenRefresh struct {
	user *user.Service
}

func (t *tokenRefresh) Handle(params authentication.TokenRefreshParams) middleware.Responder {

	return authentication.NewTokenRefreshOK().WithPayload("newToken")
}
