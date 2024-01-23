package main
import "net/http"

func main(){
	//dbTest()
	
	//Test: Using http.post can I send the zip file to the db?
	
	http.HandleFunc(webPagesHome, InitialPageLoaderTest)
	
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8000", nil)
}