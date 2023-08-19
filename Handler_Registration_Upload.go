package main

import (
	"context"
    "database/sql"
    "log"
    "os"

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

func ConnectDB() (*sql.DB ,error){
	//Connect to database
    db, err := sql.Open("mysql", os.Getenv("DSN"))
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }

	//Ping to test connection
    if err := db.Ping(); err != nil {
        log.Fatalf("failed to ping: %v", err)
    }

    log.Println("Successfully connected to PlanetScale!")

	return db, nil
}

func dbTest() {
	println("j")
	println(os.Getenv("DNS"))
	println("j")
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	//Register file in Mysql system
    filename := "compress.zip"

	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

    //Creating query
    stmt, err := db.Prepare("INSERT INTO testDB1 (pnumber, pzip) VALUES (?, ?)")
    if err != nil{
        log.Fatal(err)
    }
    pNumber := "000000000"

    //Execute query
    res, err := stmt.Exec(pNumber, b)
    if err != nil{
        panic(err)
    }

    println("Insert id: ", res)
}