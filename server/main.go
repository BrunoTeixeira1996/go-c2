package main

import (
    "encoding/gob"
    "fmt"
    "log"
    "net"
    "os"
)

const (
    SERVER_HOST = "localhost"
    SERVER_PORT = "9988"
    SERVER_TYPE = "tcp"
)

type Client struct {
    Id string
}

//TODO:
// - Recieve Client struct from socket
// - Verify that this is a new Client
// - Add this client to a slice of clients
// - After all this, send confirmation to client socket
// - Client receive this confirmation and keeps listening on his socket

func handleConnection(conn net.Conn) {
    dec := gob.NewDecoder(conn)
    client := &Client{}
    dec.Decode(client)

    fmt.Printf("Received :%+v", client)

    conn.Close()
}

// Function that handles the errors
func run() error {
    fmt.Println("Server Running ... ")

    server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
    if err != nil {
        log.Println("Error listening:", err.Error())
        return err
    }

    defer server.Close()

    fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
    fmt.Println("Waiting for client...")

    for {
        connection, err := server.Accept()

        if err != nil {
            log.Println("Error accepting: ", err.Error())
            return err
        }

        fmt.Println("Client connected")
        go handleConnection(connection)
    }
}

func main() {
    if err := run(); err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
}
