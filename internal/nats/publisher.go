package nats

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func StartNATSPublisher(nc *nats.Conn) {
	// Send a message
	subject := "updates"
	message := "Hello NATS!"

	if err := nc.Publish(subject, []byte(message)); err != nil {
		log.Fatalf("❌ Failed to publish message: %v", err)
	}
	fmt.Println("✅ Message sent:", subject, "->", message)
}
