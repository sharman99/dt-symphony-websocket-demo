# dt-symphony-websocket-demo

This example contains a Golang client that connects to any number of servers, prints out all messages received from servers, and sends out a message every second. This example contains a Node.js server and a Golang server.

When a server receives a message, it prints the message, then appends the server type (in this case, either "NODE SERVER" or "GO SERVER") to the message and returns it to the client. 

To run the example, start the servers:
    $ node nodejs-server.js
    $ cd websocket-demo
    $ go run go-server.go

Next, start the client:

    $ go run go-client.go