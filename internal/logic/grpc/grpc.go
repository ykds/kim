package grpc

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"kim/internal/logic/config"
	"kim/internal/logic/global"
	"kim/internal/logic/service"
	"kim/proto/logic"
	"net"
	"time"
)

func grpcLogger() logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, level logging.Level, msg string, fields ...any) {
		switch level {
		case logging.LevelDebug:
			global.Logger.Debugf(fmt.Sprintf("%s, field: %v", msg, fields))
		case logging.LevelInfo:
			global.Logger.Infof(fmt.Sprintf("%s, field: %v", msg, fields))
		case logging.LevelWarn:
			global.Logger.Warnf(fmt.Sprintf("%s, field: %v", msg, fields))
		case logging.LevelError:
			global.Logger.Errorf(fmt.Sprintf("%s, field: %v", msg, fields))
		}
	})
}

func NewGrpcServer(cfg config.ServerConfig, service *service.Service) *grpc.Server {
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor()),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 3 * time.Minute,
			Time:              10 * time.Second,
			Timeout:           3 * time.Second,
		}),
	)
	logic.RegisterLogicServiceServer(srv, &LogicGrpcServer{srv: service})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	if err != nil {
		panic(err)
	}
	go func() {
		err := srv.Serve(listen)
		if err != nil {
			panic(err)
		}
	}()
	return srv
}

var _ logic.LogicServiceServer = &LogicGrpcServer{}

type LogicGrpcServer struct {
	logic.UnimplementedLogicServiceServer
	srv *service.Service
}

func (l LogicGrpcServer) Auth(ctx context.Context, req *logic.AuthReq) (*logic.AuthResp, error) {
	userId, err := l.srv.UserService.Auth(req.GetToken())
	if err != nil {
		return nil, err
	}
	return &logic.AuthResp{
		UserId: int32(userId),
	}, nil
}

func (l LogicGrpcServer) HeartBeat(ctx context.Context, req *logic.HeartBeatReq) (*logic.HeartBeatResp, error) {
	err := l.srv.UserService.HeartBeat(req.GetServerId(), uint(req.GetUserId()))
	if err != nil {
		return nil, err
	}
	return &logic.HeartBeatResp{}, nil
}
