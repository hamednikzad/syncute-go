package connections

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"syncute-go/messages"
	"time"
)

type Client struct {
	Token         string
	RemoteAddress string
	connection    *websocket.Conn
}

func (client Client) Start() {
	client.connect()

	client.consume()
}

func (client Client) connect() {
	fmt.Printf("Connecting to %s...\n", client.RemoteAddress)
	connection, _, err := websocket.DefaultDialer.Dial(client.RemoteAddress, http.Header{
		"access_token": []string{client.Token},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	client.connection = connection
}

func (client Client) consume() {
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
			messages.ProcessTextMessage(message)
			break
		case websocket.BinaryMessage:
			messages.ProcessBinaryMessage(message)
			break
		default:
			fmt.Println("Unknown")
		}
	}
}

func (client Client) send(message string) {
	fmt.Println("sending:: ", message)
	client.connection.WriteMessage(websocket.TextMessage, []byte(message))
}

func (client Client) sendSomeMessages() {
	client.send("Hello server, this is client")
	time.Sleep(2 * time.Second)
	client.send("What you up to?")
	time.Sleep(3 * time.Second)
	client.send("OK, bye now")
	time.Sleep(10 * time.Second)
}
