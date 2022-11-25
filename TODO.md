# Todo

- Make logs work to a log file and print only the necessary
  - Create log folder
    - Create file that logs registers and exists
    - Create file that logs commands per client with command + response (each client has a command_log.log file)

- Server could exit but clients should continue running and when server starts again, goes check what clients are still on

- Implement mutexes to control goroutines

- Implement command encription on server and command decription on client with one time key given on c2 server startup
