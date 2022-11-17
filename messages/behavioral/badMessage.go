package behavioral

type BadMessage struct {
	Type    string            `json:"Type"`
	Content BadMessageContent `json:"Content"`
}
type BadMessageContent struct {
	Message string `json:"Message"`
}
