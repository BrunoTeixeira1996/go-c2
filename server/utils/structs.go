package utils

import (
    "net"
    "log"
    "os"
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


type Client struct {
    Id       string
    Hostname string
    Os       string
    Arch     string
    IP       string
    Port     string
    Uid      string
}