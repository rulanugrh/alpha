package helpers

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rulanugrh/alpha/book/config"
)

var Channel *amqp.Channel

func GetConnection() *amqp.Channel {
	conf := config.GetConfig()

	logon := fmt.Sprintf("amqp://%s:%s@%s:%s",
		conf.RabbitMQ.User,
		conf.RabbitMQ.Pass,
		conf.RabbitMQ.Host,
		conf.RabbitMQ.Port,
	)

	conn, err := amqp.Dial(logon)
	if err != nil {
		FailOnError(err, "something error for login")
	}
	defer conn.Close()

	ch, errs := conn.Channel()
	if errs != nil {
		FailOnError(errs, "something error for channel")
	}
	defer ch.Close()

	Channel = ch
	return ch

}
