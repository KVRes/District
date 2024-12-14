package exchange

import (
	"sync"

	"github.com/KVRes/District/exchange/channels"
)

type ExchangeT[T any] struct {
	m sync.Map
}

func NewExchange[T any]() *ExchangeT[T] {
	return &ExchangeT[T]{
		m: sync.Map{},
	}
}

func (e *ExchangeT[T]) Register(key string, meta channels.Meta) bool {
	if _, ok := e.m.Load(key); ok {
		return false
	}
	e.m.Store(key, channels.NewChannel[T](meta))
	return true
}

func (e *ExchangeT[T]) Unregister(key string) {
	e.m.Delete(key)
}

func (e *ExchangeT[T]) GetCh(key string) (channels.IChannel[T], bool) {
	ch, ok := e.m.Load(key)
	if !ok || ch == nil {
		return nil, false
	}
	return ch.(channels.IChannel[T]), true
}
