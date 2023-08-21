# WebSocketTest

Very simple project to test extending life time of extension serviceworker by server-sent websocket messages

## Run the server

It's written in go, so you'll need go installed. Go 1.18 or later should do.

From a terminal, in {project-location}/server, do

* ```go build .```
* ```./websockettest```

The server uses tcp port 8080, so nothing else should run there

It sends the message "PING" every 10 seconds

Leave it running

## Extension
 
In chrome, load the extension from {project-location}/extension

After ~30 seconds the extension deactivates, despite recieving messages.

