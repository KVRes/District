package channels

type Meta struct {
	Flags int32
	Buf   int
}

type baseChannel struct {
	meta Meta
}

func NewChannel[T any](meta Meta) IChannel[T] {
	return &ChannelT[T]{
		baseChannel: newBaseChannel(meta),
		ch:          make(chan T, meta.Buf),
	}
}

type IChannel[T any] interface {
	Meta() Meta
	Send(msg T)
	Recv() T
	Len() int
	Cap() int
}

func newBaseChannel(meta Meta) baseChannel {
	return baseChannel{
		meta: meta,
	}
}

func (c baseChannel) Meta() Meta {
	return c.meta
}
