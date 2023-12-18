package job

import (
	"context"
	"encoding/json"
	"fmt"
	"kim/internal/protocol"
	"kim/proto/comet"
	"sync"
)

type Server struct {
	Id     int32
	m      sync.RWMutex
	comets map[int32]comet.CometClient

	queue string
}

func NewServer(cfg ServerConfig) *Server {
	return &Server{
		Id:     cfg.Id,
		queue:  cfg.Queue,
		comets: make(map[int32]comet.CometClient),
	}
}

func (s *Server) GetComet(serverId int32) (comet.CometClient, bool) {
	s.m.RLock()
	defer s.m.RUnlock()
	client, ok := s.comets[serverId]
	return client, ok
}

func (s *Server) AddComet(serverId int32, client comet.CometClient) {
	s.m.Lock()
	defer s.m.Unlock()
	s.comets[serverId] = client
}

func (s *Server) DelComet(serverId int32) {
	s.m.Lock()
	defer s.m.Unlock()
	delete(s.comets, serverId)
}

func (s *Server) Consume() {
	channel, err := RabbitMQ.Channel()
	if err != nil {
		return
	}
	channel.Qos(10, 0, false)

	consume, err := channel.Consume(s.queue, fmt.Sprintf("job_%d", s.Id), false, false, false, false, nil)
	if err != nil {
		return
	}
	for {
		select {
		case msg := <-consume:
			//message := &logic.Message{}
			//err := json.Unmarshal(msg.Body, message)
			//if err != nil {
			//	Logger.Errorf("消息格式错误, body, %s, err: %v", string(msg.Body), err)
			//	msg.Reject(false)
			//	return
			//}
			noti := &protocol.Notification{}
			err = json.Unmarshal(msg.Body, noti)
			if err != nil {
				Logger.Errorf("消息格式错误, notification, %s, err: %v", string(msg.Body), err)
				msg.Reject(false)
				return
			}
			client, ok := s.GetComet(noti.ServerId)
			if !ok {
				Logger.Errorf("Comet Server: %d 不存在", noti.ServerId)
				msg.Reject(true)
				continue
			}
			ct, err := json.Marshal(noti.Content)
			if err != nil {
				Logger.Errorf("序列化内容失败, Content, %v, err: %v", noti.Content, err)
				msg.Reject(false)
				continue
			}
			_, err = client.PushMessage(context.Background(), &comet.PushMessageReq{
				Message: &comet.Message{
					Type:      int32(noti.Type),
					Content:   ct,
					UserId:    int32(noti.UserId),
					Timestamp: noti.Timestamp,
				},
			})
			if err != nil {
				Logger.Errorf("推送消息失败, err: %v", err)
				msg.Reject(true)
				return
			}
			msg.Ack(false)
		}
	}
}
