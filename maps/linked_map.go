// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package maps

type LinkedMap[K comparable, V any] interface {
	Put(key K, value V)
	Get(key K) (V, bool)
	GetNillable(key K) *V
	Keys() []K
	Values() []V
	Entries() []Entry[K, V]
}

type linkedMap[K comparable, V any] struct {
	delegate map[K]V
	keys     []K
}

func NewLinkedMap[K comparable, V any]() *linkedMap[K, V] {
	return &linkedMap[K, V]{
		delegate: make(map[K]V),
		keys:     make([]K, 0),
	}
}

func (m *linkedMap[K, V]) Put(key K, value V) {
	m.keys = append(m.keys, key)
	m.delegate[key] = value
}

func (m *linkedMap[K, V]) Get(key K) (V, bool) {
	value, ok := m.delegate[key]
	return value, ok
}

func (m *linkedMap[K, V]) GetNillable(key K) *V {
	value, ok := m.delegate[key]
	if ok {
		return &value
	}
	return nil
}

func (m *linkedMap[K, V]) Keys() []K {
	return m.keys
}

func (m *linkedMap[K, V]) Values() []V {
	res := make([]V, len(m.keys))
	for i := range m.keys {
		res[i] = m.delegate[m.keys[i]]
	}
	return res
}

func (m *linkedMap[K, V]) Entries() []Entry[K, V] {
	res := make([]Entry[K, V], len(m.keys))
	for i := range m.keys {
		res[i] = Entry[K, V]{
			Key:   m.keys[i],
			Value: m.delegate[m.keys[i]],
		}
	}
	return res
}
