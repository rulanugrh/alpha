package helpers

import (
	"encoding/json"
	"log"

	"github.com/rulanugrh/alpha/order/entities/domain"
	"github.com/rulanugrh/alpha/order/entities/web"
	"github.com/rulanugrh/alpha/order/repository/book"
	"github.com/rulanugrh/alpha/order/repository/user"
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
			var payload domain.User
			err := json.Unmarshal(d.Body, &payload)
			if err != nil {
				log.Printf("Something error unmarshal %s:", err)
			}

			user.NewUserRepository().Create(payload)

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
			var payload domain.Book
			err := json.Unmarshal(d.Body, &payload)
			if err != nil {
				log.Printf("Something error unmarshal %s:", err)
			}

			book.NewBookRepository().Create(payload)

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
			var payload web.BookDelete
			log.Printf("Receive message %s: ", d.Body)

			err := json.Unmarshal(d.Body, &payload)
			if err != nil {
				log.Printf("error cant unmarshal")
			}

			book.NewBookRepository().Delete(payload.Id)
		}
	}()

	log.Printf("Waiting for messages")
	<-forever
	return nil
}
