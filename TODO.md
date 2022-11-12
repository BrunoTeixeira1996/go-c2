# Todo

- Create server struct that has slice of clients + ip, port and server type (DONE)
  - Register client socket

- Register logic (Go routine that handles incoming registers and exits on Ctrl+C)
  - Client receive this confirmation and keeps listening on his socket 
  - Work with multiple clients to test this

- Go routine that handles stdin to send comands to clients from server

- Create log folder
    - Create file that logs registers and exists
    - Create file that logs commands per client with command + response (each client has a command_log.log file)

- Implement mutexes to control goroutines

- Implement command encription on server and command decription on client with one time key given on c2 server startup
