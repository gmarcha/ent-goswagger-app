package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	oidc "github.com/coreos/go-oidc"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/restapi/operations/authentication"
	"github.com/gmarcha/ent-goswagger-app/internal/modules/user"
	e "github.com/gmarcha/ent-goswagger-app/internal/utils"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/oauth2"
)

type login struct {
	oauthState  string
	oauthConfig *oauth2.Config
}

func (l *login) Handle(params authentication.LoginParams) middleware.Responder {
	return middleware.ResponderFunc(
		func(w http.ResponseWriter, pr runtime.Producer) {
			http.Redirect(w, params.HTTPRequest, l.oauthConfig.AuthCodeURL(l.oauthState), http.StatusFound)
		})
}

type callback struct {
	userInfoUrl          string
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
	accessTokenState     string
	refreshTokenState    string
	oauthState           string
	oauthConfig          *oauth2.Config
	user                 *user.Service
	rdb                  *redis.Client
}

func (c *callback) Handle(params authentication.CallbackParams) middleware.Responder {

	// Retrieve state and authorization code from query parameters;
	// if state differs, return an error.
	if params.HTTPRequest.URL.Query().Get("state") != c.oauthState {
		return authentication.NewCallbackUnprocessableEntity().WithPayload(e.Err(422, fmt.Errorf("invalid state")))
	}
	authCode := params.HTTPRequest.URL.Query().Get("code")

	ctx := context.Background()
	client := &http.Client{}

	// Retrieve an OAuth token from 42 API.
	clientCtx := oidc.ClientContext(ctx, client)
	token, err := c.oauthConfig.Exchange(clientCtx, authCode)
	if err != nil {
		return authentication.NewCallbackInternalServerError().WithPayload(e.Err(500, err))
	}

	// Fetch user informations in 42 API with access token.
	userInfo, err := fetchUserInfo(client, c.userInfoUrl, token.AccessToken)
	if err != nil {
		return authentication.NewCallbackInternalServerError().WithPayload(e.Err(500, err))
	}

	// Create or update user informations in database.
	user, err := c.user.SetUserOnLogin(ctx, userInfo)
	if err != nil {
		return authentication.NewCallbackInternalServerError().WithPayload(e.Err(500, err))
	}

	err = saveOauthTokenInStore(c.rdb, user.Login, token)
	if err != nil {
		return authentication.NewCallbackInternalServerError().WithPayload(e.Err(500, err))
	}

	// Create and return a JSON web token to authenticate API consumers.
	accessToken, err := createToken(ctx, c.user, c.accessTokenState, user.Login, c.accessTokenDuration)
	if err != nil {
		return authentication.NewCallbackInternalServerError().WithPayload(e.Err(500, err))
	}

	refreshToken, err := createToken(ctx, c.user, c.refreshTokenState, user.Login, c.refreshTokenDuration)
	if err != nil {
		return authentication.NewCallbackInternalServerError().WithPayload(e.Err(500, err))
	}

	return authentication.NewCallbackOK().WithPayload(&models.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

type tokenInfo struct {
	accessTokenState string
	rdb              *redis.Client
}

func (t *tokenInfo) Handle(params authentication.TokenInfoParams) middleware.Responder {

	token, err := readAndParseToken(t.rdb, t.accessTokenState, params.Authorization)
	if err != nil || !token.Valid {
		return authentication.NewTokenInfoUnauthorized().WithPayload(e.Err(401, err))
	}

	claims := token.Claims.(*userClaims)

	return authentication.NewTokenInfoOK().WithPayload(&models.TokenInfo{
		ExpiresAt: strfmt.DateTime(claims.ExpiresAt.Time),
		Login:     claims.Issuer,
	})
}

type tokenRefresh struct {
	userInfoUrl          string
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
	accessTokenState     string
	refreshTokenState    string
	oauthConfig          *oauth2.Config
	user                 *user.Service
	rdb                  *redis.Client
}

func (t *tokenRefresh) Handle(params authentication.TokenRefreshParams) middleware.Responder {

	token, err := readAndParseToken(t.rdb, t.refreshTokenState, params.Authorization)
	if err != nil || !token.Valid {
		return authentication.NewTokenRefreshUnauthorized().WithPayload(e.Err(401, err))
	}
	err = validateRefreshToken(t.rdb, token)
	if err == nil {
		return authentication.NewTokenRefreshUnauthorized().WithPayload(e.Err(401, fmt.Errorf("invalid token")))
	} else if err != redis.Nil {
		return authentication.NewTokenRefreshInternalServerError().WithPayload(e.Err(500, err))
	}
	if blacklistRefreshToken(t.rdb, token) != nil {
		return authentication.NewTokenRefreshInternalServerError().WithPayload(e.Err(500, err))
	}

	claims := token.Claims.(*userClaims)

	oauthToken, err := retrieveOauthTokenFromStore(t.rdb, claims.Issuer)
	if err != nil {
		return authentication.NewTokenRefreshInternalServerError().WithPayload(e.Err(500, err))
	}

	ctx := context.Background()
	client := t.oauthConfig.Client(ctx, oauthToken)
	res, err := client.Get(t.userInfoUrl)
	if err != nil {
		return authentication.NewTokenRefreshInternalServerError().WithPayload(e.Err(500, err))
	}

	if res.StatusCode != 200 {
		return authentication.NewTokenRefreshUnauthorized().WithPayload(e.Err(401, fmt.Errorf("invalid token")))
	}

	err = saveOauthTokenInStore(t.rdb, claims.Issuer, oauthToken)
	if err != nil {
		return authentication.NewTokenRefreshInternalServerError().WithPayload(e.Err(500, err))
	}

	// Create and return a JSON web token to authenticate API consumers.
	accessToken, err := createToken(ctx, t.user, t.accessTokenState, claims.Issuer, t.accessTokenDuration)
	if err != nil {
		return authentication.NewTokenRefreshInternalServerError().WithPayload(e.Err(500, err))
	}

	refreshToken, err := createToken(ctx, t.user, t.refreshTokenState, claims.Issuer, t.refreshTokenDuration)
	if err != nil {
		return authentication.NewTokenRefreshInternalServerError().WithPayload(e.Err(500, err))
	}

	return authentication.NewTokenRefreshOK().WithPayload(&models.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
