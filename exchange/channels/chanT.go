package channels

type ChannelT[T any] struct {
	baseChannel
	ch chan T
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
