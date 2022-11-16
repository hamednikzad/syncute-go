package behavioral

const CommandName = "bad_message"

type BadMessage struct {
	Type    string            `json:"Type"`
	Content BadMessageContent `json:"Content"`
}
type BadMessageContent struct {
	Message string `json:"Message"`
	//Content SectionContent
}
