package exchange

type ChannelT[T any] struct {
	ch chan T
	// listener []net.Listener
	// l_lck    sync.RWMutex
}

/*

func (c *ChannelT[T]) AddListener(listener net.Listener) {
	c.l_lck.Lock()
	defer c.l_lck.Unlock()
	c.listener = append(c.listener, listener)
}

func (c *ChannelT[T]) RemoveListener(listener net.Listener) {
	c.l_lck.Lock()
	defer c.l_lck.Unlock()
	for i, l := range c.listener {
		if l == listener {
			c.listener = append(c.listener[:i], c.listener[i+1:]...)
			break
		}
	}
}

func (c *ChannelT[T]) AllListener() []net.Listener {
	c.l_lck.RLock()
	defer c.l_lck.RUnlock()
	return c.listener
}
*/

func NewChannel[T any](buf int) *ChannelT[T] {
	return &ChannelT[T]{
		ch: make(chan T, buf),
	}
}

func (c *ChannelT[T]) Send(msg T) {
	c.ch <- msg
}

func (c *ChannelT[T]) Recv() T {
	return <-c.ch
}

func (c *ChannelT[T]) Len() int {
	return len(c.ch)
}

func (c *ChannelT[T]) Cap() int {
	return cap(c.ch)
}
