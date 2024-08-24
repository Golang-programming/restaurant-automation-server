package websocket

// Message represents the structure of a WebSocket message.
type Message struct {
	Type    string `json:"type"`    // e.g., "update", "notification"
	Content string `json:"content"` // The actual message content
}
