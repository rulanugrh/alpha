package helpers

import (
	"log"

	"github.com/rulanugrh/alpha/order/config"
)

func UserCreated() error {
	queue, errs := Channel.QueueDeclare("user-created", false, false, false, false, nil)
	if errs != nil {
		FailError(errs, "cant declare queue")
	}
	message, err := Channel.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		FailError(err, "cant join this consume")
	}
	var forever chan struct{}

	go func() {
		for d := range message {
			log.Printf("Receive message %s: ", d.Body)
			config.DB.Create(&d.Body)
		}
	}()

	log.Printf("Waiting for messages")
	<-forever
	return nil
}

func UserDeleted() error {
	queue, errs := Channel.QueueDeclare("user-deleted", false, false, false, false, nil)
	if errs != nil {
		FailError(errs, "cant declare queue")
	}
	message, err := Channel.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		FailError(err, "cant join this consume")
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

func BookCreated() error {
	queue, errs := Channel.QueueDeclare("book-created", false, false, false, false, nil)
	if errs != nil {
		FailError(errs, "cant declare queue")
	}
	message, err := Channel.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		FailError(err, "cant join this consume")
	}
	var forever chan struct{}

	go func() {
		for d := range message {
			log.Printf("Receive message %s: ", d.Body)
			config.DB.Create(&d.Body)
		}
	}()

	log.Printf("Waiting for messages")
	<-forever
	return nil
}

func BookDeleted() error {
	queue, errs := Channel.QueueDeclare("book-deleted", false, false, false, false, nil)
	if errs != nil {
		FailError(errs, "cant declare queue")
	}
	message, err := Channel.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		FailError(err, "cant join this consume")
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
