package global

import (
	"kim/internal/logic/config"
	"kim/internal/pkg/log"
	"kim/internal/pkg/mq"
	"kim/internal/pkg/mysql"
	"kim/internal/pkg/redis"
)

const (
	Token     = "token"
	UserIdKey = "user_id"
)

const (
	NotificationExchangeName = "notification_exchange"
	NotificationQueueName    = "notification_queue"
	NotificationRoutingKey   = "notification_routing_key"
)

var (
	Conf *config.Config

	Database *mysql.Mysql
	Redis    *redis.Redis
	Logger   *log.Logger
	RabbitMQ *mq.RabbitMQClient
	Channel  *mq.RabbitMQChannel
)

func InitMQ() {
	var err error
	Channel, err = RabbitMQ.GetChannel()
	if err != nil {
		panic(err)
	}
	Channel.DeclareExchange(NotificationExchangeName, "direct", mq.WithDurable(true))
	Channel.DeclareQueue(NotificationQueueName, mq.WithDurable(true))
	Channel.BindingQueue(NotificationExchangeName, NotificationQueueName, NotificationRoutingKey, mq.WithDurable(true))
}
