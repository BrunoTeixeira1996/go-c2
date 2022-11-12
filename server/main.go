package main

import (
    "encoding/gob"
    "fmt"
    "log"
    "net"
    "os"
    "time"
)

type Client struct {
    Id         string
    Hostname   string
    Os         string
    Arch       string
    IP         string
    Port       string
}

type Server struct {
    Host         string
    Port         string
    Type         string
    ServerSocket net.Listener
    Clients      []Client
}

// Starts boilerplate data on server
func (server *Server) assignData() {
    server.Host = "localhost"
    server.Port = "9988"
    server.Type = "tcp"
}

// Send message informing client that he was registered
func (server *Server) sendConfirmationMessageToClient(client *Client) {
    clientConn, err := net.Dial("tcp", client.IP+":"+client.Port)
    if err != nil {
        log.Println("Error connecting to client socket:", err.Error())
    } else {
        clientConn.Write([]byte("REGISTERED"))
    }
}

// Function that handles the registration phase
func handleRegister(conn net.Conn, server *Server) {
    dec := gob.NewDecoder(conn)
    client := &Client{}
    dec.Decode(client)
    
    for _, registeredClient := range server.Clients{
        if client.Id == registeredClient.Id {
            log.Println("This client already exists")
            conn.Close()
        }
    }

    // Appends new client to Clients slice
    server.Clients = append(server.Clients, *client)
    log.Printf("Client %s added to slice\n", client.Id)

    // Sends confirmation mesage to client so he knows that he was accepted
    server.sendConfirmationMessageToClient(client)
}

// Starts server
func startServer() error {
    var err error
    server := &Server{}
    server.assignData()
    
    server.ServerSocket, err = net.Listen(server.Type, server.Host+":"+server.Port)
    if err != nil {
        log.Println("Error listening:", err.Error())
        return err
    }

    defer server.ServerSocket.Close()

    log.Println("Listening on " + server.Host + ":" + server.Port)


    mainChannel:=make(chan struct{})
    defer close(mainChannel)

    go func() {
        for {
            select {
            case <-time.After(1*time.Second): // do action at interval
            case <-mainChannel: // main closed, time to stop
                return 
            }
            showClients(*server) // the action
        }
    }()

    // Waits for new Client connections
    for {
        connection, err := server.ServerSocket.Accept()
        if err != nil {
            log.Println("Error accepting: ", err.Error())
            return err
        }
        go handleRegister(connection, server)
    }

}

// Debug function
func showClients(server Server) {
    for _, client := range server.Clients {
        fmt.Println(client)
    }
}

// Function that handles the errors
func run() error {
    err := startServer()
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
