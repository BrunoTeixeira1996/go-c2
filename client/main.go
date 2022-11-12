package main

import (
    "encoding/gob"
    "fmt"
    "log"
    "net"
    "os"
    "runtime"
    "github.com/google/uuid"
)

const (
    SERVER_HOST = "localhost"
    SERVER_PORT = "9988"
    SERVER_TYPE = "tcp"
)

type Client struct {
    Id       string
    Hostname string
    Os       string
    Arch     string
    IP       string
    Port     string
}

func (cli *Client) generateRandomId() {
    id := uuid.New()
    cli.Id = id.String()
}

func (cli *Client) getClientOsAndArch() {
    cli.Os = runtime.GOOS
    cli.Arch = runtime.GOARCH
}

// TODO: Only handles linux for now
func (cli *Client) getHostName() {
    cli.Hostname, _ = os.Hostname()
}

func createNewCli() Client{
    // TODO: create random ports for now to work with multiple clients
    client := Client{IP: "localhost", Port: "9000"}
    client.generateRandomId()
    client.getClientOsAndArch()
    client.getHostName()

    return client
}

// Function to register on server socket
func registerOnServerSocket() (Client, error) {
    conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
    if err != nil {
        log.Println("Error connecting:", err.Error())
        return Client{}, err
    }
    log.Println("Connected with server")

    // Send client struct to register in server socket
    encoder := gob.NewEncoder(conn)
    client := createNewCli()
    encoder.Encode(client)

    return client, nil
}

// Function that verifies if client was registered
func verifyRegister(client Client)  error {
    // Get info from server to verify if we are registered
    clientSocket, err := net.Listen("tcp", client.IP+":"+client.Port)
    if err != nil {
        log.Println("Error listening:", err.Error())
        return err
    }

    defer clientSocket.Close()

    fmt.Println("Listening on client socket")

    connection, err := clientSocket.Accept()
    if err != nil {
        log.Println("Error accepting server message:", err.Error())
        return err
    }
    
    serverMsg := make([]byte, 1024)
    serverMsgLen, err := connection.Read(serverMsg)


    // TODO: make a new register message
    if string(serverMsg[:serverMsgLen]) != "REGISTERED" {
        return fmt.Errorf("Something went wrong and the client was not registered by server")
    }
    
    fmt.Println(string(serverMsg[:serverMsgLen]))

    return nil
}

// Function that handles the errors
func run() error {
    // Registers on server socket
    client, err := registerOnServerSocket()
    if  err != nil {
        return err
    }

    // Verifies if the register was well done
    if err := verifyRegister(client); err != nil {
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
