package auth

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gmarcha/ent-goswagger-app/internal/modules/user"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
)

func createToken(ctx context.Context, user *user.Service, accessTokenState string, login string, duration time.Duration) (string, error) {

	u, err := user.ReadUserByLogin(ctx, login)
	if err != nil {
		return "", fmt.Errorf("failed to fetch user")
	}

	claims := &userClaims{
		make([]string, 0, 8),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			Issuer:    u.Login,
		},
	}

	for _, role := range u.Edges.Roles {
		claims.Roles = append(claims.Roles, role.Name)
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString([]byte(accessTokenState))
	if err != nil {
		return "", fmt.Errorf("failed to create token")
	}

	return tokenString, nil
}

func parseToken(rdb *redis.Client, accessTokenState string, tokenString string) (*jwt.Token, error) {

	token, err := jwt.ParseWithClaims(tokenString, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(accessTokenState), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func readAndParseToken(rdb *redis.Client, accessTokenState string, authorizationHeader string) (*jwt.Token, error) {

	bearerToken := strings.Split(authorizationHeader, " ")

	if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		return nil, fmt.Errorf("invalid header")
	}

	token, err := parseToken(rdb, accessTokenState, bearerToken[1])
	if err != nil {
		return nil, err
	}
	return token, nil
}

func saveOauthTokenInStore(rdb *redis.Client, login string, token *oauth2.Token) error {

	ctx := context.Background()
	err := rdb.HSet(ctx, login, map[string]interface{}{
		"AccessToken":  token.AccessToken,
		"TokenType":    token.TokenType,
		"RefreshToken": token.RefreshToken,
		"Expiry":       token.Expiry,
	}).Err()
	if err != nil {
		return err
	}
	return nil
}

func retrieveOauthTokenFromStore(rdb *redis.Client, login string) (*oauth2.Token, error) {

	ctx := context.Background()
	val, err := rdb.HGetAll(ctx, login).Result()
	if err != nil {
		return nil, err
	}

	expiryTime, err := time.Parse(time.RFC3339, val["Expiry"])
	if err != nil {
		return nil, err
	}
	oauthToken := &oauth2.Token{
		AccessToken:  val["AccessToken"],
		TokenType:    val["TokenType"],
		RefreshToken: val["RefreshToken"],
		Expiry:       expiryTime,
	}
	return oauthToken, nil
}

func validateRefreshToken(rdb *redis.Client, token *jwt.Token) error {

	ctx := context.Background()
	return rdb.Get(ctx, token.Raw).Err()
}

func blacklistRefreshToken(rdb *redis.Client, token *jwt.Token) error {

	ctx := context.Background()
	return rdb.Set(ctx, token.Raw, "true", time.Hour*96).Err()
}
