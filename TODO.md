# Todo

- Having a problem on `Accept()` on server since it blocks all goroutines
  - https://stackoverflow.com/questions/29948497/tcp-accept-and-go-concurrency-model

- Register logic (Go routine that handles incoming registers and exits on Ctrl+C)
  - Client receive this confirmation and keeps listening on his socket (Client receives this info but exits for now)

- Go routine that handles stdin to send comands to clients from server

- Create log folder
    - Create file that logs registers and exists
    - Create file that logs commands per client with command + response (each client has a command_log.log file)

- Implement mutexes to control goroutines

- Implement command encription on server and command decription on client with one time key given on c2 server startup
