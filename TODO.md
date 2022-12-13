# Todo

- Working on `execCommand`, next is to send to client the command, execute and then wait for the response
    - This is working, next step is to wait for commands in client and printf those to see if its working (DONE)
    - Next is to understand why the client is only receiving the `use` command from the server instead of the correct command (DONE)
    - Next is to execute that command in client and send back the response to the server

- Server could exit but clients should continue running and when server starts again, goes check what clients are still on

- Implement mutexes to control goroutines

- Implement command encription on server and command decription on client with one time key given on c2 server startup
