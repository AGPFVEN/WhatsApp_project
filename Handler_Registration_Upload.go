package main

import (
	"context"
)

func HandlerRegistrationUpload(phoneNumber string,isBrowserClosed context.Context) (){
	//Check if the browser is closed
	if (struct {}{} == <-isBrowserClosed.Done()){
		println("Browser Closed")
	}

	println("Continue program")
}