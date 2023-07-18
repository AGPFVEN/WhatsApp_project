package main

import (
	"net/http"
)

func main(){
	http.HandleFunc("/", LogInPageLoader)
	http.ListenAndServe(":8000", nil)
}