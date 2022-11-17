package messages

import (
	"encoding/json"
	"fmt"
	"syncute-go/messages/resources"

	messagesBeh "syncute-go/messages/behavioral"
)

type MessageType string

const (
	BadMessageType          string = "bad_message"
	ReadyType               string = "ready"
	GetAllResourcesType     string = "get_resources"
	AllResourcesListType    string = "resources"
	DownloadResourcesType   string = "download"
	NewResourceReceivedType string = "new_resource"
)

func getMessageType(jsonMessage []byte) (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal(jsonMessage, &data)

	mType := data["Type"]
	if err != nil || mType == nil {
		fmt.Println(err)
		return string(BadMessageType), err
	}

	return mType.(string), nil
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

func parseAllResourcesListMessage(jsonMessage []byte) (resources.AllResourcesListMessage, error) {
	var message resources.AllResourcesListMessage
	err := json.Unmarshal(jsonMessage, &message)
	if err != nil {
		fmt.Println(err)
		return message, err
	}
	return message, nil
}
