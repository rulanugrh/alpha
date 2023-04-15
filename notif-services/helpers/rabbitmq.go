package helpers

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rulanugrh/alpha/notif/config"
)

func GetConnection() (*amqp.Channel, error) {
	conf := config.GetConfig()

	logon := fmt.Sprintf("amqp://%s:%s@%s:%s",
		conf.RabbitMQ.User,
		conf.RabbitMQ.Pass,
		conf.RabbitMQ.Host,
		conf.RabbitMQ.Port,
	)

	conn, err := amqp.Dial(logon)
	if err != nil {
		FailError(err, "something error for login")
	}
	defer conn.Close()

	ch, errs := conn.Channel()
	if errs != nil {
		FailError(errs, "something error for channel")
	}
	defer ch.Close()

	// Channel = ch
	return ch, nil
}
