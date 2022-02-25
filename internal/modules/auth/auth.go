package auth

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

func authenticate(token string, scopes []string) (*models.Principal, error) {

	for _, scope := range scopes {
		fmt.Println("authentication scope:", scope)
	}

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
