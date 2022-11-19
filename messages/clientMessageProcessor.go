package messages

import (
	"errors"
	"log"
	"os"
	"syncute-go/helpers"
	"syncute-go/messages/resources"
)

func ProcessTextMessage(sendTextMessage func(message []byte), sendBinaryMessage func(message []byte), jsonMessage []byte) error {
	messageType, err := getMessageType(jsonMessage)
	if err != nil {
		return err
	}

	switch messageType {
	case BadMessageType:
		onBadMessage(jsonMessage, sendTextMessage)
		break
	case ReadyType:
		onReadyMessage(sendTextMessage)
		break
	case AllResourcesListType:
		onAllResourcesListMessage(jsonMessage, sendTextMessage, sendBinaryMessage)
		break
	case NewResourceReceivedType:
		onNewResourceReceivedMessage(jsonMessage, sendTextMessage)
		break
	default:
		return errors.New("unknown message " + string(messageType))
	}

	return nil
}

func onBadMessage(jsonMessage []byte, send func(message []byte)) {
	badMessage, err := parseBadMessage(jsonMessage)
	if err != nil {
		return
	}
	log.Printf("BadMessaged received: %s\n", badMessage.Content.Message)

	send(createBadJsonMessage(badMessage))
}

func onReadyMessage(send func(message []byte)) {
	log.Println("Send GetAllResources message")
	var message = createGetAllResourcesJsonMessage()
	send(message)
}

func onAllResourcesListMessage(message []byte, sendTextMessage func(message []byte), sendBinaryMessage func(message []byte)) {
	allResourcesListMessage, err := parseAllResourcesListMessage(message)
	if err != nil {
		return
	}
	log.Println("AllResourcesListMessage received:")
	for i := range allResourcesListMessage.Content.Resources {
		log.Printf("Resource %d: %s\n", i, allResourcesListMessage.Content.Resources[i])
	}
	var serverResources = allResourcesListMessage.Content.Resources

	var localResources = helpers.GetAllFilesWithChecksum()

	var shouldDownloads = helpers.DifferenceResources(serverResources, localResources)
	var shouldUploads = helpers.DifferenceResources(localResources, serverResources)
	var intersects = helpers.IntersectResources(serverResources, localResources)

	//log.Println("shouldDownloads", shouldDownloads)
	//log.Println("shouldUploads", shouldUploads)
	//log.Println("intersects", intersects)

	for i := range intersects {
		for j := range localResources {
			if localResources[j].RelativePath == intersects[i].RelativePath && localResources[j].Checksum != intersects[i].Checksum {
				shouldDownloads = append(shouldDownloads, intersects[i])
				break
			}
		}
	}

	UploadResources(shouldUploads, sendBinaryMessage)

	downloadResources(shouldDownloads, sendTextMessage)
}

func onNewResourceReceivedMessage(message []byte, sendTextMessage func(message []byte)) {
	newResourceReceivedMessage, err := parseNewResourceReceivedMessage(message)
	if err != nil {
		return
	}
	log.Printf("NewResourceReceivedMessage received: %s\n", newResourceReceivedMessage.Content.Resource)

	downloadResources([]resources.Resource{{
		RelativePath: newResourceReceivedMessage.Content.Resource,
	},
	}, sendTextMessage)
}

func getRelativePath(resources []resources.Resource) []string {
	var result []string
	for i := range resources {
		result = append(result, resources[i].RelativePath)
	}
	return result
}

func downloadResources(downloads []resources.Resource, send func(message []byte)) {
	if len(downloads) <= 0 {
		log.Println("There is nothing to download")
		return
	}

	send(createDownloadResourcesJsonMessage(getRelativePath(downloads)))
}

func UploadResource(resource resources.Resource, send func(message []byte)) {
	sendFile(resource, send)
}

func UploadResources(uploads []resources.Resource, send func(message []byte)) {
	if len(uploads) <= 0 {
		log.Println("There is nothing to upload")
		return
	}

	for i := range uploads {
		sendFile(uploads[i], send)
	}
}

func SerializeResource(resource resources.Resource) []byte {
	fileData, err := os.ReadFile(resource.FullPath)
	if err != nil {

	}
	var fileNameByte = []byte(resource.RelativePath)
	var clientData = make([]byte, 4+len(fileNameByte)+len(fileData))

	var fileNameLen = helpers.GetBytesOfUInt32(uint32(len(fileNameByte)))

	copy(clientData[0:4], fileNameLen)
	copy(clientData[4:len(fileNameByte)+4], fileNameByte)
	copy(clientData[4+len(fileNameByte):], fileData)

	return clientData
}

func sendFile(resource resources.Resource, send func(message []byte)) {
	data := SerializeResource(resource)

	send(data)
}

var CurrentResources []resources.Resource

func IsResourceExistAlready(fullPath string) bool {
	for i := range CurrentResources {
		if CurrentResources[i].FullPath == fullPath {
			return true
		}
	}
	return false
}

func ProcessBinaryMessage(binary []byte) {
	log.Printf("Binary message with size %d received\n", len(binary))

	newFilePath := helpers.WriteResource(binary)
	newResource := helpers.GetResourceByFullPath(newFilePath, "")

	CurrentResources = append(CurrentResources, newResource)
}
