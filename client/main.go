package main

import (
    "bufio"
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

// Function that handles the errors
func run() error {
    log.Println("Connecting to", connType, "server", connHost+":"+connPort)
    conn, err := net.Dial(connType, connHost+":"+connPort)

    if err != nil {
        fmt.Println("Error connecting:", err.Error())
        os.Exit(1)
    }

    reader := bufio.NewReader(os.Stdin)

    // run loop forever, until exit.
    for {
        fmt.Print("Text to send: ")
        input, _ := reader.ReadString('\n')

        // Send to socket connection.
        conn.Write([]byte(input))

        // // Listen for relay.
        // message, _ := bufio.NewReader(conn).ReadString('\n')

        // // Print server relay.
        // log.Print("Server relay: " + message)
    }
}

func main() {
    if err := run(); err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
}
