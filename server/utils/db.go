package utils

import (
	"database/sql"
	"fmt"
	"log"
    "strings"
	_ "github.com/mattn/go-sqlite3"
)

// Function to setup a db connection
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
            &client.Id,
            &client.Hostname,
            &client.Os,
            &client.Arch,
            &client.IP,
            &client.Port,
            &client.Uid)

        // TODO: make a better format for this
        fmt.Printf("Id: %s Uid:%s Os: %s\n", client.Id, client.Uid, client.Os)
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

// Function to check if client exists in database
func CheckClientExistence(id string) (Client, error) {
    db := setupDBConnection()
    defer db.Close()

    var client Client
    var query string

    // check if is Id or Uid
    if strings.Contains(id, "-") {
        query = "SELECT * FROM Clients WHERE Uid = ?"
    } else {
        query = "SELECT * FROM Clients WHERE Id = ?"
    }

    if err := db.QueryRow(query, id).
            Scan(&client.Id,
                 &client.Hostname,
                 &client.Os,
                 &client.Arch,
                 &client.IP,
                 &client.Port,
                 &client.Uid); err != nil {

         if err == sql.ErrNoRows {
            return Client {}, fmt.Errorf("There are no Clients with that Uid (%s)\n", id)
        }

        return Client{}, err
    }

    return client, nil
}


// TODO
func DeleteClientQuery() {

}