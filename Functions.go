package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func SelectContact(givenCtx context.Context) (){
	//This function checks the number of the user
	err := chromedp.Run(givenCtx,
		chromedp.WaitReady("._3ndVb.fbgy3m38.ft2m32mm.oq31bsqd.nu34rnf1", chromedp.ByQuery),
		chromedp.WaitReady("body"),
		chromedp.Click("/html/body/div[1]/div/div/div[4]/header/div[2]/div/span/div[3]/div/span"),
					  ///html/body/div[1]/div/div/div[4]/header/div[2]/div/span/div[3]/div/span
		chromedp.WaitReady("body"),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Useful page ready")
}

	