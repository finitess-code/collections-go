// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package maps

import "sync"

type syncLinkedMap[K comparable, V any] struct {
	lock     sync.RWMutex
	delegate map[K]V
	keys     []K
}

func NewSyncLinkedMap[K comparable, V any]() *syncLinkedMap[K, V] {
	return &syncLinkedMap[K, V]{
		delegate: make(map[K]V),
		keys:     make([]K, 0),
	}
}

func NewSyncLinkedMapWithSize[K comparable, V any](size int) *syncLinkedMap[K, V] {
	return &syncLinkedMap[K, V]{
		delegate: make(map[K]V, size),
		keys:     make([]K, 0, size),
	}
}

func (m *syncLinkedMap[K, V]) Put(key K, value V) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.keys = append(m.keys, key)
	m.delegate[key] = value
}

func (m *syncLinkedMap[K, V]) Get(key K) (V, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	value, ok := m.delegate[key]
	return value, ok
}

func (m *syncLinkedMap[K, V]) GetNillable(key K) *V {
	m.lock.RLock()
	defer m.lock.RUnlock()

	value, ok := m.delegate[key]
	if ok {
		return &value
	}
	return nil
}

func (m *syncLinkedMap[K, V]) Keys() []K {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.keys
}

func (m *syncLinkedMap[K, V]) Values() []V {
	m.lock.RLock()
	defer m.lock.RUnlock()

	res := make([]V, len(m.keys))
	for i := range m.keys {
		res[i] = m.delegate[m.keys[i]]
	}
	return res
}

func (m *syncLinkedMap[K, V]) Entries() []Entry[K, V] {
	m.lock.RLock()
	defer m.lock.RUnlock()

	res := make([]Entry[K, V], len(m.keys))
	for i := range m.keys {
		res[i] = Entry[K, V]{
			Key:   m.keys[i],
			Value: m.delegate[m.keys[i]],
		}
	}
	return res
}
