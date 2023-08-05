package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func HandlerRegistrationUpload(phoneNumber string, isAllocatorClosed context.Context, isBrowserClosed context.Context) (){
	//Check if the browser is closed
	if (struct {}{} == <-isBrowserClosed.Done() && struct{}{} == <-isAllocatorClosed.Done()){
		println("Browser Closed")
	}

	println("continue")
	
	//TODO: quit this
	time.Sleep(1* time.Minute)

	MyZip("compress.zip", "./myUsers")
	println("hello there")
	MyUnzip("myUsers2", "compress.zip")

	//Test
	log.Println("Initializing Browser...")
	allocatorCtx, allocatorCancel := chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:], 
			chromedp.Flag("headless", false),
			chromedp.UserDataDir("myUsers2"),
			)...
	)
	defer allocatorCancel()

	//Browser is closed at the end of this function
	bor, browserCancel := chromedp.NewContext(allocatorCtx)
	defer browserCancel()
	erri := chromedp.Run(bor,
		chromedp.Navigate("http://web.whatsapp.com/"),
		chromedp.WaitEnabled("._10aH-", chromedp.ByQuery),
	)
	if erri != nil {
		log.Fatal(erri)
	}	
}