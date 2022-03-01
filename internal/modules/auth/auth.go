package auth

import (
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
	"github.com/go-openapi/errors"
	"github.com/go-redis/redis/v8"
)

type authenticatorFunc func(string, []string) (*models.Principal, error)

func (f authenticatorFunc) auth(tokenString string, scopes []string) (*models.Principal, error) {
	return f(tokenString, scopes)
}

type authenticator interface {
	auth(string, []string) (*models.Principal, error)
}

func authenticate(rdb *redis.Client, accessTokenState string) authenticator {

	return authenticatorFunc(func(tokenString string, scopes []string) (*models.Principal, error) {

		token, err := parseToken(rdb, accessTokenState, tokenString)
		if err != nil || !token.Valid {
			return nil, errors.New(401, "Unauthorized")
		}

		claims := token.Claims.(*userClaims)

		permissions := map[string]bool{"public": true}

		for _, role := range claims.Roles {
			switch role {
			case "tutor":
				permissions["event"] = true
				permissions["user"] = true
				permissions["user:subscription"] = true
			case "calendar":
				permissions["event"] = true
				permissions["event:write"] = true
				permissions["user"] = true
			case "admin":
				permissions["event"] = true
				permissions["event:write"] = true
				permissions["user"] = true
				permissions["user:subscription"] = true
				permissions["user:write"] = true
			}
		}

		for _, scope := range scopes {
			if _, ok := permissions[scope]; !ok {
				return nil, errors.New(403, "Forbidden")
			}
		}

		prin := models.Principal(claims.Issuer)
		return &prin, nil
	})
}
