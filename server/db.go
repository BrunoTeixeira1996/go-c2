package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/mattn/go-sqlite3"
)


type Client struct {
    Id         int
    Hostname   string
    Os         string
    Arch       string
    Ip         string
    Port       string
}


func testQuery() {
    db, err := sql.Open("sqlite3", "/home/brun0/Desktop/workspace/go-c2/c2_db")
    
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

    rows, _ := db.Query("SELECT * FROM Clients")

    var client Client
    for rows.Next() {
        rows.Scan(
            &client.Id,
            &client.Hostname,
            &client.Os,
            &client.Arch,
            &client.Ip,
            &client.Port)

        fmt.Printf("Id:%d Os: %s\n",client.Id, client.Os)
    }
}

func main() {
    testQuery()
}
