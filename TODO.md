# Todo

- Check the server code and check the TODO I was here, but basicaly I was searching for nc lib in go so I can connect to the rev shell

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
