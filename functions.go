package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func SelectContact(givenCtx context.Context) (returnedCtx context.Context){
	//This function checks the number of the user	

	err := chromedp.Run(givenCtx,
		chromedp.WaitEnabled("._604FD", chromedp.ByQuery),
		chromedp.WaitReady("body"),
		chromedp.Click("/html/body/div[1]/div/div/div[4]/header/div[2]/div/span/div[3]/div/span"), 
		chromedp.WaitReady("body"),
		//chromedp.Attributes("/html/body/div[1]/div/div/div[3]/div[1]/span/div/span/div/div[2]/div[5]/div/div/div[11]/div[1]/div/div[2]/div[1]/div/span[1]", &data),
	)
	if err != nil {
		log.Fatal(err)
	}

	return givenCtx
}