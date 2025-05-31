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
			enum.QueueName.GoogleQ: {
				Queue:     enum.QueueName.GoogleQ,
				Durable:   false,
				AutoDel:   false,
				Exclusive: true,
				NoWait:    false,
				Args:      nil,
			},
			enum.QueueName.GithubQ: {
				Queue:     enum.QueueName.GithubQ,
				Durable:   false,
				AutoDel:   false,
				Exclusive: true,
				NoWait:    false,
				Args:      nil,
			},
		},
		BindQueue: map[string]service.BindQueueParams{
			enum.QueueName.GoogleQ: {
				Name:     enum.QueueName.GoogleQ,
				Key:      enum.RoutingKey.GoogleRK,
				Exchange: enum.RpcExchange.Name,
				NoWait:   false,
				Args:     nil,
			},
			enum.QueueName.GithubQ: {
				Name:     enum.QueueName.GithubQ,
				Key:      enum.RoutingKey.GithubRK,
				Exchange: enum.RpcExchange.Name,
				NoWait:   false,
				Args:     nil,
			},
		},
		ConsumeMsg: map[string]service.ConsumeMsgParams{
			enum.QueueName.GoogleQ: {
				Queue:     enum.QueueName.GoogleQ,
				Consumer:  enum.Channel.Consumer,
				AutoAck:   true,
				Exclusive: false,
				NoLocal:   false,
				NoWait:    false,
				Args:      nil,
			},
			enum.QueueName.GithubQ: {
				Queue:     enum.QueueName.GithubQ,
				Consumer:  enum.Channel.Consumer,
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
