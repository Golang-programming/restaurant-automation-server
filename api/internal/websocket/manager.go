package websocket

/* import "sync"

// Manager is responsible for managing all active WebSocket connections.
var clients map[*Client]bool // Active clients connected to the WebSocket server
var mu sync.Mutex            // Mutex to protect concurrent access to the clients map

// NewManager creates a new WebSocket connection manager.
func init() {
	clients = make(map[*Client]bool)
}

// AddClient registers a new client to the manager.
// func AddClient(client *Client) {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	clients[client] = true
// }

// RemoveClient removes a client from the manager.
// func RemoveClient(client *Client) {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	if _, ok := clients[client]; ok {
// 		delete(clients, client)
// 		close(client.Send)
// 	}
// }

// Broadcast sends a message to all connected clients.
func Broadcast(message *Message) {
	mu.Lock()
	defer mu.Unlock()
	for client := range clients {
		select {
		case client.Send <- message:
		default:
			close(client.Send)
			delete(clients, client)
		}
	}
}
*/
