package resources

const AllResourcesListType = "resources"

type AllResourcesListMessage struct {
	Type    string                  `json:"Type"`
	Content AllResourcesListContent `json:"Content"`
}

type AllResourcesListContent struct {
	Resources []Resource `json:"Resources"`
}

type Resource struct {
	ResourceName string `json:"-"`
	FullPath     string `json:"-"`
	RelativePath string `json:"Path"`
	Checksum     string `json:"Checksum"`
}

func CreateAllResourcesListMessage(resources []Resource) AllResourcesListMessage {
	return AllResourcesListMessage{
		Type: AllResourcesListType,
		Content: AllResourcesListContent{
			Resources: resources,
		},
	}
}
