package utils

import (
	"database/sql"
	"fmt"
	"log"

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
        fmt.Printf("Uid:%s Os: %s\n",client.Uid, client.Os)
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
func CheckClientExistence(clientUid string) error {
    db := setupDBConnection()
    defer db.Close()

    var client Client

    if err := db.QueryRow("SELECT * FROM Clients WHERE Uid = ?", clientUid).
            Scan(&client.Id,
                 &client.Hostname,
                 &client.Os,
                 &client.Arch,
                 &client.IP,
                 &client.Port,
                 &client.Uid); err != nil {

         if err == sql.ErrNoRows {
            return fmt.Errorf("There are no Clients with that Uid (%s)\n", clientUid)
        }

        return err
    }

    return nil
}


// TODO
func DeleteClientQuery() {

}