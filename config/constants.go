package config

import (
	"os"

	"golang.org/x/oauth2"
)

const(
	//Web pages
	WebPagesHome = "/"

	//Prompts
	PromptStartBrowser = "Initializing Browser..."
)

type product struct{
	phone string
	zip_bytes []byte
}

func LoadOauthConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID: os.Getenv("ONEDRIVE_AUTH_ID"),
		ClientSecret: os.Getenv("ONEDRIVE_CLIENT_SECRET"),
		Scopes: []string{"files.readwrite", "offline_access"},
		Endpoint: oauth2.Endpoint{
			AuthURL: os.Getenv("ONEDRIVE_AUTH_ENDPOINT"),
			TokenURL: os.Getenv("ONEDRIVE_TOKEN_ENDPOINT"),
		},
	}

	return conf
}