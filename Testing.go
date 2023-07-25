package main

import (
	"context"
	"log"
	"net/http"

	"github.com/chromedp/chromedp"
)

func my_test(w http.ResponseWriter, r *http.Request) () {

	log.Println("Initializing Browser...")
	allocatorCtx, allocatorCancel := chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:], 
			chromedp.Flag("headless", false),
			chromedp.UserDataDir("myUsers"),
			)...
	)
	defer allocatorCancel()

	//Browser is closed at the end of this function
	_, browserCancel := chromedp.NewContext(allocatorCtx)
	defer browserCancel()

	//Testing string
	testString := "This is my test"

	//Next step of the process
	go HandlerRegistrationUpload(testString, allocatorCtx)
}