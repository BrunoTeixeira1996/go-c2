# Todo

- Working on `execCommand`, next is to send to client the command, execute and then wait for the response

- Server could exit but clients should continue running and when server starts again, goes check what clients are still on

- Implement mutexes to control goroutines

- Implement command encription on server and command decription on client with one time key given on c2 server startup
