package auth

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func authenticate(token string, scopes []string) (*models.Principal, error) {

	ok, err := authenticated(token)
	if err != nil {
		return nil, errors.New("authentication error")
	}
	if !ok {
		return nil, errors.New("invalid token")
	}
	prin := models.Principal(token)
	return &prin, nil
}

func authenticated(token string) (bool, error) {

	tokenInfoUrl := os.Getenv("API_TOKENINFO_URL")
	req, err := http.NewRequest("GET", tokenInfoUrl, nil)
	if err != nil {
		return false, fmt.Errorf("http request: %v", err)
	}
	bearerToken := "Bearer " + token
	req.Header.Add("Authorization", bearerToken)

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return false, fmt.Errorf("http request: %v", err)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("fail to get response: %v", err)
	}

	if resp.StatusCode != 200 {
		return false, nil
	}
	return true, nil
}
