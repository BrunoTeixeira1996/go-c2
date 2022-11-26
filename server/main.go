package main

import (
    "encoding/gob"
    "fmt"
    "net"
    "os"
    "bufio"
    "strings"
    "log"
)

type Logger struct {
    infoLogger *log.Logger
    errorLogger *log.Logger
}

func (l *Logger) start(logFile *os.File) {
    l.infoLogger = log.New(logFile, "INFO: ", log.LstdFlags | log.Lshortfile)
    l.errorLogger = log.New(logFile, "ERROR: ", log.LstdFlags | log.Lshortfile)
}

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
func (server *Server) sendConfirmationMessageToClient(client *Client, logger *Logger) {
    clientConn, err := net.Dial("tcp", client.IP+":"+client.Port)
    if err != nil {
        logger.errorLogger.Println("Error connecting to client socket:", err.Error())
    } else {
        clientConn.Write([]byte("REGISTERED"))
    }
}

// Function that handles the registration phase
func handleRegister(conn net.Conn, server *Server, logger *Logger) {
    dec := gob.NewDecoder(conn)
    client := &Client{}
    dec.Decode(client)
    
    for _, registeredClient := range server.Clients{
        if client.Id == registeredClient.Id {
            logger.infoLogger.Println("This client already exists")
            conn.Close()
        }
    }

    // Appends new client to Clients slice
    server.Clients = append(server.Clients, *client)

    logger.infoLogger.Println("Client "+client.Id+" added to slice")

    // Sends confirmation mesage to client so he knows that he was accepted
    server.sendConfirmationMessageToClient(client, logger)
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
func respondsToStdin(server *Server, logger *Logger) {
    for {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("\ncommand > ")
        input, err := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        if err != nil{
            logger.errorLogger.Println("Error while reading the user input")
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
func startServer(server *Server, logger *Logger) {
    server.assignData()
    var err error

    server.ServerSocket, err = net.Listen(server.Type, server.Host+":"+server.Port)
    if err != nil {
        logger.errorLogger.Println("Error listening:", err.Error())
    }

    defer server.ServerSocket.Close()

    logger.infoLogger.Println("Listening on "+ server.Host + ":" + server.Port)

    // Waits for new Client connections
    for {

        connection, err := server.ServerSocket.Accept()
        if err != nil {
            logger.errorLogger.Println("Error accepting: ", err.Error())
        }

        go handleRegister(connection, server, logger)
    }

}

// Function that handles the errors
func run() error {
    logFile, err := os.OpenFile("/home/brun0/Desktop/workspace/go-c2/log/file.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666)
    if err != nil {
        log.Fatalf("Error while seting up the log file", err)
    }

    defer logFile.Close()

    logger := &Logger{}
    logger.start(logFile)

    server := &Server{}

    go startServer(server, logger) // starts listening for clients
    respondsToStdin(server, logger) // responds to stdin commands

    return nil
}

func main() {
    if err := run(); err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
}