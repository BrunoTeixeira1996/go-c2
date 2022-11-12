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
    Id         string
    Hostname   string
    Os         string
    Arch       string
    IP         string
    Port       string
}

//TODO:
// - Register logic (Go routine that handles incoming registers and exits on Ctrl+C
    // - Verify that this is a new Client (DONE)
    // - Add this client to a slice of client (DONE)
    // - After all this, send confirmation to client socket (DONE)
    // - Work with multiple clients to test this (Need to randomize ports on client.go)
    // - Client receive this confirmation and keeps listening on his socket

// - Go routine that handles stdin to send comands to clients from server
// - Create log folder
    // - Create file that logs registers and exists
    // - Create file that logs commands per client with command + response (each client has a command_log.log file)

// - Implement mutexes to control goroutines
// - Implement command encription on server and command decription on client with one time key given on c2 server startup

func handleConnection(conn net.Conn) {
    dec := gob.NewDecoder(conn)
    client := &Client{}
    dec.Decode(client)

    fmt.Printf("Received :%+v", client)

    conn.Close()
}

func handleRegister(conn net.Conn, clientsSlice *[]Client) {
    dec := gob.NewDecoder(conn)
    client := &Client{}
    dec.Decode(client)

    for _, registeredClient := range *clientsSlice{
        if client.Id == registeredClient.Id {
            fmt.Println("This client already exists")
            conn.Close()
        }
    }

    *clientsSlice = append(*clientsSlice, *client)
    fmt.Printf("Client %s added to slice\n", client.Id)
    

    // Sends confirmation to client socket
    clientConn, err := net.Dial("tcp", client.IP+":"+client.Port)
    if err != nil {
        log.Println("Error connecting to client socket:", err.Error())
    } else {
        log.Println("Connected with client")
        clientConn.Write([]byte("REGISTERED"))
    }
}

func startRegisterSocket() error {
    server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
    if err != nil {
        log.Println("Error listening:", err.Error())
        return err
    }

    defer server.Close()

    fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)

    clientsSlice := &[]Client{}

    for {
        connection, err := server.Accept()

        if err != nil {
            log.Println("Error accepting: ", err.Error())
            return err
        }

        go handleRegister(connection, clientsSlice)
    }
}

// Function that handles the errors
func run() error {
    // Starts server register socket
    err := startRegisterSocket()
    if err != nil {
        return err
    }

    return nil
}

func main() {
    if err := run(); err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
}
