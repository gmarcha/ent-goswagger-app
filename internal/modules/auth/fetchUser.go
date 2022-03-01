package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
)

func fetchUserInfo(client *http.Client, userInfoUrl string, token string) (*ent.User, error) {

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

	req, err := http.NewRequest("GET", userInfoUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request")
	}

	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user info")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("invalid token")
	}

	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user info")
	}

	user := &ent.User{
		Login:     info.Login,
		FirstName: info.FirstName,
		LastName:  info.LastName,
		ImagePath: info.Image,
	}
	return user, nil
}
