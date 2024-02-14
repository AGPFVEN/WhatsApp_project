package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/agpfven/WhatsApp_project/config"
	"github.com/goh-chunlin/go-onedrive/onedrive"
	"golang.org/x/oauth2"
)

//Only admin logs in
func OnedriveLogin(w http.ResponseWriter, r *http.Request){
	onedriveConfig := config.LoadOauthConfig()
	url := onedriveConfig.AuthCodeURL("randomstate")

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func OnedriveCallback(res http.ResponseWriter, req *http.Request){
	// state
	state := req.URL.Query()["state"][0]
	if state != "randomstate" {
		fmt.Fprintln(res, "States doesn't match")
		return
	}

	// code
	code := req.URL.Query()["code"][0]

	// configuration
	onedriveConfig := config.LoadOauthConfig()

	// exchange code for token
	token, err := onedriveConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Fprintln(res, "Code-Token exchange failed")
	}

	// use oneDrive API to get admin info
	ctx:= context.Background()
	ts := oauth2.StaticTokenSource(token)
	tc := oauth2.NewClient(ctx, ts)
	client_od := onedrive.NewClient(tc)

	drives, err := client_od.Drives.List(ctx)
	if err != nil{
		fmt.Println("Error at listing")
	}

	fmt.Printf("drives: %v\n", drives)
}