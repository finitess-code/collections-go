// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWhenEmpty(t *testing.T) {
	m := NewLinkedMap[int, string]()

	actual, isPresent := m.Get(1)

	assert.IsType(t, "", actual)
	assert.Len(t, actual, 0)
	assert.False(t, isPresent)
}

func TestGet(t *testing.T) {
	m := NewLinkedMap[int, string]()
	m.Put(1, "one")

	actual, isPresent := m.Get(1)
	assert.True(t, isPresent)
	assert.Equal(t, "one", actual)
}

func TestGetNillableWhenEmpty(t *testing.T) {
	m := NewLinkedMap[int, string]()

	actual := m.GetNillable(1)

	assert.Nil(t, actual)
}

func TestGetNillable(t *testing.T) {
	m := NewLinkedMap[int, string]()
	m.Put(2, "two")

	actual := m.GetNillable(2)
	assert.Equal(t, ptr("two"), actual)
}

func TestKeysWhenEmpty(t *testing.T) {
	m := NewLinkedMap[int, string]()

	actual := m.Keys()

	assert.Empty(t, actual)
}

func TestKeys(t *testing.T) {
	m := NewLinkedMap[int, string]()
	m.Put(2, "two")
	m.Put(3, "three")
	m.Put(1, "one")

	actual := m.Keys()

	assert.Exactly(t, []int{2, 3, 1}, actual)
}

func TestValuesWhenEmpty(t *testing.T) {
	m := NewLinkedMap[int, string]()

	actual := m.Values()

	assert.Empty(t, actual)
}

func TestValues(t *testing.T) {
	m := NewLinkedMap[int, string]()
	m.Put(3, "three")
	m.Put(1, "one")
	m.Put(2, "two")

	actual := m.Values()

	assert.Exactly(t, []string{"three", "one", "two"}, actual)
}

func TestEntriesWhenEmpty(t *testing.T) {
	m := NewLinkedMap[int, string]()

	actual := m.Entries()

	assert.Empty(t, actual)
}

func TestEntries(t *testing.T) {
	m := NewLinkedMap[int, string]()
	m.Put(3, "three")
	m.Put(2, "two")
	m.Put(1, "one")

	actual := m.Entries()

	assert.Exactly(t, []Entry[int, string]{
		{Key: 3, Value: "three"},
		{Key: 2, Value: "two"},
		{Key: 1, Value: "one"},
	}, actual)
}

func TestLinkedMapWithCustomComparableTypes(t *testing.T) {
	m := NewLinkedMap[comparableTypeKey, string]()
	m.Put(comparableTypeKey{someInt: 2, someString: "two"}, "two value")
	m.Put(comparableTypeKey{someInt: 1, someString: "one"}, "one value")

	get, isPresent := m.Get(comparableTypeKey{someInt: 1, someString: "one"})
	assert.True(t, isPresent)
	assert.Equal(t, "one value", get)

	getNillable := m.GetNillable(comparableTypeKey{someInt: 2, someString: "two"})
	assert.Equal(t, ptr("two value"), getNillable)

	keys := m.Keys()
	assert.Exactly(t, []comparableTypeKey{
		{someInt: 2, someString: "two"},
		{someInt: 1, someString: "one"},
	}, keys)

	values := m.Values()
	assert.Exactly(t, []string{"two value", "one value"}, values)

	entries := m.Entries()
	assert.Exactly(t, []Entry[comparableTypeKey, string]{
		{comparableTypeKey{someInt: 2, someString: "two"}, "two value"},
		{comparableTypeKey{someInt: 1, someString: "one"}, "one value"},
	}, entries)
}

func ptr[T any](t T) *T {
	return &t
}

type comparableTypeKey struct {
	someInt    int
	someString string
}
