package integration

import (
	"bytes"
	"net/http"
)

const api = "http://localhost:5000/v2"

func Get(client *http.Client, url string) (*http.Response, error) {

	resp, err := client.Get(api + url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetAuth(client *http.Client, url, token string) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodGet, api+url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func PostAuth(client *http.Client, url, token, body string) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodPost, api+url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func PatchAuth(client *http.Client, url, token, body string) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodPatch, api+url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func DeleteAuth(client *http.Client, url, token string) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodDelete, api+url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
