package connection

import amqp "github.com/rabbitmq/amqp091-go"

func RabbitMQConn() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://anekra:1234@localhost:5672")
	if err != nil {
		return nil, err
	}

	return conn, nil
}
