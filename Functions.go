package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func SelectContact(givenCtx context.Context) (returnedCtx context.Context){
	//This function checks the number of the user
	println("hit")
	err := chromedp.Run(givenCtx,
		chromedp.WaitEnabled("._604FD", chromedp.ByQuery),
		chromedp.WaitReady("body"),
		chromedp.Click("/html/body/div[1]/div/div/div[4]/header/div[2]/div/span/div[3]/div/span"), 
		chromedp.WaitReady("body"),
	)
	if err != nil {
		log.Fatal(err)
	}
	println("hit")

	return givenCtx
}