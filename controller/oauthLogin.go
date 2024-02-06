package controller

import (
	"net/http"
)

func OnedriveLogin(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//Create oauthState cookie
	oauthState := uti
}