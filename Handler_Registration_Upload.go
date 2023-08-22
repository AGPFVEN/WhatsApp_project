package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func HandlerRegistrationUpload(phoneNumber string, isAllocatorClosed context.Context, isBrowserClosed context.Context) (){
	//Check if the browser is closed
	if (struct {}{} == <-isBrowserClosed.Done() && struct{}{} == <-isAllocatorClosed.Done()){
		println("Browser Closed")
	}

	//Compress browser sesion
	MyZip("compress.zip", "./myUsers")
}

//Idea: After creating a file check its size and use it in driver and session (All in handler)

func ConnectDB() (*sql.DB ,error){
	//Connect to database
    myConfig, err := mysql.ParseDSN(os.Getenv("DSN"))
    if err != nil{
        log.Fatal(err)
    }

    myConfig.MaxAllowedPacket = 90 << 20 //Less than 89.2MB

    println(myConfig.MaxAllowedPacket)
    
    myConn, err:= mysql.NewConnector(myConfig)
    if err != nil{
        log.Fatal(err)
    }

    db := sql.OpenDB(myConn)

	//Ping to test connection
    if err := db.Ping(); err != nil {
        log.Fatalf("failed to ping: %v", err)
    }

    log.Println("Successfully connected to PlanetScale!")

	return db, nil
}

func dbTest() {
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