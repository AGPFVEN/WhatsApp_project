package main

import (
	"context"
	"net/http"

	//"database/sql"
	"log"
	"os"

	"github.com/agpfven/WhatsApp_project/config"
)

//Env var OneDrive auth code (ONEDRIVE_AUTH_CODE)

func HandlerRegistrationUpload(phoneNumber string, isAllocatorClosed context.Context, isBrowserClosed context.Context) (){
	//Check if the browser is closed
	if (struct {}{} == <-isBrowserClosed.Done() && struct{}{} == <-isAllocatorClosed.Done()){
		println("Browser Closed")
	}

	//Compress browser sesion
	MyZip("compress.zip", "./myUsers")
}

func dbTest() { 
    print("https://login.live.com/oauth20_authorize.srf?client_id={"+
    os.Getenv("ONEDRIVE_AUTH_ID")+
    "}&scope={readwrite offline_access}&response_type=code&redirect_uri={" + 
    config.WebPagesHome + "}\n")
    //Getenv("ONEDRIVE_AUTH_ID")
    resp, err := http.Get("https://login.live.com/oauth20_authorize.srf?client_id={"+
        os.Getenv("ONEDRIVE_AUTH_ID")+
        "}&scope={readwrite offline_access}&response_type=code&redirect_uri={}")
    if err != nil{
        log.Fatal("Error requesting OneDrive authorization token\n", err)
    }

    log.Printf("Type: %T, value: %v", resp, resp)

    /*
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
    //_, err = filePointer.Read(buf)
    if err != nil{
        log.Fatal(err)
    }
    println("First read completed")

    //Insert (I am reading how to use this insert in order to introduce the zip by chunks)
    stmt, err := db.Prepare("INSERT INTO testDB1 values (?, ?)")
    if err != nil{
        log.Fatal(err)
    }

    queryResult, err:= stmt.Exec(testNumber, 010)
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
    }*/

    //Use between runs
    //delete from testDB1 where pnumber = "000000001";

    /*
    Investigate

    Console:
    2023/09/18 12:43:35 Successfully connected to PlanetScale!

    First read completed
    First query executed
    2023/09/18 12:43:35 1 <nil>

    Another read completed
    2023/09/18 12:43:42 Error 1105 (HY000): target: wss_test.-.primary: vttablet: rpc error: code = Unavailable desc = error reading from server: EOF
    */

    /*
    //Testing
    stmt, err = db.Prepare("update testDB1 set pzip = concat(pzip, ?) where pnumber = ?")
    if err != nil{
        log.Fatal(err)
    }
    queryResult, err = stmt.Exec(010, testNumber)
    if err != nil{
        log.Fatal(err)
    }
    log.Println(queryResult.RowsAffected())
    println()
    */
}