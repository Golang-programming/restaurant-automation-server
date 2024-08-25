package websocket

import "sync"

// "net/http"
// "sync"

// "github.com/gin-gonic/gin"
// "github.com/gorilla/websocket"

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	mu         sync.Mutex
}

// var hubInstance *Hub
// var once sync.Once

// func GetHub() *Hub {
// 	// once.Do(func() {
// 	// 	hubInstance = &Hub{
// 	// 		Broadcast:  make(chan []byte),
// 	// 		Register:   make(chan *Client),
// 	// 		Unregister: make(chan *Client),
// 	// 		Clients:    make(map[*Client]bool),
// 	// 	}
// 	// 	go hubInstance.run()
// 	// })
// 	// return hubInstance
// }

// func (h *Hub) run() {
// 	// for {
// 	// 	select {
// 	// 	case client := <-h.Register:
// 	// 		h.mu.Lock()
// 	// 		h.Clients[client] = true
// 	// 		h.mu.Unlock()
// 	// 	case client := <-h.Unregister:
// 	// 		h.mu.Lock()
// 	// 		if _, ok := h.Clients[client]; ok {
// 	// 			delete(h.Clients, client)
// 	// 			close(client.Send)
// 	// 		}
// 	// 		h.mu.Unlock()
// 	// 	case message := <-h.Broadcast:
// 	// 		h.mu.Lock()
// 	// 		for client := range h.Clients {
// 	// 			select {
// 	// 			case client.Send <- message:
// 	// 			default:
// 	// 				close(client.Send)
// 	// 				delete(h.Clients, client)
// 	// 			}
// 	// 		}
// 	// 		h.mu.Unlock()
// 	// 	}
// 	// }
// }

// func (h *Hub) ServeWS(c *gin.Context) {
// 	// upgrader := websocket.Upgrader{
// 	// 	CheckOrigin: func(r *http.Request) bool {
// 	// 		return true
// 	// 	},
// 	// }
// 	// conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	// if err != nil {
// 	// 	return
// 	// }
// 	// client := &Client{Conn: conn, Send: make(chan []byte, 256)}
// 	// h.Register <- client

// 	// go client.writePump()
// 	// go client.readPump()
// }
