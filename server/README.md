# Redis Server
A server application which imitates Redis. It uses an in memory map to store key-value pairs.

### Install Go
This project is implemented using Go. If Go is not yet installed, please download and install from [here](https://golang.org/doc/install)

### Run server
```bash
go mod tidy
go run .
```

### Commands
Allowed commands are `DUMP`, `SET`, `INCR`, `RENAME`, `HELP` and used like below :
```bash 
DUMP <key>
SET <key> <value>
INCR <key>
RENAME <key> <key>
HELP
```

### Run unit test
```bash
go to the folder `cd storage`, `cd config`, `cd commands`
go test -v
go test -v -run <function test name>
go test -cover
```