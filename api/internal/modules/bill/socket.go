package bill

import (
	"encoding/json"
	"log"

	"github.co/golang-programming/restaurant/api/internal/websocket"
)

// BillWebSocket handles WebSocket logic for the "bill" module.
type BillWebSocket struct {
	hub        *websocket.Hub        // Central WebSocket hub
	natsClient *websocket.NatsClient // NATS client for scalability
}

// NewBillWebSocket initializes a BillWebSocket handler.
func NewBillWebSocket(hub *websocket.Hub, natsClient *websocket.NatsClient) *BillWebSocket {
	return &BillWebSocket{
		hub:        hub,
		natsClient: natsClient,
	}
}

// HandleMessage processes incoming WebSocket messages for bills.
func (ws *BillWebSocket) HandleMessage(client *websocket.Client, message []byte) {
	var msg websocket.Message
	if err := json.Unmarshal(message, &msg); err != nil {
		log.Printf("Error parsing WebSocket message: %v", err)
		return
	}

	switch msg.Event {
	case "payBill":
		ws.handlePayBill(msg.Payload)
	}
}

// handlePayBill processes the "payBill" event.
func (ws *BillWebSocket) handlePayBill(payload interface{}) {
	message := websocket.NewMessage("billPaid", payload)
	ws.hub.Broadcast(message)
	ws.natsClient.Publish("bills", message) // Optional for scaling
}
