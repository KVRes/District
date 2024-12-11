package exchange

import "sync"

type Exchange struct {
	m sync.Map
}

func (e *Exchange) Register(key string, buf int) {
	if _, ok := e.m.Load(key); ok {
		return
	}
	e.m.Store(key, NewChannel(buf))
}

func (e *Exchange) Unregister(key string) {
	e.m.Delete(key)
}

func (e *Exchange) GetCh(key string) (*Channel, bool) {
	ch, ok := e.m.Load(key)
	if !ok || ch == nil {
		return nil, false
	}
	return ch.(*Channel), true
}
