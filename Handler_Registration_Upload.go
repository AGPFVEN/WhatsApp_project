package main

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
)

func HandlerRegistrationUpload(phoneNumber string, isAllocatorClosed context.Context, isBrowserClosed context.Context) (){
	//Check if the browser is closed
	if (struct {}{} == <-isBrowserClosed.Done() && struct{}{} == <-isAllocatorClosed.Done()){
		println("Browser Closed")
	}

	//Compress browser sesion
	MyZip("compress.zip", "./myUsers")
}