package messages

import (
	"encoding/json"
	"fmt"

	messagesBeh "syncute-go/messages/behavioral"
)

type MessageType byte

const (
	Unknown    MessageType = iota
	BadMessage MessageType = iota
)

func getMessageType(jsonMessage []byte) (MessageType, error) {
	var data map[string]interface{}
	err := json.Unmarshal(jsonMessage, &data)

	if err != nil {
		fmt.Println(err)
		return Unknown, err
	}
	return BadMessage, nil
}

func parseBadMessage(jsonMessage []byte) (messagesBeh.BadMessage, error) {
	var message messagesBeh.BadMessage
	err := json.Unmarshal(jsonMessage, &message)
	if err != nil {
		fmt.Println(err)
		return message, err
	}
	return message, nil
}
