package integration

import (
	"bytes"
	"net/http"
)

const api = "http://localhost:5000/v2"

// Get sends an HTTP GET request to a location and returns an HTTP response.
func Get(client *http.Client, url string) (*http.Response, error) {

	resp, err := client.Get(api + url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetAuth sends an HTTP GET request with an authorization header
// containing a bearer token and returns an HTTP response.
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

// PostAuth sends an HTTP POST request with an authorization header
// containing a bearer token and returns an HTTP response.
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

// PatchAuth sends an HTTP PATCH request with an authorization header
// containing a bearer token and returns an HTTP response.
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

// DeleteAuth sends an HTTP DELETE request with an authorization header
// containing a bearer token and returns an HTTP response.
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
