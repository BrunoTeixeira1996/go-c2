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

// Function that shows all clients from Clients table
func SelectAllClientsQuery() error {
    db := setupDBConnection()

    defer db.Close()

    rows, err := db.Query("SELECT * FROM Clients")

    if err != nil {
        return err
    }

    var client Client
    for rows.Next() {
        rows.Scan(
            &client.Hostname,
            &client.Os,
            &client.Arch,
            &client.IP,
            &client.Port,
            &client.Uid)

        // TODO: make a better format for this
        fmt.Printf("Id:%s Os: %s\n",client.Uid, client.Os)
    }

    return nil
}

// Function that inserts new clients into database
func InsertNewClientQuery(client Client) error {
    db := setupDBConnection()
    defer db.Close()

    statement, err := db.Prepare("INSERT INTO Clients (Hostname, Os, Arch, Ip, Port, Uid) VALUES (?, ?, ?, ?, ?, ?)")
    if err != nil {
        return err
    }

    _, err = statement.Exec(client.Hostname, client.Os, client.Arch, client.IP, client.Port, client.Uid)

    if err != nil {
        return err
    }

    return nil
}


// TODO
func DeleteClientQuery() {

}