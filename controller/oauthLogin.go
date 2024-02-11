package controller

import (
	"net/http"

	"github.com/agpfven/WhatsApp_project/config"
)

func OnedriveLogin(w http.ResponseWriter, r *http.Request){
	onedriveConfig := config.LoadOauthConfig()
	url := onedriveConfig.AuthCodeURL("randomstate")

	http.Redirect(w, r, url, http.StatusSeeOther)
}