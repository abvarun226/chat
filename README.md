# chat server

A very simple chat server using `github.com/gorilla/websocket`
* a user can send a private message to another user
* a user can broadcast a message to all users currently online

## Build

In order to build/run the server:
```
$ go run main.go
```

## Example
Use the `client/client.go` to connect to server and send a message.

```
$ cd client

# Start "user2" client
$ go run client.go "user2" "" ""

# Send a message "from user1" from "user1" to "user2"
$ go run client.go "user1" "user2" "from user1"
```

To broadcast a message to all connected users

```
$ cd client

# Start "user2" client
$ go run client.go "user2" "" ""

# Start "user3" client
$ go run client.go "user3" "" ""

# Broadcast a message "from user1" from "user1" to all connected users
$ go run client.go "user1" "GLOBAL" "from user1"
```