# Todo

- Working on `execCommand`, next is to send to client the command, execute and then wait for the response
    - This is working, next step is to wait for commands in client and printf those to see if its working (DONE)
    - Next is to understand why the client is only receiving the `use` command from the server instead of the correct command (DONE)
    - Next is to execute that command in client (DONE)
    - Next is to create web app in server that receives POST requests from client (DONE)
    - Next is to get the real command from the server + the output from the client + the timestamp and save that on the database when making the POST
    - Next is to encode the command and the result into b64

- Server could exit but clients should continue running and when server starts again, goes check what clients are still on

- Implement mutexes to control goroutines

- Implement command encription on server and command decription on client with one time key given on c2 server startup
