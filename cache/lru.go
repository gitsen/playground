package cache

import (
	"container/list"
	"sync"
)

type Lru struct {
	c        map[Key]*list.Element
	l        list.List
	mu       sync.Mutex
	capacity int
}

func NewLru(capacity int) *Lru {
	return &Lru{
		c:        make(map[Key]*list.Element),
		l:        list.List{},
		capacity: capacity,
	}
}

func (l *Lru) Get(k Key) (Value, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	var e *list.Element
	e, ok := l.c[k]
	if !ok {
		return nil, ErrNotFound
	} else {
		l.l.MoveToFront(e)
	}
	return e.Value, nil
}

func (l *Lru) Put(k Key, v Value) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if e, ok := l.c[k]; ok {
		e.Value = v
		l.l.MoveToFront(e)
	} else {
		l.c[k] = l.l.PushFront(v)
	}
	if l.l.Len() == l.capacity+1 {
		l.l.Remove(l.l.Back())
	}
	return nil
}

func (l *Lru) Evict(k Key) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if e, ok := l.c[k]; ok {
		l.l.Remove(e)
		delete(l.c, k)
	}
}

func (l *Lru) EvictAll() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.l.Init()
	l.c = make(map[Key]*list.Element)
}
