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

type Server struct {
    Clients      []Client
    ServerSocket net.Listener
    Host         string
    Port         string
    Type         string
}

func (server *Server) assignData() {
    server.Host = "localhost"
    server.Port = "9988"
    server.Type = "tcp"
}

type Client struct {
    Id         string
    Hostname   string
    Os         string
    Arch       string
    IP         string
    Port       string
}

func handleConnection(conn net.Conn) {
    dec := gob.NewDecoder(conn)
    client := &Client{}
    dec.Decode(client)

    fmt.Printf("Received :%+v", client)

    conn.Close()
}

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

    server.Clients = append(server.Clients, *client)
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

    for {
        connection, err := server.ServerSocket.Accept()

        if err != nil {
            log.Println("Error accepting: ", err.Error())
            return err
        }

        go handleRegister(connection, server)
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
