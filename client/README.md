# Redis-client
A client application which connects to redis server (A server application which imitates Redis)

### Run client
```bash
./mockredis-cli localhost 6379
```
### Commands
Allowed commands are `DUMP`, `SET`, `INCR`, `RENAME`, `QUIT`, `HELP` and used in following manner :
```bash
DUMP <key>
SET <key> <value>
INCR <key>
RENAME <key> <key>
QUIT
HELP
```