package prototyping

import (
	"context"
	"os"

	"golang.org/x/oauth2"
)

var AppConfig oauth2.Config

func LoadConfig(){
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID: os.Getenv("ONEDRIVE_AUTH_ID"),
		ClientSecret: os.Getenv("ONEDRIVE_CLIENT_SECRET"),
		Scopes: []string{"files.readwrite", "offline_access"},
		Endpoint: oauth2.Endpoint{
			AuthURL: os.Getenv("ONEDRIVE_AUTH_ENDPOINT"),
			TokenURL: os.Getenv("ONEDRIVE_TOKEN_ENDPOINT"),
		},
	}
}