package main

import (
	"context"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/agpfven/WhatsApp_project/config"
	"github.com/chromedp/chromedp"
)

type logInData struct{
	QrImage string
}

func InitialPageLoaderTest(w http.ResponseWriter, r *http.Request) {
	//Load html loading file
	page_template, err := template.ParseFiles("testing.html")
	if err != nil { 
		print(err)
	}
	
	//Execute template into user browser
	if page_template.Execute(w, nil) != nil{
		print(err)
	}
}

func InitialPageLoader(w http.ResponseWriter, r *http.Request) {
	qrData := make(chan string)

	//Retrive qr from what'sapp web page
	go RegistrationDataHandler(qrData)

	//Qr data into the page data
	p := logInData{QrImage: <-qrData}
	close(qrData)

	//Load html file with qr code
	t, err := template.ParseFiles("log_in.html")
	if err != nil {
		print(err)
	}
	
	//Execute template into user browser
	if t.Execute(w, p) != nil{
		print(err)
	}
}

func RegistrationDataHandler(ch chan string) (){
	//Initializing Browser Context (if headless mode is not disabled this doesn't work)
	log.Println(config.PromptStartBrowser)
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
	browserCtx, browserCancel := chromedp.NewContext(allocatorCtx)
	defer browserCancel()

	//Extract QR data from wss page
	GetQrCode(ch, browserCtx)

	//Retrive User's phone number
	userPhoneNumber := RetriveNumber(browserCtx)
	log.Printf("Users phone number: %s", userPhoneNumber)

	//This is done in order to let the whatsapp web page to synchronize with the mobile app
	time.Sleep(1 * time.Minute)

	chromedp.Cancel(browserCtx)

	//Next step of the process
	go HandlerRegistrationUpload(userPhoneNumber, allocatorCtx, browserCtx)
}

//This function retrives the user phone number
func RetriveNumber(givenCtx context.Context) (string){
	//This function checks the number of the user using a channel
	SelectContact(givenCtx)

	var data map[string] string
	err := chromedp.Run(givenCtx,
		chromedp.Attributes("/html/body/div[1]/div/div/div[3]/div[1]/span/div/span/div/div[2]/div[5]/div/div/div[11]/div[1]/div/div[2]/div[1]/div/span[1]", &data),
	)
	if err != nil {
		log.Fatal(err)
	}

	return data["title"]
}

//This functions retrives the image of the qr code of the wss page
func GetQrCode(auxiliarCh chan string, browserCtx context.Context) () {
	log.Println("Extracting QR data...")

	//Where the attributes data will be stored
	var data map[string]string

	//Go to WSS webpage, wait for QR and extract its information
	err := chromedp.Run(browserCtx,
		chromedp.Navigate("http://web.whatsapp.com/"),
		chromedp.WaitEnabled("._10aH-", chromedp.ByQuery),
		chromedp.Attributes("/html/body/div[1]/div/div/div[3]/div[1]/div/div/div[2]/div", &data),
		)
	if err != nil {
		log.Fatal(err)
	}

	//Pass the QR data information to the channel
	auxiliarCh <- data["data-ref"]
}