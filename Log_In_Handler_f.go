package main

import (
	"context"
	"log"
	"net/http"
	"text/template"

	"github.com/chromedp/chromedp"
)

type logInData struct{
	QrImage string
}

func LogInPageLoader(w http.ResponseWriter, r *http.Request) {
	//Retrive qr from what'sapp web page
	tmpQrPng, logInCtx := GetQrCode()
	
	//Qr data into the page data
	p := logInData{QrImage: tmpQrPng}

	//Load html file with qr code
	t, err := template.ParseFiles("log_in.html")
	if err != nil {
		print(err)
	}
	
	//Execute template into user browser
	if t.Execute(w, p) != nil{
		print(err)
	}

	//Return client's number
	ch := make(chan string)
	go RetriveNumber(logInCtx)
	clientsNumber := <-ch

	//clientsNumber := RetriveNumber(logInCtx)
	//clientsNumber := 
	//RetriveNumber(logInCtx)
	print(clientsNumber)

	//-------------------------test--------------------------
	//check if context works fine, see how to
	print(logInCtx)
}

func RetriveNumber(givenCtx context.Context) (string){
	//This function checks the number of the user using a channel
	contactCtx := SelectContact(givenCtx)

	var data map[string] string
	err := chromedp.Run(contactCtx,
		chromedp.Attributes("/html/body/div[1]/div/div/div[3]/div[1]/span/div/span/div/div[2]/div[5]/div/div/div[11]/div[1]/div/div[2]/div[1]/div/span[1]", &data),
	)
	if err != nil {
		log.Fatal(err)
	}

	return data["title"]
}

//This functions retrives the image of the qr code of the wss page
func GetQrCode() (string, context.Context) {
	//Initializing Browser Context (if headless mode is not disabled this doesn't work)
	execCtx, _ := chromedp.NewExecAllocator(
		context.Background(),
		append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))...
	)

	browserCtx, _ := chromedp.NewContext(execCtx)

	//var data map
	var data map[string]string

	log.Println("Initializing Browser...")

	err := chromedp.Run(browserCtx,
		chromedp.Navigate("http://web.whatsapp.com/"),
		chromedp.WaitEnabled("._10aH-", chromedp.ByQuery),
		chromedp.Attributes("/html/body/div[1]/div/div/div[3]/div[1]/div/div/div[2]/div", &data),
		)
	if err != nil {
		log.Fatal(err)
	}
	return data["data-ref"], browserCtx
}