package websocket

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

// NatsClient is responsible for publishing and subscribing to WebSocket messages.
var nc *nats.Conn

// NewNatsClient initializes a new NATS client.
func NewNatsClient(url string) {
	nConnection, err := nats.Connect(url)
	if err != nil {
		log.Fatalln(err)
	}

	nc = nConnection
}

// Subscribe subscribes to a given subject and handles messages with the provided handler.
func Subscribe(subject string, handler func(*nats.Msg)) error {
	_, err := nc.Subscribe(subject, handler)
	return err
}

// Publish publishes a message to a given subject.
func Publish(subject string, message *Message) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return nc.Publish(subject, data)
}

func Close() {
	nc.Close()
}
