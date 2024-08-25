package customer

/*
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ConnectToWS(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Fatalln("Error while upgrading connection:", err)
	}

	WSHanlders(conn)

	defer conn.Close()
}

func WSHanlders(conn *websocket.Conn) {
	for {

		messageType, message, err := conn.ReadMessage()

		if err != nil {
			log.Fatalln("Error reading message:", err)
			break
		}

		// Unmarshal the message into WSMessage struct
		var msg structs.WSMessage
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}

		switch msg.Type {
		case "onMessage":
			handleOnMessage(conn, msg.Payload)
		default:
			log.Println("Unknown event type:", msg.Type)
		}

	}
} */

/* messageType, msg, err := conn.ReadMessage()
if err != nil {
	log.Println("Error while reading message:", err)
	// break
}

switch messageType {
case websocket.TextMessage:
	log.Printf("Received text message: %s", msg)
default:
	log.Printf("Received unknown message type: %d", messageType)
}

// Echo the received message back to the client
err = conn.WriteMessage(messageType, msg)
if err != nil {
	log.Println("Error while writing message:", err)
} */

/*
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
func init() {

	&BillWebSocket{
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
	// ws.hub.Broadcast(message)
	// ws.natsClient.Publish("bills", message) // Optional for scaling
}
*/
