package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"personal/TODO-list/pkg/setting"
	"time"

	"golang.org/x/oauth2"
	githubOAuth "golang.org/x/oauth2/github"
)

// GetGithubOAuthConfig will return the config to call github Login
func GetGithubOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     setting.GithubSetting.ClientId,
		ClientSecret: setting.GithubSetting.ClientSecret,
		RedirectURL:  setting.GithubSetting.RedirectUrl,
		Endpoint:     githubOAuth.Endpoint,
		Scopes:       []string{"email"},
	}
}

// Represents the response received from Github
type githubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func GetGithubAccessToken(code string) (string, error) {
	var OAuth2Config = GetGithubOAuthConfig()

	// Set us the request body as JSON
	requestBodyMap := map[string]string{
		"client_id":     OAuth2Config.ClientID,
		"client_secret": OAuth2Config.ClientSecret,
		"code":          code,
	}
	requestJSON, _ := json.Marshal(requestBodyMap)

	// POST request to set URL
	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBuffer(requestJSON))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := http.Client{
		Timeout: time.Second * 30,
	}

	// Get the response
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// Response body converted to stringified JSON
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Convert stringified JSON to a struct object of type githubAccessTokenResponse
	var ghresp githubAccessTokenResponse
	json.Unmarshal(respBody, &ghresp)

	// Return the access token (as the rest of the
	// details are relatively unnecessary for us)
	return ghresp.AccessToken, nil
}

type GithubUserResult struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Location string `json:"location"`
}

func GetGithubUserData(accessToken string) (GithubUserResult, error) {
	// Get request to a set URL
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return GithubUserResult{}, err
	}

	// Set the Authorization header before sending the request
	// Authorization: token XXXXXXXXXXXXXXXXXXXXXXXXXXX
	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)
	client := http.Client{
		Timeout: time.Second * 30,
	}

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return GithubUserResult{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return GithubUserResult{}, errors.New("could_not_retrieve_user")
	}

	// Read the response as a byte slice
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("string(respBody): %v\n", string(respBody))

	var githubUserResult GithubUserResult
	if err := json.Unmarshal(respBody, &githubUserResult); err != nil {
		return GithubUserResult{}, err
	}

	// Convert byte slice to string and return
	return githubUserResult, nil
}
