package grpc

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"kim/internal/comet/config"
	"kim/internal/comet/server"
	"kim/proto/comet"
	"net"
	"time"
)

func NewGrpcServer(cfg config.GrpcConfig, server *server.Server) *grpc.Server {
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor()),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 3 * time.Minute,
			Time:              10 * time.Second,
			Timeout:           3 * time.Second,
		}),
	)
	comet.RegisterCometServer(srv, &Server{srv: server})
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

var _ comet.CometServer = &Server{}

type Server struct {
	comet.UnimplementedCometServer
	srv *server.Server
}

func (g Server) PushMessage(ctx context.Context, req *comet.PushMessageReq) (*comet.PushMessageResp, error) {
	uid := uint(req.Message.UserId)
	ch, ok := g.srv.GetBucket(uid).Get(uid)
	if ok {
		ch.Put(req.Message)
	}
	return &comet.PushMessageResp{}, nil
}
