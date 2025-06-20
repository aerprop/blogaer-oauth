package config

import (
	"blogaer-oauth/internal/service"
	"blogaer-oauth/internal/utils/enum"
)

type RabbitMQConfig struct {
	AssertQueue map[string]service.AssertQueueParams
	BindQueue   map[string]service.BindQueueParams
	ConsumeMsg  map[string]service.ConsumeMsgParams
}

func LoadRabbitMQConfig() *RabbitMQConfig {
	queueConfig := &RabbitMQConfig{
		AssertQueue: map[string]service.AssertQueueParams{
			enum.QueueName.Google: {
				Queue:     enum.QueueName.Google,
				Durable:   false,
				AutoDel:   false,
				Exclusive: true,
				NoWait:    false,
				Args:      nil,
			},
			enum.QueueName.Github: {
				Queue:     enum.QueueName.Github,
				Durable:   false,
				AutoDel:   false,
				Exclusive: true,
				NoWait:    false,
				Args:      nil,
			},
		},
		BindQueue: map[string]service.BindQueueParams{
			enum.QueueName.Google: {
				Name:     enum.QueueName.Google,
				Key:      enum.RoutingKey.GoogleRK,
				Exchange: enum.RpcExchange.Name,
				NoWait:   false,
				Args:     nil,
			},
			enum.QueueName.Github: {
				Name:     enum.QueueName.Github,
				Key:      enum.RoutingKey.GithubRK,
				Exchange: enum.RpcExchange.Name,
				NoWait:   false,
				Args:     nil,
			},
		},
		ConsumeMsg: map[string]service.ConsumeMsgParams{
			enum.QueueName.Google: {
				Queue:     enum.QueueName.Google,
				Consumer:  enum.Consumer.Google,
				AutoAck:   true,
				Exclusive: false,
				NoLocal:   false,
				NoWait:    false,
				Args:      nil,
			},
			enum.QueueName.Github: {
				Queue:     enum.QueueName.Github,
				Consumer:  enum.Consumer.Github,
				AutoAck:   true,
				Exclusive: false,
				NoLocal:   false,
				NoWait:    false,
				Args:      nil,
			},
		},
	}

	return queueConfig
}
