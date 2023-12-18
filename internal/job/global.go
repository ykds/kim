package job

import (
	"kim/internal/pkg/log"
	"kim/internal/pkg/mq"
)

var (
	Conf *Config

	RabbitMQ *mq.RabbitMQClient
	Logger   *log.Logger
)
