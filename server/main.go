package main

import (
    "encoding/gob"
    "fmt"
    "net"
    "os"
    "bufio"
    "strings"
    "log"
    "go-c2/server/utils"
)

const help = `help -> shows help
showClients -> shows connected available clients
exit -> exits the server
`

type Logger struct {
    infoLogger *log.Logger
    errorLogger *log.Logger
}

func (l *Logger) start(logFile *os.File) {
    l.infoLogger = log.New(logFile, "INFO: ", log.LstdFlags | log.Lshortfile)
    l.errorLogger = log.New(logFile, "ERROR: ", log.LstdFlags | log.Lshortfile)
}


type Server struct {
    Host         string
    Port         string
    Type         string
    ServerSocket net.Listener
    Clients      []utils.Client
}

// Starts boilerplate data on server
func (server *Server) assignData() {
    server.Host = "localhost"
    server.Port = "9988"
    server.Type = "tcp"
}

// Send message informing client that he was registered
func (server *Server) sendConfirmationMessageToClient(client *utils.Client, logger *Logger) {
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
    client := &utils.Client{}
    dec.Decode(client)

    err := utils.InsertNewClientQuery(*client)
    if err != nil {
        log.Fatal(err)
    }

    logger.infoLogger.Println("Client "+client.Id+" added to database")

    // Sends confirmation mesage to client so he knows that he was accepted
    server.sendConfirmationMessageToClient(client, logger)
}

// Function to show clients
func showClients(server *Server) {
    fmt.Println("==============")
    for _, client := range server.Clients {
        fmt.Println(client)
    }
    fmt.Println("==============")
}

// Function to execute commands on client
func useClient(input string, logger *Logger) {
    //TODO: send command to client
    // use <client id>
    // if back exit this, otherwise its an infinite loop here
  for {
        if len(strings.Split(input, " ")) == 1 {
            fmt.Println("Please provide the client id to execute commands")
            break
        }
        reader := bufio.NewReader(os.Stdin)
        fmt.Printf("\ncommand in client > ")
        commandToClient, err := reader.ReadString('\n')
        commandToClient = strings.TrimSpace(commandToClient)

        if err != nil {
            logger.errorLogger.Println("Error while reading the user input")
            break
        } else {
            // TODO: verify if the client exists
            // to do this, I should register the client to a db and then check the client status
            // after that I could keep working on this
            if commandToClient == "back" {
                break
            } else {
                fmt.Println("Sent " + commandToClient + " to client")
            }
        }
    }
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
            // Working on this
            switch {

            case strings.Contains(input,"help"):
                fmt.Printf(help)

            case strings.Contains(input, "showClients"):
                showClients(server)

            case strings.HasPrefix(input, "use"):
                useClient(input, logger)

            case strings.Contains(input, "exit"):
                fmt.Println("Going to exit...")
                return
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
    logFile, err := os.OpenFile("/home/brun0/Desktop/go-c2/log/file.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666)
    if err != nil {
        log.Fatalf("Error while seting up the log file %s", err)
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