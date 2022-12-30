# Todo

- For now its only possible to connect over rev shell (it would be cooler to use a bind shell instead) on port 9991 (localhost obvs because this is for learning only)

- Server could exit but clients should continue running and when server starts again, goes check what clients are still on

- Send/View commands via web UI (webserver that resided on server)

## Bugs

- Bug when trying to use more than one command
    - `ls -la` doesn't work
    - `cat something` doesn't work

- Bug when using the command `shell` on server, for some reason it works on the first attempt but doesn't on the next ones
    - If I use `shell` -> `no` after this, it never works
    - If I use another command before using `shell` It doesn't work

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
    - The workaround was a time.sleep()