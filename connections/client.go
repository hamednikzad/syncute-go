package connections

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"syncute-go/messages"
)

type Client struct {
	Token         string
	RemoteAddress string
	connection    *websocket.Conn
}

func (client *Client) Start() {
	client.connect()

	client.consume()
}

func (client *Client) connect() {
	fmt.Printf("Connecting to %s...\n", client.RemoteAddress)
	var err error
	client.connection, _, err = websocket.DefaultDialer.Dial(client.RemoteAddress, http.Header{
		"access_token": []string{client.Token},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}
func Send() {

}
func (client *Client) consume() {
	fmt.Println("Waiting for new message...")
	// Must be started in a go routine
	for {
		// I exit the loop upon error here so the program doesn't panic when one side closes the connection
		messageType, message, err := client.connection.ReadMessage()

		if err != nil {
			return
		}
		fmt.Println("server says >>", string(message))

		switch messageType {
		case websocket.TextMessage:
			messages.ProcessTextMessage(client.send, message)
			break
		case websocket.BinaryMessage:
			messages.ProcessBinaryMessage(message)
			break
		default:
			fmt.Println("Unknown")
		}
	}
}

func (client *Client) send(message []byte) {
	fmt.Println("sending:: ", string(message))
	client.connection.WriteMessage(websocket.TextMessage, message)
}
