# Todo

- Create db and register clients in there
  - Store commands and results sent from server to client

- Define commands to send to clients in a simple way

- Server could exit but clients should continue running and when server starts again, goes check what clients are still on

- Implement mutexes to control goroutines

- Implement command encription on server and command decription on client with one time key given on c2 server startup
