package auth

import (
	"os"
	"time"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/restapi/operations"
	"github.com/gmarcha/ent-goswagger-app/internal/modules/user"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
)

type userClaims struct {
	Roles []string
	jwt.RegisteredClaims
}

// Init sets authentication route handlers.
func Init(api *operations.TutorAPI, db *ent.Client, rdb *redis.Client) {

	accessTokenDuration := time.Minute * 30
	refreshTokenDuration := time.Hour * 72
	accessTokenState := os.Getenv("ACCESS_TOKEN_STATE")
	refreshTokenState := os.Getenv("REFRESH_TOKEN_STATE")
	oauthState := os.Getenv("AUTH_STATE")
	oauthConfig := createConfig()
	userInfoUrl := os.Getenv("API_USERINFO_URL")

	userService := &user.Service{User: db.User}

	api.OAuth2Auth = authenticate(rdb, accessTokenState).auth
	api.AuthenticationLoginHandler = &login{
		oauthState:  oauthState,
		oauthConfig: oauthConfig,
	}
	api.AuthenticationCallbackHandler = &callback{
		userInfoUrl:          userInfoUrl,
		accessTokenDuration:  accessTokenDuration,
		refreshTokenDuration: refreshTokenDuration,
		accessTokenState:     accessTokenState,
		refreshTokenState:    refreshTokenState,
		oauthState:           oauthState,
		oauthConfig:          oauthConfig,
		user:                 userService,
		rdb:                  rdb,
	}
	api.AuthenticationAuthorizeHandler = &authorize{
		userInfoUrl:          userInfoUrl,
		accessTokenDuration:  accessTokenDuration,
		refreshTokenDuration: refreshTokenDuration,
		accessTokenState:     accessTokenState,
		refreshTokenState:    refreshTokenState,
		oauthState:           oauthState,
		oauthConfig:          oauthConfig,
		user:                 userService,
		rdb:                  rdb,
	}
	api.AuthenticationTokenInfoHandler = &tokenInfo{
		accessTokenState: accessTokenState,
		rdb:              rdb,
	}
	api.AuthenticationTokenRefreshHandler = &tokenRefresh{
		userInfoUrl:          userInfoUrl,
		accessTokenDuration:  accessTokenDuration,
		refreshTokenDuration: refreshTokenDuration,
		accessTokenState:     accessTokenState,
		refreshTokenState:    refreshTokenState,
		oauthConfig:          oauthConfig,
		user:                 userService,
		rdb:                  rdb,
	}
}

func createConfig() *oauth2.Config {

	clientID := os.Getenv("API_CLIENT_ID")
	clientSecret := os.Getenv("API_CLIENT_SECRET")
	authUrl := os.Getenv("API_AUTH_URL")
	tokenUrl := os.Getenv("API_TOKEN_URL")
	callbackUrl := os.Getenv("API_CALLBACK_URL")

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"public"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  authUrl,
			TokenURL: tokenUrl,
		},
		RedirectURL: callbackUrl,
	}
}
