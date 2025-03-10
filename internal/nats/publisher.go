package nats

import (
	cartpkg "github.com/milo1150/cart-demo-pkg/pkg"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

func PublishUserCreated(nc *nats.Conn, log *zap.Logger, userId uint) {
	subject := "user.created"
	data := cartpkg.UintToBytes(userId)

	if err := nc.Publish(subject, data); err != nil {
		log.Error("Failed to publish",
			zap.String("subject", subject),
			zap.Error(err),
		)
		return
	}
}
