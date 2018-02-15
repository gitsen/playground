package cache

import (
	"github.com/pkg/errors"
)

type Key interface{}
type Value interface{}

var ErrNotFound = errors.New("key not found")
var ErrCacheFull = errors.New("cache has reached its maximum capacity")

type Cache interface {
	Get(k Key) (Value, error)
	Put(k Key, v Value) error
	Evict(k Key)
	EvictAll()
}
