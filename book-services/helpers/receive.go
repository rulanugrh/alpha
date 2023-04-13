package helpers

import (
	"encoding/json"
	"log"

	"github.com/rulanugrh/alpha/book/entities/domain"
	books "github.com/rulanugrh/alpha/book/repository/book"
)

func UserCreated() error {
	queue, errs := Channel.QueueDeclare("user-created", false, false, false, false, nil)
	if errs != nil {
		FailOnError(errs, "cant declare queue")
	}
	message, err := Channel.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		FailOnError(err, "cant join this consume")
	}
	var forever chan struct{}

	go func() {
		for d := range message {
			log.Printf("Receive message %s: ", d.Body)
			var payload domain.Seller
			err := json.Unmarshal(d.Body, &payload)
			if err != nil {
				log.Printf("Something error unmarshal %s:", err)
			}

			books.NewBookRepository().Seller(payload)

		}
	}()

	log.Printf("Waiting for messages")
	<-forever
	return nil
}

func UserDeleted() error {
	queue, errs := Channel.QueueDeclare("user-deleted", false, false, false, false, nil)
	if errs != nil {
		FailOnError(errs, "cant declare queue")
	}
	message, err := Channel.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		FailOnError(err, "cant join this consume")
	}
	var forever chan struct{}

	go func() {
		for d := range message {
			log.Printf("Receive message %s: ", d.Body)
		}
	}()

	log.Printf("Waiting for messages")
	<-forever
	return nil
}
