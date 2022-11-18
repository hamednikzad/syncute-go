package resources

const DownloadResourcesType = "download"

type DownloadResourcesMessage struct {
	Type    string                   `json:"Type"`
	Content DownloadResourcesContent `json:"Content"`
}

type DownloadResourcesContent struct {
	Resources []string `json:"Resources"`
}

func CreateDownloadResourcesMessage(resources []string) DownloadResourcesMessage {
	return DownloadResourcesMessage{
		Type: DownloadResourcesType,
		Content: DownloadResourcesContent{
			Resources: resources,
		},
	}
}
