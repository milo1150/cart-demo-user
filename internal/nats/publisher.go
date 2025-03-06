package nats

import (
	"strconv"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

func PublishUserCreated(nc *nats.Conn, log *zap.Logger) {
	subject := "user.created"
	userId := strconv.Itoa(22) // TODO: use real user id
	data := []byte(userId)

	if err := nc.Publish(subject, data); err != nil {
		log.Error("Failed to publish",
			zap.String("subject", subject),
			zap.Error(err),
		)
		return
	}
}
