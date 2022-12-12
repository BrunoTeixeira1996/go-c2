package utils

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/mattn/go-sqlite3"
)


func setupDBConnection() *sql.DB {
    db, err := sql.Open("sqlite3", "/home/brun0/Desktop/go-c2/c2_db")
    if err != nil {
        log.Fatal(err)
    }
    return db
}


func SelectAllClientsQuery() {
    db := setupDBConnection()

    defer db.Close()

    rows, err := db.Query("SELECT * FROM Clients")

    if err != nil {
        log.Fatal(err)
    }

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

// TODO
func InsertNewClientQuery(client Client) {
    db := setupDBConnection()
    defer db.Close()

    statement, err := db.Prepare("INSERT INTO Clients (Hostname, Os, Arch, Ip, Port) VALUES (?, ?, ?, ?, ?)")

    if err != nil {
        log.Fatal(err)
    }

    statement.Exec(client.Hostname, client.Os, client.Arch, client.Ip, client.Port)
}


// TODO
func DeleteClientQuery() {

}