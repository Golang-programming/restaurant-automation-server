package websocket

import (
	"log"

	"github.com/nats-io/nats.go"
)

// NATS Client
var natsConn *nats.Conn

func InitNATS() {
	var err error
	natsConn, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("Error connecting to NATS:", err)
	}
}

func PublishToNATS(subject string, message []byte) {
	if natsConn != nil {
		natsConn.Publish(subject, message)
	}
}

func SubscribeToNATS(subject string, handler func(*nats.Msg)) {
	if natsConn != nil {
		natsConn.Subscribe(subject, handler)
	}
}
