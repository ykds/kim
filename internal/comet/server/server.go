package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"kim/internal/comet/config"
	"kim/internal/comet/global"
	"kim/internal/pkg/response"
	"kim/proto/logic"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{}

type Server struct {
	cfg        config.ServerConfig
	id         int32
	bucketsNum int
	buckets    []*Bucket

	logicClient logic.LogicServiceClient
}

func NewServer(cfg config.ServerConfig) *Server {
	srv := &Server{
		cfg:        cfg,
		id:         cfg.ID,
		bucketsNum: cfg.BucketNum,
		buckets:    make([]*Bucket, cfg.BucketNum),
	}
	for i := 0; i < srv.bucketsNum; i++ {
		srv.buckets[i] = NewBucket()
	}
	srv.logicClient = dialLogicServer(cfg.LogicHost, cfg.LogicPort)

	engine := gin.Default()
	engine.GET("/ws", srv.handleWebsocket)
	httpServer := http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Handler: engine,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			return
		}
	}()
	return srv
}

func dialLogicServer(host, port string) logic.LogicServiceClient {
	conn, err := grpc.DialContext(context.Background(), fmt.Sprintf("%s:%s", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	return logic.NewLogicServiceClient(conn)
}

func (s *Server) handleWebsocket(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	auth, err := s.logicClient.Auth(context.Background(), &logic.AuthReq{
		Token: token,
	})
	if err != nil {
		response.HandleResponse(ctx, errors.New("连接失败"), nil)
		return
	}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	userId := uint(auth.GetUserId())
	ch := NewChannel(conn, s.cfg.HeartBeatInterval)
	b := s.GetBucket(userId)
	b.Put(userId, ch)
	ctx.Done()

	go func() {
		for {
			select {
			case msg := <-ch.Signal():
				body, err := json.Marshal(msg)
				if err != nil {
					global.Logger.Errorf("序列化消息失败, msg: %v, err: %v", msg, err)
					continue
				}
				err = ch.WriteMessage(body)
				if err != nil {
					global.Logger.Errorf("发送消息失败, err: %v", err)
					continue
				}
			case <-ch.Done():
				return
			}
		}
	}()
	go func() {
		for {
			msgType, _, err := ch.conn.ReadMessage()
			if err != nil {
				global.Logger.Errorf("读取消息失败, err: %v", err)
				ch.Close()
				return
			}
			switch msgType {
			case websocket.PingMessage:
				i := 0
				for ; i < 3; i++ {
					ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
					_, err = s.logicClient.HeartBeat(ctx, &logic.HeartBeatReq{ServerId: s.id, UserId: int32(userId)})
					if err != nil {
						cancel()
						continue
					}
					cancel()
				}
				global.Logger.Errorf("发送心跳失败, err: %v", err)
				if i == 4 {
					_ = ch.Close()
					return
				}
			}
		}
	}()
	go func() {
		_, err = s.logicClient.HeartBeat(ctx, &logic.HeartBeatReq{ServerId: s.id, UserId: int32(userId)})
		if err != nil {
			global.Logger.Errorf("发送心跳失败, err: %v", err)
			ch.Close()
			return
		}
		for {
			select {
			case <-ch.HeartBeat():
				global.Logger.Errorf("uid: %d, 心跳超时，连接断开", userId)
				ch.Close()
				s.logicClient.DisConnect(context.Background(), &logic.DisConnectReq{UserId: int32(userId)})
				return
			case <-ch.Done():
				return
			}
		}
	}()
}

func (s *Server) GetBucket(userId uint) *Bucket {
	return s.buckets[int(userId)%s.bucketsNum]
}

func (s *Server) GetId() int32 {
	return s.id
}

func (s *Server) Addr() string {
	return fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port)
}
