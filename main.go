package main

import (
	"syncute-go/connections"
)

func main() {
	//	rawJSON := []byte(`{
	//  "Type": "bad_message",
	//  "Content": {
	//    "Message": "Bad requestttt"
	//  }
	//}`)
	return
	client := connections.Client{
		RemoteAddress: "ws://localhost:5000/ws",
		Token:         "y3ocz7Aiv16jRRY4yfMmqVQvuV2wPuLSOO0HbpNE",
	}
	client.Start()
}
