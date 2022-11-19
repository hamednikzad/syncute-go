package resources

const NewResourceReceivedType = "new_resource"

type NewResourceReceivedMessage struct {
	Type    string                     `json:"Type"`
	Content NewResourceReceivedContent `json:"Content"`
}

type NewResourceReceivedContent struct {
	Resource string `json:"Resource"`
}

func CreateNewResourceReceivedMessage(resource string) NewResourceReceivedMessage {
	return NewResourceReceivedMessage{
		Type: NewResourceReceivedType,
		Content: NewResourceReceivedContent{
			Resource: resource,
		},
	}
}
