package cache

import (
	"testing"

	"math/rand"

	"github.com/stretchr/testify/assert"
)

func TestLru(t *testing.T) {
	lru := NewLru(3)
	keys := []interface{}{1, 2, 3}
	for _, k := range keys {
		assert.NoError(t, lru.Put(k, k))
	}
	for _, k := range keys {
		v, e := lru.Get(k)
		assert.NoError(t, e)
		assert.Equal(t, k, v)
	}
	// 3 was last accessed so should be in front
	assert.Equal(t, lru.l.Front().Value, 3)
	assert.NoError(t, lru.Put(4, 4))
	// 4 was just inserted so should be in the front
	assert.Equal(t, lru.l.Front().Value, 4)
	// 1 was last accessed and should be evicted
	assert.Equal(t, lru.l.Back().Value, 2)
	v, e := lru.Get(2)
	assert.NoError(t, e)
	assert.Equal(t, 2, v)
	assert.Equal(t, lru.l.Front().Value, 2)

	_, e = lru.Get(55)
	assert.Equal(t, ErrNotFound, e)

	lru.EvictAll()
	assert.Equal(t, 0, lru.l.Len())
	assert.Equal(t, 0, len(lru.c))
}

func BenchmarkLru(b *testing.B) {
	size := 10000
	lru := NewLru(size)
	keys := make([]interface{}, size)
	for i := 0; i < size; i++ {
		keys[i] = rand.Int()
	}
	for _, k := range keys {
		lru.Put(k, k)
	}
	s := int32(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lru.Get(keys[rand.Int31n(s)])
	}
}
