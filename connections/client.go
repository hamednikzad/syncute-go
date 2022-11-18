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

func (client *Client) consume() {
	fmt.Println("Waiting for new message...")

	for {
		messageType, message, err := client.connection.ReadMessage()

		if err != nil {
			return
		}

		switch messageType {
		case websocket.TextMessage:
			fmt.Println("server: text message >>", string(message))
			go messages.ProcessTextMessage(client.sendTextMessage, client.sendBinaryMessage, message)
			break
		case websocket.BinaryMessage:
			go messages.ProcessBinaryMessage(message)
			break
		default:
			fmt.Println("Unknown")
		}
	}
}

func (client *Client) sendTextMessage(message []byte) {
	fmt.Println("sending: ", string(message))
	client.connection.WriteMessage(websocket.TextMessage, message)
}

func (client *Client) sendBinaryMessage(message []byte) {
	fmt.Println("sending binary message: ")
	client.connection.WriteMessage(websocket.BinaryMessage, message)
}
