package structs

type WSMessage struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}
