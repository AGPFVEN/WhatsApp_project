package main

import (
	"net/http"
)

func main(){
	http.HandleFunc("/", InitialPageLoader)
	http.ListenAndServe(":8000", nil)
}