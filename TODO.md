# Todo

- Make the `use` command to use index based clients insead of only UID

- For now its only possible to connect over rev shell (it should be cooler to use a bind shell instead) on port 9991 (localhost obvs because this is for learning only)

- Server could exit but clients should continue running and when server starts again, goes check what clients are still on

- Implement mutexes to control goroutines

- Implement command encription on server and command decription on client with one time key given on c2 server startup

## Bugs

- Bug when using the command `shell` on server, for some reason it works on the first attempt but doesn't on the next ones
    - Looks like it works if I `use` -> `shell` -> `yes` -> `back` -> `use`
    - But it should work with
        - `use` -> `shell` -> `yes` -> `shell`
        - `use` -> `shell` -> `no` -> `shell` -> `yes`


- Bug when using a client because the output of the command goes after the next input
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

## Plan todo

