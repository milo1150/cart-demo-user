package nats

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

// Connects to NATS_URL: nats://nats:4222 (Docker service name).
func ConnectNATS() *nats.Conn {
	NATSUrl := os.Getenv("NATS_URL")

	nc, err := nats.Connect(NATSUrl)
	if err != nil {
		log.Fatalf("Failed to connect NATS server")
	}

	return nc
}
