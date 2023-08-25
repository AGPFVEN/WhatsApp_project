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
    filename := "compress.zip"
    fileDescriptor, err := os.Open(filename)
    if err != nil{
        log.Fatal("Error opening file\n", err)
    }
    defer fileDescriptor.Close()

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
    log.Println("Successfully connected to PlanetScale!")


    //First read (this read uses insert)
    readFile, err := fileDescriptor.Read(buf)

    //Insert
    stmt, err := db.Prepare("INSERT INTO testDB")

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