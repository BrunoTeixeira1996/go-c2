# Todo

- Wait for `stdin` and for `Accept`
  - https://gist.github.com/mmirolim/fe3a77eb29af192ca9de
  
- Register logic (Go routine that handles incoming registers and exits on Ctrl+C)
  - Client receive this confirmation and keeps listening on his socket (Client receives this info but exits for now so I added a Scanf to prevent the exit)

- Create log folder
    - Create file that logs registers and exists
    - Create file that logs commands per client with command + response (each client has a command_log.log file)

- Implement mutexes to control goroutines

- Implement command encription on server and command decription on client with one time key given on c2 server startup
