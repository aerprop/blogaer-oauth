package main

import (
	"blogaer-oauth/internal/messaging/connection"
	"blogaer-oauth/internal/messaging/rpc"
	"blogaer-oauth/internal/service"
	"blogaer-oauth/internal/utils/config"
	"blogaer-oauth/internal/utils/enum"
	"blogaer-oauth/internal/utils/helper"
	"blogaer-oauth/internal/utils/types"
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := connection.RabbitMQConn()
	helper.OnError(err, "Failed to connect to RabbitMQ!")
	defer conn.Close()
	fmt.Println("Connected to RabbitMQ ✔✔✔")

	consumerChan, err := service.InitChannel(conn)
	helper.OnError(err, "Failed to create consumer channel!")
	defer consumerChan.Close()

	publisherChan, err := service.InitChannel(conn)
	helper.OnError(err, "Failed to create publisher channel!")
	defer publisherChan.Close()

	err = consumerChan.Qos(1, 0, false)
	helper.OnError(err, "Failed to declare consumer channel Qos!")
	defer consumerChan.Close()

	rabbitMQConf := config.LoadRabbitMQConfig()
	queues := make(map[string]<-chan amqp.Delivery)

	for queueName, conf := range rabbitMQConf.AssertQueue {
		_, err := service.AssertQueue(consumerChan, &conf)
		helper.OnError(err, "Failed to assert queue!")

		queueBindConf := rabbitMQConf.BindQueue[queueName]
		err = service.BindQueue(consumerChan, &queueBindConf)
		helper.OnError(err, "Failed to bind queue!")

		consumeRpcMsgConf := rabbitMQConf.ConsumeMsg[queueName]
		msgs, err := service.ConsumeMsg(consumerChan, &consumeRpcMsgConf)
		helper.OnError(err, "Failed to register a consumer!")

		queues[queueName] = msgs
	}

	forever := make(chan struct{})

	go func() {
		ctx := context.Background()
		for queueName, messages := range queues {
			go func(name string, msgs <-chan amqp.Delivery) {
				for delivery := range msgs {
					var message types.Message
					err := json.Unmarshal(delivery.Body, &message)
					if err != nil {
						helper.OnError(err, "Failed to parse json data!")
					}

					switch delivery.RoutingKey {
					case enum.RoutingKey.GoogleRK:
						rpc.GoogleOauth(ctx, publisherChan, delivery, message.Code)
					case enum.RoutingKey.GithubRK:
						rpc.GithubOauth(ctx, publisherChan, delivery, message.Code)
					}
				}
			}(queueName, messages)
		}
	}()

	<-forever
}
