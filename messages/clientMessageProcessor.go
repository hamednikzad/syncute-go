package messages

import (
	"errors"
	"fmt"
	"syncute-go/helpers"
	"syncute-go/messages/resources"
)

func ProcessTextMessage(send func(message []byte), jsonMessage []byte) error {
	messageType, err := getMessageType(jsonMessage)
	if err != nil {
		return err
	}

	switch messageType {
	case BadMessageType:
		onBadMessage(jsonMessage, send)
		break
	case ReadyType:
		onReadyMessage(send)
		break
	case AllResourcesListType:
		onAllResourcesListMessage(jsonMessage)
		break
	//case DownloadResourcesType:
	//	onDownloadResourcesMessage(jsonMessage)
	//	break
	case NewResourceReceivedType:
		onNewResourceReceivedMessage(jsonMessage)
		break
	default:
		return errors.New("unknown message " + string(messageType))
	}

	return nil
}

func onNewResourceReceivedMessage(message []byte) {
	fmt.Printf("NewResourceReceivedMessage: %s", "") //message.Content.Resource);

	//await DownloadResources(new List<Resource>()
	//{
	//	new()
	//	{
	//		RelativePath = message.Content.Resource
	//	}
	//});
}

func onDownloadResourcesMessage(message []byte) {

}

func getRelativePath(resources []resources.Resource) []string {
	var result []string
	for i := range resources {
		result = append(result, resources[i].RelativePath)
	}
	return result
}

func onAllResourcesListMessage(message []byte) {
	allResourcesListMessage, err := parseAllResourcesListMessage(message)
	if err != nil {
		return
	}
	fmt.Println("AllResourcesListMessage received:")
	for i := range allResourcesListMessage.Content.Resources {
		fmt.Printf("Resource %d: %s\n", i, allResourcesListMessage.Content.Resources[i])
	}
	var serverResources = allResourcesListMessage.Content.Resources
	var serverRelativePaths = getRelativePath(serverResources)

	var localResources = helpers.GetAllFilesWithChecksum()
	var localRelativePaths = getRelativePath(localResources)

	var shouldDownloads = helpers.Difference(serverRelativePaths, localRelativePaths)
	var shouldUploads = helpers.Difference(localRelativePaths, serverRelativePaths)
	var intersects = helpers.Intersect(serverRelativePaths, localRelativePaths)

	fmt.Println("shouldDownloads", shouldDownloads)
	fmt.Println("shouldUploads", shouldUploads)
	fmt.Println("intersects", intersects)
	//foreach (var resource in intersects)
	//{
	//if (localResources.Any(r => r.RelativePath == resource.RelativePath && r.Checksum != resource.Checksum))
	//{
	//shouldDownloads.Add(resource);
	//}
	//}
	//
	//await UploadResources(shouldUploads);
	//
	//await DownloadResources(shouldDownloads);
}

func onReadyMessage(send func(message []byte)) {
	fmt.Println("Send GetAllResources message")
	var message = createGetAllResourcesJsonMessage()
	send(message)
}

func ProcessBinaryMessage(binary []byte) {

}

func onBadMessage(jsonMessage []byte, send func(message []byte)) {
	badMessage, err := parseBadMessage(jsonMessage)
	if err != nil {
		return
	}
	fmt.Printf("BadMessaged received: %s\n", badMessage.Content.Message)

	send(createBadJsonMessage(badMessage))
}
