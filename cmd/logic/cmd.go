package main

import (
	"gopkg.in/yaml.v3"
	"kim/internal/logic/api"
	"kim/internal/logic/global"
	"kim/internal/logic/grpc"
	"kim/internal/pkg/log"
	"kim/internal/pkg/mq"
	"kim/internal/pkg/mysql"
	"kim/internal/pkg/redis"
	"kim/internal/pkg/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initConfig()
	global.Database = mysql.InitMysql(global.Conf.Mysql)
	global.Redis = redis.NewRedis(global.Conf.Redis)
	global.Logger = log.InitLogger(global.Conf.Log)
	global.RabbitMQ = mq.NewRabbitMQClient(global.Conf.RabbitMQ)
	global.InitMQ()

	httpServer := server.NewHttpServer(global.Conf.HttpServer, global.Logger)
	httpServer.RegisterRouter(api.InitRouter)
	go func() {
		httpServer.Run()
	}()
	grpcServer := grpc.NewGrpcServer(global.Conf.GrpcServer, api.GetService())

	global.Logger.Infof("http server started")
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case <-c:
			httpServer.Stop()
			grpcServer.GracefulStop()
			global.Logger.Infof("http server stop")
			return
		}
	}
}

func initConfig() {
	file, err := os.ReadFile("./cmd/logic/config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, &global.Conf)
	if err != nil {
		panic(err)
	}
}
