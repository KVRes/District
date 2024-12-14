package exchange

type ChMeta struct {
	Flags int32
	Buf   int
}

type ChannelT[T any] struct {
	meta ChMeta
	ch   chan T
}

func NewChannel[T any](meta ChMeta) *ChannelT[T] {
	return &ChannelT[T]{
		meta: meta,
		ch:   make(chan T, meta.Buf),
	}
}

func (c *ChannelT[T]) Meta() ChMeta {
	return c.meta
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
