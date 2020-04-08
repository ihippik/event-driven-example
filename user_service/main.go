package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/nats-io/go-nats"
	"github.com/sirupsen/logrus"
)

// publisher
type publisher interface {
	Publish(string, []byte) error
}

// Service represent еру service structure.
type Service struct {
	publisher publisher
}

// User represents the user structure.
type User struct {
	ID   uuid.UUID
	Name string
}

// newService create new Service instance.
func newService(pub publisher) *Service {
	return &Service{publisher: pub}
}

func main() {
	logrus.Infoln("user service is starting on port 8080..")
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		logrus.WithError(err).Fatalln("nats connection error")
	}
	srv := newService(nc)
	http.HandleFunc("/users", srv.AddUser)
	if err = http.ListenAndServe(":8080", nil); err != nil {
		logrus.WithError(err).Fatalln("listenAndServe error")
	}
}

// AddUser add new user to db.
func (s Service) AddUser(w http.ResponseWriter, req *http.Request) {
	var user User
	// FIXME: all errors was ignored for clarity.
	// TODO: allow only POST method.
	// decode user request body into our structure.
	decoder := json.NewDecoder(req.Body)
	_ = decoder.Decode(&user)

	// TODO: validate user data.
	user.ID = uuid.New()

	// TODO:  save model into db.
	data, _ := json.Marshal(user)
	// asynchronously notify services that need to know about creating a user.
	const eventSubject = "user:created"
	_ = s.publisher.Publish(eventSubject, data)

	// creating a new entity width 201 http code (REST).
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(data)
}
