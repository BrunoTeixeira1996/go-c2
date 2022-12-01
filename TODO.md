# Todo

- Create db and register clients in there
    - Created, now I am trying to understand how to select, insert and delete (https://www.youtube.com/watch?v=YpDVQC8hfik)



```
Clients
---------
Id autoincrement PK
Hostname varchar
Os varchar
Arch varchar
IP varchar
Port varchar

Commands
---------
Id autoincrement PK
Id_client FK
Command varchar
Result varchar
Timestamp date
```

- Define commands to send to clients in a simple way

- Server could exit but clients should continue running and when server starts again, goes check what clients are still on

- Implement mutexes to control goroutines

- Implement command encription on server and command decription on client with one time key given on c2 server startup
