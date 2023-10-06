package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
	"github.com/skip2/go-qrcode"
)

func main(){
	//Initializing Browser Context (if headless mode is not disabled this doesn't work)
	ctx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))...
	)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	//var data map
	var data map[string]string

	log.Println("Initializing Browser...")

	err := chromedp.Run(ctx,
		chromedp.Navigate("http://web.whatsapp.com/"),
		chromedp.WaitEnabled(".b77wc", chromedp.ByQuery),
		chromedp.Attributes("._2UwZ_", &data),
		)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data["data-ref"])

	qrcode.WriteFile(data["data-ref"], qrcode.Medium, 256, "qr.png")
}