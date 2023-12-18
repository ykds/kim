package job

import (
	"context"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"kim/internal/pkg/etcd"
	"kim/proto/comet"
)

func Watch(manager *etcd.Client, srv *Server) {
	watch, err := manager.Watch()
	if err != nil {
		panic(err)
	}
	for {
		select {
		case update := <-watch:
			for _, item := range update {
				if item.Op == endpoints.Add {
					md := item.Endpoint.Metadata.(map[string]interface{})
					client, err := newCometClient(item.Endpoint.Addr)
					if err != nil {
						Logger.Errorf("连接Comet Grpc失败， error: %v", err)
						continue
					}
					srv.AddComet(int32(md["server_id"].(float64)), client)
				}
				if item.Op == endpoints.Delete {
					md := item.Endpoint.Metadata.(map[string]interface{})
					srv.DelComet(int32(md["server_id"].(float64)))
				}
			}
		}
	}

}

func newCometClient(addr string) (comet.CometClient, error) {
	conn, err := grpc.DialContext(context.Background(), addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return comet.NewCometClient(conn), nil
}