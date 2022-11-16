package messages

import (
	"errors"
	"fmt"
)

func ProcessTextMessage(jsonMessage []byte) error {
	messageType, err := getMessageType(jsonMessage)
	if err != nil {
		return err
	}
	switch messageType {
	case BadMessage:
		badMessage, err := parseBadMessage(jsonMessage)
		if err != nil {
			return err
		}
		fmt.Println("It is bad message")
		fmt.Println(badMessage)
		break
	default:
		return errors.New("unknown message")
	}

	return nil
}

func ProcessBinaryMessage(binary []byte) {

}
