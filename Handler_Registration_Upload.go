package main

import (
	"context"
	"database/sql"
	"fmt"
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

func ConnectDbWithSizePacket(myMaxAllowedPacket int) (*sql.DB ,error){
	//Connect to database
    myConfig, err := mysql.ParseDSN(os.Getenv("DSN"))
    if err != nil{
        log.Fatal(err)
    }

    myConfig.MaxAllowedPacket = myMaxAllowedPacket

    fmt.Printf("%+v\n", myConfig)

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
	//Register file in Mysql system
    filename := "compress.zip"

	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	db, err := ConnectDbWithSizePacket(len(b))
	if err != nil {
		log.Fatal(err)
	}

    //Creating query
    stmt, err := db.Prepare("SET GLOBAL max_allowed_packet=?;")
    if err != nil{
        log.Fatal(err)
    }

    //Execute query
    res, err := stmt.Exec(len(b))
    if err != nil{
        log.Fatal(err)
    }
    println("Insert id: ", res)

    stmt, err = db.Prepare("INSERT INTO testDB1 (pnumber, pzip) VALUES (?, ?)")
    if err != nil{
        log.Fatal(err)
    }
    pNumber := "000000000"

    res, err = stmt.Exec(pNumber, b)
    if err != nil{
        log.Fatal(err)
    }
    println("Insert id: ", res)
}