package server

import (
	"github.com/gorilla/websocket"
	"kim/proto/comet"
	"sync/atomic"
	"time"
)

type Channel struct {
	conn       *websocket.Conn
	ch         chan *comet.Message
	closed     chan struct{}
	isClosed   atomic.Bool
	hbticker   *time.Ticker
	hbInterval int
}

func NewChannel(conn *websocket.Conn, hbInterval int) *Channel {
	c := &Channel{
		conn:       conn,
		ch:         make(chan *comet.Message),
		closed:     make(chan struct{}),
		hbInterval: hbInterval,
		hbticker:   time.NewTicker(time.Duration(hbInterval) * time.Second),
	}
	c.isClosed.Store(false)
	return c
}

func (c *Channel) Signal() <-chan *comet.Message {
	return c.ch
}

func (c *Channel) Put(msg *comet.Message) {
	c.ch <- msg
}

func (c *Channel) WriteMessage(body []byte) error {
	return c.conn.WriteMessage(websocket.TextMessage, body)
}

func (c *Channel) Close() error {
	if !c.isClosed.Load() {
		close(c.closed)
		c.isClosed.Store(true)
		return c.conn.Close()
	}
	return nil
}

func (c *Channel) IsClosed() bool {
	return c.isClosed.Load()
}

func (c *Channel) Done() <-chan struct{} {
	return c.closed
}

func (c *Channel) HeartBeat() <-chan time.Time {
	return c.hbticker.C
}
