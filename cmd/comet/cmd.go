package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"kim/internal/comet/global"
	"kim/internal/comet/grpc"
	"kim/internal/comet/server"
	"kim/internal/pkg/etcd"
	"kim/internal/pkg/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initConfig()
	global.Logger = log.InitLogger(global.Conf.Log)

	srv := server.NewServer(global.Conf.Server)
	grpcServer := grpc.NewGrpcServer(global.Conf.GrpcServer, srv)

	etcdManager := etcd.NewEtcd(global.Conf.Etcd)
	key := fmt.Sprintf("/comet_%d", srv.GetId())
	err := etcdManager.Register(key, fmt.Sprintf("%s:%s", global.Conf.GrpcServer.Host, global.Conf.GrpcServer.Port), map[string]interface{}{"server_id": srv.GetId()})
	if err != nil {
		panic(err)
	}

	global.Logger.Infof("comet server started")
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case <-c:
			etcdManager.UnRegister(key)
			grpcServer.GracefulStop()
			global.Logger.Infof("comet server stop")
			return
		}
	}
}

func initConfig() {
	file, err := os.ReadFile("./cmd/comet/config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, &global.Conf)
	if err != nil {
		panic(err)
	}
}
