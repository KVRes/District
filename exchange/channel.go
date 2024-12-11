package exchange

import (
	"net"
	"sync"
)

type Channel struct {
	ch       chan []byte
	listener []net.Listener
	l_lck    sync.RWMutex
}

func (c *Channel) AddListener(listener net.Listener) {
	c.l_lck.Lock()
	defer c.l_lck.Unlock()
	c.listener = append(c.listener, listener)
}

func (c *Channel) RemoveListener(listener net.Listener) {
	c.l_lck.Lock()
	defer c.l_lck.Unlock()
	for i, l := range c.listener {
		if l == listener {
			c.listener = append(c.listener[:i], c.listener[i+1:]...)
			break
		}
	}
}

func (c *Channel) AllListener() []net.Listener {
	c.l_lck.RLock()
	defer c.l_lck.RUnlock()
	return c.listener
}

func NewChannel(buf int) *Channel {
	return &Channel{
		ch: make(chan []byte, buf),
	}
}

func (c *Channel) Send(msg []byte) {
	c.ch <- msg
}

func (c *Channel) Recv() []byte {
	return <-c.ch
}
