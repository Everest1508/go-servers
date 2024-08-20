# Go TCP Server
#### How to start go TCP Server
1. `cd tcp`
2. ` go run . `
3. Go server will start listening on `localhost` and on `8000` port.
4. Now your go TCP server is started ans ready to connect with client.

#### Conncet with TCP Server.
There are two way to connect with server and perform read write.

- Telnet
  1. Open terminal and type `telnet localhost 8000`.
  2. Now you are connected with server and can type any message and send to server.

- Go Clinet
  1. In main project there is Clinet Code written in Go.
  2. Type `cd client` and `go run .`
  3. It will execute and send message to server.

# Go HTTP server

#### How to start go HTTP server
1. `cd http`
2. `go run .`
3. Go HTTP server will start listening on `localhost` and `8000` port.
4. Open any browser and request on localhost:8000

