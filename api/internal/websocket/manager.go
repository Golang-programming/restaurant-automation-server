package websocket

import "sync"

// Manager is responsible for managing all active WebSocket connections.
type Manager struct {
	clients map[*Client]bool // Active clients connected to the WebSocket server
	mu      sync.Mutex       // Mutex to protect concurrent access to the clients map
}

// NewManager creates a new WebSocket connection manager.
func NewManager() *Manager {
	return &Manager{
		clients: make(map[*Client]bool),
	}
}

// AddClient registers a new client to the manager.
func (m *Manager) AddClient(client *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.clients[client] = true
}

// RemoveClient removes a client from the manager.
func (m *Manager) RemoveClient(client *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.clients[client]; ok {
		delete(m.clients, client)
		close(client.Send)
	}
}

// Broadcast sends a message to all connected clients.
func (m *Manager) Broadcast(message *Message) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for client := range m.clients {
		select {
		case client.Send <- message:
		default:
			close(client.Send)
			delete(m.clients, client)
		}
	}
}
