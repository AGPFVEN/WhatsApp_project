package main

import (
    "database/sql"
    "log"

     "github.com/go-sql-driver/mysql"
)

func dbTest() {
	//Connect
    db, err := sql.Open("mysql", credential1)
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }
    defer db.Close()

	//Verify connection
    if err := db.Ping(); err != nil {
        log.Fatalf("failed to ping: %v", err)
    }
    log.Println("Successfully connected to PlanetScale!")

	//Register file in Mysql system
    filename := "compress.zip"
    mysql.RegisterLocalFile(filename)

    //Creating query
    stmt, err := db.Prepare("INSERT INTO testDB1 (pnumber, pzip) VALUES (?, 0010101010)")
    if err != nil{
        panic(err)
    }
    pNumber := "000000000"

    //Execute query
    res, err := stmt.Exec(pNumber)
    if err != nil{
        panic(err)
    }

    println("Insert id: ", res)
}
