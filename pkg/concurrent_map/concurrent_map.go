package concurrentmap

import (
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

func (m *ConcurrentMap[K, V]) Get(key K) V {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	return m.data[key]
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
