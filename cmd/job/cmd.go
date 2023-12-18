package main

import (
	"gopkg.in/yaml.v3"
	"kim/internal/job"
	"kim/internal/pkg/etcd"
	"kim/internal/pkg/log"
	"kim/internal/pkg/mq"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initConfig()
	job.Logger = log.InitLogger(job.Conf.Log)
	job.RabbitMQ = mq.NewRabbitMQClient(job.Conf.RabbitMQ)

	srv := job.NewServer(job.Conf.Server)
	etcdManager := etcd.NewEtcd(job.Conf.Etcd)
	go job.Watch(etcdManager, srv)
	go srv.Consume()

	job.Logger.Infof("job server started")
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case <-c:
			job.Logger.Infof("job server stop")
			return
		}
	}
}

func initConfig() {
	file, err := os.ReadFile("./cmd/job/config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, &job.Conf)
	if err != nil {
		panic(err)
	}
}
