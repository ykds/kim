package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

type Config struct {
	Urls     []string `json:"urls" yaml:"urls"`
	Endpoint string   `json:"endpoint"`
}

type Client struct {
	cfg         Config
	cli         *clientv3.Client
	manager     endpoints.Manager
	registerMap map[string]context.CancelFunc
}

func NewEtcd(cfg Config) *Client {
	cli, err := clientv3.NewFromURLs(cfg.Urls)
	if err != nil {
		panic(err)
	}
	manager, err := endpoints.NewManager(cli, cfg.Endpoint)
	if err != nil {
		panic(err)
	}
	return &Client{manager: manager, cli: cli, cfg: cfg, registerMap: make(map[string]context.CancelFunc)}
}

func (c *Client) put(key string, addr string, md map[string]interface{}, lease clientv3.Lease) (clientv3.Lease, clientv3.LeaseID, error) {
	if lease == nil {
		lease = clientv3.NewLease(c.cli)
	}
	grant, err := lease.Grant(context.Background(), 30)
	if err != nil {
		return lease, 0, err
	}
	return lease, grant.ID, c.manager.AddEndpoint(context.Background(), c.cfg.Endpoint+key, endpoints.Endpoint{Addr: addr, Metadata: md}, clientv3.WithLease(grant.ID))
}

func (c *Client) Register(key string, addr string, md map[string]interface{}) error {
	lease, leaseID, err := c.put(key, addr, md, nil)
	if err != nil {
		return err
	}
	alive, err := lease.KeepAlive(context.Background(), leaseID)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	c.registerMap[key] = cancel
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case _ = <-alive:
				alive, err = lease.KeepAlive(context.Background(), leaseID)
				if err != nil {
					return
				}
			}
		}
	}()
	return nil
}

func (c *Client) UnRegister(key string) {
	cancel, ok := c.registerMap[key]
	if ok {
		cancel()
		delete(c.registerMap, key)
	}
	_, _ = c.cli.Delete(context.Background(), key)
}

func (c *Client) Watch() (endpoints.WatchChannel, error) {
	return c.manager.NewWatchChannel(context.Background())
}
