package main

import (
    "encoding/gob"
    "fmt"
    "log"
    "net"
    "os"
    //"time"
    "bufio"
    "strings"
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
        //log.Println("Error connecting to client socket:", err.Error())
        fmt.Println("Error connecting to client socket:", err.Error())
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
            //log.Println("This client already exists")
            fmt.Println("This client already exists")
            conn.Close()
        }
    }

    // Appends new client to Clients slice
    server.Clients = append(server.Clients, *client)
    //log.Printf("Client %s added to slice\n", client.Id)
    fmt.Printf("\nClient %s added to slice\n", client.Id)

    // Sends confirmation mesage to client so he knows that he was accepted
    server.sendConfirmationMessageToClient(client)
}

// Debug function
func showClients(server *Server) {
    fmt.Println("==============")
    for _, client := range server.Clients {
        fmt.Println(client)
    }
    fmt.Println("==============")
}

// Function that waits for stdin commands
func respondsToStdin(server *Server) {
    for {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("\ncommand > ")
        input, err := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        if err != nil{
            log.Println("Error while reading the user input")
        } else {
            // TODO: switch case to grab the command
            if input == "showClients" {
                showClients(server)
            } else if input == "exit" {
                fmt.Println("Going to exit ...")
                break
            }
        }

    }

}


// Starts server
func startServer(server *Server) {
    server.assignData()
    var err error

    server.ServerSocket, err = net.Listen(server.Type, server.Host+":"+server.Port)
    if err != nil {
        log.Println("Error listening:", err.Error())
    }

    defer server.ServerSocket.Close()

    //log.Println("\nListening on " + server.Host + ":" + server.Port)
    // Waits for new Client connections
    for {

        connection, err := server.ServerSocket.Accept()
        if err != nil {
            log.Println("Error accepting: ", err.Error())
        }
        go handleRegister(connection, server)
    }

}


// Function that handles the errors
func run() error {
    server := &Server{}

    go startServer(server) // starts listening for clients
    respondsToStdin(server) // responds to stdin commands

    return nil
}

func main() {
    if err := run(); err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
}
