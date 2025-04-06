package concurrentmap

import (
	"maps"
	"sync"
)

type ConcurrentMap[K comparable, V any] struct {
	data  map[K]V
	mutex sync.RWMutex
}

func newConcurrentMap[K comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{
		data: make(map[K]V),
	}
}

func NewConcurrentMap[K comparable, V any]() IConcurrentMap[K, V] {
	return newConcurrentMap[K, V]()
}

func (m *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	value, exists := m.data[key]
	return value, exists
}

func (m *ConcurrentMap[K, V]) Set(key K, value V) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.data[key] = value
}

func (m *ConcurrentMap[K, V]) Delete(key K) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.data, key)
}

func (m *ConcurrentMap[K, V]) Len() int {
	return len(m.data)
}

func (m *ConcurrentMap[K, V]) Iter() func(func(K, V) bool) {
	return m.Next
}

func (m *ConcurrentMap[K, V]) Next(yield func(K, V) bool) {
	snap := make(map[K]V, len(m.data))
	m.snapshot(snap)
	for key, value := range snap {
		if !yield(key, value) {
			return
		}
	}
}

func (m *ConcurrentMap[K, V]) snapshot(dst map[K]V) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	maps.Copy(dst, m.data)
}
