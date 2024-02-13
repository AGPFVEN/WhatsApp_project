package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/agpfven/WhatsApp_project/config"
)

//Only admin logs in
func OnedriveLogin(w http.ResponseWriter, r *http.Request){
	onedriveConfig := config.LoadOauthConfig()
	url := onedriveConfig.AuthCodeURL("randomstate")

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func OnedriveCallback(res http.ResponseWriter, req *http.Request){
	//State
	state := req.URL.Query()["state"][0]
	if state != "randomstate" {
		fmt.Fprintln(res, "States doesn't match")
		return
	}

	//Code
	code := req.URL.Query()["code"][0]

	//Configuration
	onedriveConfig := config.LoadOauthConfig()

	//Exchange code for token
	token, err := onedriveConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Fprintln(res, "Code-Token exchange failed")
	}

	//use token????
}