package helpers

import "log"

func UserCreated() error {
	queue, errs := channel.QueueDeclare("user-created", false, false, false, false, nil)
	if errs != nil {
		FailError(errs, "cant declare queue")
	}
	message, err := channel.Consume(queue.Name, "", true, false, false, false, nil)
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

func UserDeleted() error {
	queue, errs := channel.QueueDeclare("user-deleted", false, false, false, false, nil)
	if errs != nil {
		FailError(errs, "cant declare queue")
	}
	message, err := channel.Consume(queue.Name, "", true, false, false, false, nil)
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
	queue, errs := channel.QueueDeclare("book-created", false, false, false, false, nil)
	if errs != nil {
		FailError(errs, "cant declare queue")
	}
	message, err := channel.Consume(queue.Name, "", true, false, false, false, nil)
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

func BookDeleted() error {
	queue, errs := channel.QueueDeclare("book-deleted", false, false, false, false, nil)
	if errs != nil {
		FailError(errs, "cant declare queue")
	}
	message, err := channel.Consume(queue.Name, "", true, false, false, false, nil)
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
