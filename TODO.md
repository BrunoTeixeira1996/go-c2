# Todo

- Define a function that receive input and send to client so this is like an exec function and the function only gets a string and sends to client
    - Modify the `useClient` function

- Server could exit but clients should continue running and when server starts again, goes check what clients are still on

- Implement mutexes to control goroutines

- Implement command encription on server and command decription on client with one time key given on c2 server startup
