package main

import (
	"log"

	"middleware/example/internal/consumer"
	"middleware/example/internal/helpers"
)

func main() {
	helpers.InitNats()

	if err := consumer.StartAlerterConsumer(); err != nil {
		log.Fatalf("Erreur lors du démarrage du consumer alerter: %v", err)
	}
}
