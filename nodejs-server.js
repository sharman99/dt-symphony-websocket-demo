var WebSocketServer = require('ws').Server
  , wss = new WebSocketServer({ port: 8080 });

wss.on('connection', function connection(ws) {
  ws.on('message', function incoming(message) {
    console.log("recv from client: " + message.toString());
    ws.send("NODE SERVER: " + message);
    console.log("write to client: " + message.toString());
  });
});