package main

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

// UserEvent represents the user event structure.
type UserEvent struct {
	ID   uuid.UUID
	Name string
}

func main() {
	var user UserEvent
	// FIXME: all errors was ignored for clarity.
	logrus.Infoln("audit service is starting. Press 'Enter' to exit...")
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		logrus.WithError(err).Fatalln("nats connection error")
	}

	const eventSubject = "user:created"
	_, _ = nc.Subscribe(eventSubject, func(msg *nats.Msg) {
		_ = json.Unmarshal(msg.Data, &user)
		logrus.WithFields(logrus.Fields{
			"user_id":   user.ID,
			"user_name": user.Name,
		}).
			Infoln("process event")
	})
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
