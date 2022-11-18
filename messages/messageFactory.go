package messages

import (
	"encoding/json"
	messagesBeh "syncute-go/messages/behavioral"
	"syncute-go/messages/resources"
)

func createBadJsonMessage(message messagesBeh.BadMessage) []byte {
	return []byte("")
}

func createGetAllResourcesJsonMessage() []byte {
	message := resources.CreateGetAllResourcesMessage()
	jsonMessage, _ := json.Marshal(message)
	return jsonMessage
}

func createDownloadResourcesJsonMessage(paths []string) []byte {
	message := resources.CreateDownloadResourcesMessage(paths)
	jsonMessage, _ := json.Marshal(message)
	return jsonMessage
}
