package main

//import "net/http"

//func main(){
////dbTest()

////Test: Using http.post can I send the zip file to the db?

//http.HandleFunc(webPagesHome, InitialPageLoaderTest)

////http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

//http.ListenAndServe(":8000", nil)
//}
import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/goh-chunlin/go-onedrive/onedrive"
	"golang.org/x/oauth2"
)

 func main(){
	//Request token code----------------------------------------
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
	//// Redirect user to consent page to ask for permission
	//// for the scopes specified above.
	//url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	//fmt.Printf("Visit the URL for the auth dialog: %v", url)

	//Use token code-------------------------------------------
	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	print("hi\n")
	code := os.Getenv("ONEDRIVE_TOKEN_REDEEM_CODE")
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}

	// Use the custom HTTP client when requesting a token.
	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	print("Exchange done\n")

	client := conf.Client(ctx, tok)
	fmt.Println("It worked")

	//----------------------------OneDrive Mine------------------
	lib_client := onedrive.NewClient(client)
	log.Printf("CONTEXT %+V", ctx)
	somenoi, err := lib_client.DriveItems.List(ctx, "")
	if err != nil{
		log.Fatal(err)
	}
	print(somenoi.DriveItems)
}