package exchange

import "sync"

type ExchangeT[T any] struct {
	m sync.Map
}

func NewExchange[T any]() *ExchangeT[T] {
	return &ExchangeT[T]{
		m: sync.Map{},
	}
}

func (e *ExchangeT[T]) Register(key string, buf int) bool {
	if _, ok := e.m.Load(key); ok {
		return false
	}
	e.m.Store(key, NewChannel[T](buf))
	return true
}

func (e *ExchangeT[T]) Unregister(key string) {
	e.m.Delete(key)
}

func (e *ExchangeT[T]) GetCh(key string) (*ChannelT[T], bool) {
	ch, ok := e.m.Load(key)
	if !ok || ch == nil {
		return nil, false
	}
	return ch.(*ChannelT[T]), true
}
