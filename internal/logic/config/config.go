package config

import (
	"kim/internal/pkg/log"
	"kim/internal/pkg/mq"
	"kim/internal/pkg/mysql"
	"kim/internal/pkg/redis"
	"kim/internal/pkg/server"
)

type Config struct {
	Mysql      mysql.Config            `json:"mysql" yaml:"mysql"`
	Redis      redis.Config            `json:"redis" yaml:"redis"`
	Log        log.Config              `json:"log" yaml:"log"`
	HttpServer server.HttpServerConfig `json:"http_server" yaml:"http_server"`
	RabbitMQ   mq.Config               `json:"rabbitmq" yaml:"rabbitmq"`
}
