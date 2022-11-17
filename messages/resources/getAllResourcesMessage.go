package resources

const GetAllResourcesType = "get_resources"

type GetAllResourcesMessage struct {
	Type string `json:"Type"`
}

func CreateGetAllResourcesMessage() GetAllResourcesMessage {
	return GetAllResourcesMessage{
		Type: GetAllResourcesType,
	}
}
