package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/chromedp/chromedp"
)


func main(){
	http.HandleFunc("/", LogInHandler)
	http.ListenAndServe(":8000", nil)
}

type logInData struct{
	QrImage string
}

func LogInHandler(w http.ResponseWriter, r *http.Request) {
	//Retrive qr from what'sapp web page
	tmpQrPng, logInCtx := GetQrCode()

	//-------------------------test--------------------------
	//check if context works fine, see how to

	//Qr data into the page data
	p := logInData{QrImage: tmpQrPng}

	//Load html file with qr code
	t, _ := template.ParseFiles("log_in.html")
	fmt.Println(t.Execute(w, p))

	//Return client's number
	//RetriveNumber(logInCtx)
	//clientsNumber := RetriveNumber(logInCtx)

	print(logInCtx)
}

func RetriveNumber(givenCtx context.Context) (){
	//This function checks the number of the user
	err := chromedp.Run(givenCtx,
		chromedp.WaitEnabled("._3ndVb", chromedp.ByQuery),
		//chromedp.Click("/html/body/div[1]/div/div/div[4]/header/div[2]/div/span/div[3]/div/span/svg"),
	)
	if err != nil {
		log.Fatal(err)
	}
	print("Waiting done")
	time.Sleep(5 * time.Second)
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