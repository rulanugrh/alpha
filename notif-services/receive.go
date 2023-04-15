package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/rulanugrh/alpha/notif/config"
	"github.com/rulanugrh/alpha/notif/helpers"
)

type Notif struct {
	OrderID  uint
	Subtotal int
	Paid     bool
	BookID   uint
}

func main() {
	fmt.Printf("START ......")
	Create()

}

func Create() {
	client := config.MngoConnection()
	defer client.Disconnect(context.TODO())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		table = config.MngoColletion("notifserv", client)
	)

	channel, err := helpers.GetConnection()
	if err != nil {
		helpers.FailError(err, "something error connection")
	}
	queue, errs := channel.QueueDeclare("notif-created", false, false, false, false, nil)
	if errs != nil {
		helpers.FailError(errs, "error declare queue")
	}

	message, errs2 := channel.Consume(queue.Name, "", true, false, false, false, nil)
	if errs2 != nil {
		helpers.FailError(errs2, "error consume")
	}

	go func() {
		for d := range message {
			log.Printf("Receive message %s: ", d.Body)
			var payload Notif
			err := json.Unmarshal(d.Body, &payload)
			if err != nil {
				helpers.FailError(err, "cant marshal")
			}

			_, errCreate := table.InsertOne(ctx, payload)
			if errCreate != nil {
				helpers.FailError(errCreate, "cannot insert to db")
			}
		}
	}()
}
