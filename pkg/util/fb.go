package util

import (
	"encoding/json"
	"errors"
	"net/http"
	"personal/TODO-list/pkg/setting"

	"golang.org/x/oauth2"
	facebookOAuth "golang.org/x/oauth2/facebook"
)

// GetFacebookOAuthConfig will return the config to call facebook Login
func GetFacebookOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     setting.FbSetting.ClientId,
		ClientSecret: setting.FbSetting.ClientSecret,
		RedirectURL:  setting.FbSetting.RedirectUrl,
		Endpoint:     facebookOAuth.Endpoint,
		Scopes:       []string{"email"},
	}
}

// FacebookUserDetails is struct used for user details
type FacebookUserDetails struct {
	ID    string
	Name  string
	Email string
}

// GetUserInfoFromFacebook will return information of user which is fetched from facebook
func GetUserInfoFromFacebook(token string) (FacebookUserDetails, error) {
	var fbUserDetails FacebookUserDetails
	facebookUserDetailsRequest, _ := http.NewRequest("GET", "https://graph.facebook.com/me?fields=id,name,email&access_token="+token, nil)
	facebookUserDetailsResponse, facebookUserDetailsResponseError := http.DefaultClient.Do(facebookUserDetailsRequest)

	if facebookUserDetailsResponseError != nil {
		return FacebookUserDetails{}, errors.New("Error occurred while getting information from Facebook")
	}

	decoder := json.NewDecoder(facebookUserDetailsResponse.Body)
	decoderErr := decoder.Decode(&fbUserDetails)
	defer facebookUserDetailsResponse.Body.Close()

	if decoderErr != nil {
		return FacebookUserDetails{}, errors.New("Error occurred while getting information from Facebook")
	}

	return fbUserDetails, nil
}
