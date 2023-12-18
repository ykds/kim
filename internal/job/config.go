package job

import (
	"kim/internal/pkg/etcd"
	"kim/internal/pkg/log"
	"kim/internal/pkg/mq"
)

type ServerConfig struct {
	Id    int32  `json:"id" yaml:"id"`
	Queue string `json:"queue" yaml:"queue"`
}

type Config struct {
	Server   ServerConfig `json:"server" yaml:"server"`
	RabbitMQ mq.Config    `json:"rabbitmq" yaml:"rabbitmq"`
	Log      log.Config   `json:"log" yaml:"log"`
	Etcd     etcd.Config  `json:"etcd" yaml:"etcd"`
}
