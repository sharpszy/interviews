package common

import (
	"sync"
)

type RWSyncMap[K comparable, V any] struct {
	data map[K]V
	lock sync.RWMutex
	size int
}

func NewRWSyncMap[K comparable, V any](size ...uint) *(RWSyncMap[K, V]) {
	var n uint = 0
	if len(size) > 0 {
		n = size[0]
	}
	return &RWSyncMap[K, V]{
		data: make(map[K]V, n),
		lock: sync.RWMutex{},
		size: 0,
	}
}

// Put a k-v entry into map and return old value
func (m *RWSyncMap[K, V]) Put(k K, v V) V {
	m.lock.Lock()
	defer m.lock.Unlock()
	r := m.data[k]
	m.data[k] = v
	m.size += 1
	return r
}

// Get value from map by k
func (m *RWSyncMap[K, V]) Get(k K) V {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.data[k]
}

// Delete k entry from map and return k's value
func (m *RWSyncMap[K, V]) Delete(k K) V {
	m.lock.Lock()
	defer m.lock.Unlock()

	var d V
	d, ok := m.data[k]
	if ok {
		delete(m.data, k)
		m.size -= 1
	}
	return d
}

func (m *RWSyncMap[K, V]) Range(f func(k K, v V) bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for kk, vv := range m.data {
		if !f(kk, vv) {
			return
		}
	}
}

func (m *RWSyncMap[K, V]) Len() int {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.size
}
