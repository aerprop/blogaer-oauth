package service

import (
	"blogaer-oauth/internal/utils/enum"
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AssertExchParams struct {
	Exchange string
	Kind     string
	Durable  bool
	AutoDel  bool
	Internal bool
	NoWait   bool
	Args     amqp.Table
}

func InitChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	config := AssertExchParams{
		Exchange: enum.RpcExchange.Name,
		Kind:     enum.RpcExchange.Type,
		Durable:  false,
		AutoDel:  false,
		Internal: false,
		NoWait:   false,
		Args:     nil,
	}
	AssertExchange(channel, &config)

	return channel, nil
}

func AssertExchange(
	channel *amqp.Channel,
	config *AssertExchParams,
) (
	*amqp.Channel,
	error,
) {
	err := channel.ExchangeDeclare(
		config.Exchange,
		config.Kind,
		config.Durable,
		config.AutoDel,
		config.Internal,
		config.NoWait,
		config.Args,
	)
	if err != nil {
		return nil, err
	}

	return channel, nil
}

type AssertQueueParams struct {
	Queue    string
	Durable  bool
	AutoDel  bool
	Internal bool
	NoWait   bool
	Args     amqp.Table
}

func AssertQueue(
	channel *amqp.Channel,
	config *AssertQueueParams,
) (
	*amqp.Queue,
	error,
) {
	queue, err := channel.QueueDeclare(
		config.Queue,
		config.Durable,
		config.AutoDel,
		config.Internal,
		config.NoWait,
		config.Args,
	)
	if err != nil {
		return nil, err
	}

	return &queue, nil
}

type BindQueueParams struct {
	Name     string
	Key      string
	Exchange string
	NoWait   bool
	Args     amqp.Table
}

func BindQueue(
	channel *amqp.Channel,
	config *BindQueueParams,
) error {
	err := channel.QueueBind(
		config.Name,
		config.Key,
		config.Exchange,
		config.NoWait,
		config.Args,
	)
	if err != nil {
		return err
	}

	return nil
}

type ConsumeMsgParams struct {
	Queue     string
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
}

func ConsumeMsg(
	channel *amqp.Channel,
	config *ConsumeMsgParams,
) (
	<-chan amqp.Delivery,
	error,
) {
	msg, err := channel.Consume(
		config.Queue,
		config.Consumer,
		config.AutoAck,
		config.Exclusive,
		config.NoLocal,
		config.NoWait,
		config.Args,
	)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

type PublishMsgParams struct {
	Ctx       context.Context
	Exchange  string
	Key       string
	Mandatory bool
	Immediate bool
	Msg       amqp.Publishing
}

func PublishMsg(
	channel *amqp.Channel,
	config *PublishMsgParams,
) error {
	err := channel.PublishWithContext(
		config.Ctx,
		config.Exchange,
		config.Key,
		config.Mandatory,
		config.Immediate,
		config.Msg,
	)
	if err != nil {
		return err
	}

	return nil
}
