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
    Hostname string
    Os string
    Arch string
    // IP
    // PORT
}

//TODO:
// CHECK EXCALIDRAW DRAWING
// - Register logic (Go routine that handles incoming registers and exits on Ctrl+C (create a channel ?))
    // - Verify that this is a new Client
    // - Add this client to a slice of clients
    // - After all this, send confirmation to client socket
    // - Client receive this confirmation and keeps listening on his socket

// - Go routine that handles stdin to send comands to clients from server
// - Create log folder
    // - Create file that logs registers and exists
    // - Create file that logs commands per client with command + response (each client has a command_log.log file)

// - Implement command encription on server and command decription on client with one time key given on c2 server startup

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
