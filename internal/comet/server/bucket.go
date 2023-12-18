package server

import "sync"

type Bucket struct {
	m     sync.RWMutex
	users map[uint]*Channel
}

func NewBucket() *Bucket {
	return &Bucket{
		users: make(map[uint]*Channel),
	}
}

func (b *Bucket) Put(userId uint, ch *Channel) {
	b.m.Lock()
	if c, ok := b.users[userId]; ok {
		c.Close()
	}
	b.users[userId] = ch
	b.m.Unlock()
}

func (b *Bucket) Get(userId uint) (ch *Channel, ok bool) {
	b.m.RLock()
	defer b.m.RUnlock()
	ch, ok = b.users[userId]
	return
}
