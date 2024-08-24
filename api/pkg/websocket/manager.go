package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

// Manager manages all WebSocket connections.
type Manager struct {
	connections map[*websocket.Conn]bool
	mutex       sync.Mutex
}

// NewManager creates a new WebSocket Manager.
func NewManager() *Manager {
	return &Manager{
		connections: make(map[*websocket.Conn]bool),
	}
}

// AddConnection adds a new WebSocket connection to the manager.
func (m *Manager) AddConnection(conn *websocket.Conn) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.connections[conn] = true
}

// RemoveConnection removes a WebSocket connection from the manager.
func (m *Manager) RemoveConnection(conn *websocket.Conn) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if _, ok := m.connections[conn]; ok {
		delete(m.connections, conn)
		conn.Close()
	}
}

// Broadcast sends a message to all connected WebSocket clients.
func (m *Manager) Broadcast(message []byte) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for conn := range m.connections {
		conn.WriteMessage(websocket.TextMessage, message)
	}
}
