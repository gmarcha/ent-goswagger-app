package integration

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
	"github.com/stretchr/testify/suite"
)

type AuthSuite struct {
	suite.Suite
	client *http.Client
	token  *models.Token
}

func (s *AuthSuite) SetupSuite() {
	s.client = &http.Client{}
	s.token = &models.Token{}
}

func (s *AuthSuite) TestAuth() {
	s.Run("Login", s.Login)
	s.Run("TokenInfo", s.TokenInfo)
	s.Run("TokenInfoInvalidToken", s.TokenInfoInvalidToken)
	s.Run("TokenRefresh", s.TokenRefresh)
	s.Run("TokenRefreshExpiredToken", s.TokenRefreshExpiredToken)
	s.Run("TokenRefreshInvalidToken", s.TokenRefreshInvalidToken)
}

func (s *AuthSuite) Login() {
	resp, err := Get(s.client, "/auth/login")
	s.Require().NoError(err)
	s.Require().Equal(200, resp.StatusCode)
	s.Require().NoError(json.NewDecoder(resp.Body).Decode(s.token))
	defer resp.Body.Close()
	s.Require().NotEmpty(s.token.AccessToken, "missing acces token in response")
	s.Require().NotEmpty(s.token.RefreshToken, "missing refresh token in response")
	time.Sleep(time.Second)
}

func (s *AuthSuite) TokenInfo() {
	resp, err := GetAuth(s.client, "/auth/token/info", s.token.AccessToken)
	s.NoError(err)
	s.Equal(200, resp.StatusCode)
	body := &models.TokenInfo{}
	s.NoError(json.NewDecoder(resp.Body).Decode(body))
	defer resp.Body.Close()
	s.NotEmpty(body.Login, "missing login in response")
	s.NotEmpty(body.ExpiresAt, "missing expiration time in response")
}

func (s *AuthSuite) TokenInfoInvalidToken() {
	resp, err := GetAuth(s.client, "/auth/token/info", "")
	s.NoError(err)
	s.Equal(401, resp.StatusCode)
}

func (s *AuthSuite) TokenRefresh() {
	resp, err := GetAuth(s.client, "/auth/token/refresh", s.token.RefreshToken)
	s.NoError(err)
	s.Equal(200, resp.StatusCode)
	body := &models.Token{}
	s.NoError(json.NewDecoder(resp.Body).Decode(body))
	defer resp.Body.Close()
	s.NotEmpty(body.AccessToken, "missing acces token in response")
	s.NotEmpty(body.RefreshToken, "missing refresh token in response")
	s.NotEqual(body.AccessToken, s.token.AccessToken, "access token is not different")
	s.NotEqual(body.RefreshToken, s.token.RefreshToken, "refresh token is not different")
	time.Sleep(time.Second)
}

func (s *AuthSuite) TokenRefreshExpiredToken() {
	resp, err := GetAuth(s.client, "/auth/token/refresh", s.token.RefreshToken)
	s.NoError(err)
	s.Equal(401, resp.StatusCode)
}

func (s *AuthSuite) TokenRefreshInvalidToken() {
	resp, err := GetAuth(s.client, "/auth/token/refresh", "")
	s.NoError(err)
	s.Equal(401, resp.StatusCode)
}
