package main

import (
	"net/http"
)

func main(){
	http.HandleFunc("/", InitialPageLoader)
	//http.HandleFunc("/", my_test)

	http.ListenAndServe(":8000", nil)
}