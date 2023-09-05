package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
    currPacketSize = 30 * 1024 * 1024 //30MiB
)

func HandlerRegistrationUpload(phoneNumber string, isAllocatorClosed context.Context, isBrowserClosed context.Context) (){
	//Check if the browser is closed
	if (struct {}{} == <-isBrowserClosed.Done() && struct{}{} == <-isAllocatorClosed.Done()){
		println("Browser Closed")
	}

	//Compress browser sesion
	MyZip("compress.zip", "./myUsers")
}

func dbTest() {
    //Open File to read its content
    testNumber := "000000001"
    filename := "compress.zip"
    filePointer, err := os.Open(filename)
    if err != nil{
        log.Fatal("Error opening file\n", err)
    }
    defer filePointer.Close()

    //Create buffer to read chunks of the file
    buf:= make([]byte, currPacketSize)

    //Connect to DB
    db, err := sql.Open("mysql", os.Getenv("DSN"))	
    if err != nil {
		log.Fatal(err)
	}

	//Ping to test connection to DB
    if err := db.Ping(); err != nil {
        log.Fatalf("failed to ping: %v", err)
    }
    log.Printf("Successfully connected to PlanetScale!\n\n")


    //First read (this read uses insert)
    fileDescriptor, err := filePointer.Read(buf)
    if err != nil{
        log.Fatal(err)
    }
    println("First read completed")

    //Insert (I am reading how to use this insert in order to introduce the zip by chunks)
    stmt, err := db.Prepare("INSERT INTO testDB values (?, ?)")
    if err != nil{
        log.Fatal(err)
    }

    queryResult, err:= stmt.Exec(testNumber, buf)
    if err != nil{
        log.Fatal(err)
    }
    println("First query executed")
    log.Println(queryResult.RowsAffected())
    println()

    //Update (Upload file in chunks)
    for fileDescriptor > 0{
        fileDescriptor, err = filePointer.Read(buf)
        if err != nil{
            log.Fatal(err)
        }
        println("Another read completed")

        stmt, err = db.Prepare("update testDB1 set pzip = concat(pzip, ?) where pnumber = ?")
        if err != nil{
            log.Fatal(err)
        }
        queryResult, err:= stmt.Exec(buf, testNumber)
        if err != nil{
            log.Fatal(err)
        }
        log.Println(queryResult.RowsAffected())
        println()
    }

	/*
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

    println(stmt)
    log.Printf("Value of my b %v\n", )
    pNumber := "000000000"
    
    res, err := stmt.Exec(pNumber, b)
    if err != nil{
        log.Fatal(err)
    }
    println("Insert id: ", res)    
    */
}