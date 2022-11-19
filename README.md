# SynCute-go
Golang client for SynCute
#
SynCute-go client is a Golang console application that connects to the server and syncs local directory with the server 
and receive new resource from other clients through the server. 
Clients are not connected to each other, they just connected to the server. 

## Run
You should edit config.yml file and set remote server address and token for connection to the server.
Run client with --help for more information.

## Test
Use this command for unit testing:
#### `go test ./...`
