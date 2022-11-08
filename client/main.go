package main

import (
    "encoding/gob"
    "fmt"
    "log"
    "net"
    "os"
)

// Application constants, defining host, port, and protocol.
const (
    connHost = "localhost"
    connPort = "9988"
    connType = "tcp"
)

type Client struct {
    Id string
}

// Function that handles the errors
func run() error {
    log.Println("Connecting to", connType, "server", connHost+":"+connPort)
    conn, err := net.Dial(connType, connHost+":"+connPort)

    if err != nil {
        fmt.Println("Error connecting:", err.Error())
        os.Exit(1)
    }

    encoder := gob.NewEncoder(conn)
    client := &Client{Id: "1"}
    encoder.Encode(client)

    // for {
    //     // TODO: create go routine that sends hearthbeat to server
    // }

    return nil
}

func main() {
    if err := run(); err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
}
