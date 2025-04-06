package helpers

import (
	"log"

	"github.com/nats-io/nats.go"
)

var NatsConn *nats.Conn

func InitNats() {
	var err error
	NatsConn, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Erreur de connexion à NATS : %v", err)
	}
	log.Println("Connecté à NATS !")
}
