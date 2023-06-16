package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"personal/TODO-list/pkg/setting"
	"time"

	"golang.org/x/oauth2"
	googleOAuth "golang.org/x/oauth2/google"
)

// GetGoogleOAuthConfig will return the config to call google Login
func GetGoogleOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     setting.GmailSetting.ClientId,
		ClientSecret: setting.GmailSetting.ClientSecret,
		RedirectURL:  setting.GmailSetting.RedirectUrl,
		Endpoint:     googleOAuth.Endpoint,
		Scopes:       []string{"email"},
	}
}

type GoogleOauthToken struct {
	Access_token string
	Id_token     string
}

func GetGoogleOauthToken(code string) (GoogleOauthToken, error) {
	const rootURl = "https://oauth2.googleapis.com/token"

	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("client_id", setting.GmailSetting.ClientId)
	values.Add("client_secret", setting.GmailSetting.ClientSecret)
	values.Add("redirect_uri", setting.GmailSetting.RedirectUrl)
	query := values.Encode()

	req, err := http.NewRequest("POST", rootURl, bytes.NewBufferString(query))
	if err != nil {
		return GoogleOauthToken{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return GoogleOauthToken{}, err
	}

	if res.StatusCode != http.StatusOK {
		return GoogleOauthToken{}, errors.New("could not retrieve token")
	}

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return GoogleOauthToken{}, err
	}

	var GoogleOauthTokenRes map[string]interface{}

	if err := json.Unmarshal(resBody.Bytes(), &GoogleOauthTokenRes); err != nil {
		return GoogleOauthToken{}, err
	}

	tokenBody := GoogleOauthToken{
		Access_token: GoogleOauthTokenRes["access_token"].(string),
		Id_token:     GoogleOauthTokenRes["id_token"].(string),
	}

	return tokenBody, nil
}

type GoogleUserResult struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
}

func GetGoogleUserData(accessToken string, idToken string) (GoogleUserResult, error) {
	rootUrl := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", accessToken)

	req, err := http.NewRequest("GET", rootUrl, nil)
	if err != nil {
		return GoogleUserResult{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", idToken))
	client := http.Client{
		Timeout: time.Second * 30,
	}

	resp, err := client.Do(req)
	if err != nil {
		return GoogleUserResult{}, err
	}

	respBody, _ := ioutil.ReadAll(resp.Body)

	var googleUserRes GoogleUserResult
	if err := json.Unmarshal(respBody, &googleUserRes); err != nil {
		return GoogleUserResult{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return GoogleUserResult{}, errors.New("could_not_retrieve_user")
	}

	return googleUserRes, nil
}
