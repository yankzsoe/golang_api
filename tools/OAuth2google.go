package tools

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var Configs oauth2.Config

func GenerateOAuthGoogleUrl() string {
	Config := &oauth2.Config{
		ClientID:     "1069137797205-9f314ntntd3ac3dl144rbl6941q8cda6.apps.googleusercontent.com",
		ClientSecret: "AJEYTAbiHyZ7w3OyCKNYAFYL",
		RedirectURL:  "https://localhost:5001/api/v1/auth/external/google-callback",
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
	authURL := Config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	Configs = *Config
	return authURL
}

func GetUserInfo(code string) UserInfoResponse {
	token, err := Configs.Exchange(context.Background(), code)
	if err != nil {
		panic(err)
	}

	client := Configs.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")

	if err != nil {
		panic(err)
	}

	var data UserInfoResponse

	defer resp.Body.Close()

	// Baca respons body sebagai []byte
	userInfo, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading user info:", err)
		panic(err)
	}
	err = json.Unmarshal(userInfo, &data)
	if err != nil {
		panic(err)
	}

	return data
}

type UserInfoResponse struct {
	Id             string `json:"id"`
	Email          string `json:"email"`
	Verified_Email bool   `json:"verified_email"`
	Name           string `json:"name"`
	Given_Name     string `json:"given_name"`
	Family_Name    string `json:"family_name"`
	Picture        string `json:"picture"`
	Locale         string `json:"locale"`
	Hd             string `json:"hd"`
}
