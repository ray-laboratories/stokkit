package memcache

import (
	"context"
	"fmt"
	"sync"
)

type Key interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type MemCache[K Key, T any] struct {
	cache map[K]T
	mutex sync.RWMutex
}

func NewMemCache[K Key, T any]() *MemCache[K, T] {
	return &MemCache[K, T]{
		cache: make(map[K]T),
	}
}

func (m *MemCache[K, T]) Get(ctx context.Context, id K) (T, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	obj, ok := m.cache[id]
	if !ok {
		var zero T
		return zero, fmt.Errorf("id=%d not found", id)
	}
	return obj, nil
}

func (m *MemCache[K, T]) GetAll(ctx context.Context) ([]T, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	objs := make([]T, 0, len(m.cache))
	for _, obj := range m.cache {
		objs = append(objs, obj)
	}
	return objs, nil
}

func (m *MemCache[K, T]) Update(ctx context.Context, id K, obj T) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.cache[id] = obj
	return nil
}

func (m *MemCache[K, T]) Create(ctx context.Context, obj T) (K, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	var maxKey K
	for k := range m.cache {
		if k > maxKey {
			maxKey = k
		}
	}
	newKey := maxKey + 1
	m.cache[newKey] = obj
	return newKey, nil
}

func (m *MemCache[K, T]) Delete(ctx context.Context, id K) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.cache, id)
	return nil
}
