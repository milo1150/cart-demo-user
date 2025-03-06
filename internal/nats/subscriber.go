package nats

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func StartNATSListener(nc *nats.Conn) {
	subject := "updates"
	_, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		fmt.Println("📩 Received NATS Message:", string(msg.Data))
	})
	if err != nil {
		log.Fatalf("❌ Failed to subscribe to NATS: %v", err)
	}
	fmt.Println("👂 Listening for messages on subject:", subject)
}
