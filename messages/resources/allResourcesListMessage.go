package resources

const AllResourcesListType = "resources"

type AllResourcesListMessage struct {
	Type    string                  `json:"Type"`
	Content AllResourcesListContent `json:"Content"`
}

type AllResourcesListContent struct {
	Resources []Resource `json:"Resources"`
}

func CreateAllResourcesListMessage(resources []Resource) AllResourcesListMessage {
	return AllResourcesListMessage{
		Type: AllResourcesListType,
		Content: AllResourcesListContent{
			Resources: resources,
		},
	}
}
