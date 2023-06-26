// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package maps

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWhenEmpty(t *testing.T) {
	for i, m := range getTestable[int, string]() {
		actual, isPresent := m.Get(1)

		assert.IsTypef(t, "", actual, msgWithIndex(i))
		assert.Lenf(t, actual, 0, msgWithIndex(i))
		assert.Falsef(t, isPresent, msgWithIndex(i))
	}
}

func TestGet(t *testing.T) {
	for i, m := range getTestable[int, string]() {
		m.Put(1, "one")

		actual, isPresent := m.Get(1)

		assert.Truef(t, isPresent, msgWithIndex(i))
		assert.Equalf(t, "one", actual, msgWithIndex(i))
	}
}

func TestGetNillableWhenEmpty(t *testing.T) {
	for i, m := range getTestable[int, string]() {
		actual := m.GetNillable(1)

		assert.Nilf(t, actual, msgWithIndex(i))
	}
}

func TestGetNillable(t *testing.T) {
	for i, m := range getTestable[int, string]() {
		m.Put(2, "two")

		actual := m.GetNillable(2)

		assert.Equalf(t, ptr("two"), actual, msgWithIndex(i))
	}

}

func TestKeysWhenEmpty(t *testing.T) {
	for i, m := range getTestable[int, string]() {
		actual := m.Keys()

		assert.Emptyf(t, actual, msgWithIndex(i))
	}
}

func TestKeys(t *testing.T) {
	for i, m := range getTestable[int, string]() {
		m.Put(2, "two")
		m.Put(3, "three")
		m.Put(1, "one")

		actual := m.Keys()

		assert.Exactlyf(t, []int{2, 3, 1}, actual, msgWithIndex(i))
	}
}

func TestValuesWhenEmpty(t *testing.T) {
	for i, m := range getTestable[int, string]() {
		actual := m.Values()

		assert.Emptyf(t, actual, msgWithIndex(i))
	}
}

func TestValues(t *testing.T) {
	for i, m := range getTestable[int, string]() {
		m.Put(3, "three")
		m.Put(1, "one")
		m.Put(2, "two")

		actual := m.Values()

		assert.Exactlyf(t, []string{"three", "one", "two"}, actual, msgWithIndex(i))
	}
}

func TestEntriesWhenEmpty(t *testing.T) {
	for i, m := range getTestable[int, string]() {
		actual := m.Entries()

		assert.Emptyf(t, actual, msgWithIndex(i))
	}
}

func TestEntries(t *testing.T) {
	for i, m := range getTestable[int, string]() {
		m.Put(3, "three")
		m.Put(2, "two")
		m.Put(1, "one")

		actual := m.Entries()

		assert.Exactlyf(t, []Entry[int, string]{
			{Key: 3, Value: "three"},
			{Key: 2, Value: "two"},
			{Key: 1, Value: "one"},
		}, actual, msgWithIndex(i))
	}
}

func TestLinkedMapWithCustomComparableTypes(t *testing.T) {
	for i, m := range getTestable[comparableTypeKey, string]() {
		m.Put(comparableTypeKey{someInt: 2, someString: "two"}, "two value")
		m.Put(comparableTypeKey{someInt: 1, someString: "one"}, "one value")

		get, isPresent := m.Get(comparableTypeKey{someInt: 1, someString: "one"})
		assert.Truef(t, isPresent, msgWithIndex(i))
		assert.Equalf(t, "one value", get, msgWithIndex(i))

		getNillable := m.GetNillable(comparableTypeKey{someInt: 2, someString: "two"})
		assert.Equalf(t, ptr("two value"), getNillable, msgWithIndex(i))

		keys := m.Keys()
		assert.Exactlyf(t, []comparableTypeKey{
			{someInt: 2, someString: "two"},
			{someInt: 1, someString: "one"},
		}, keys, msgWithIndex(i))

		values := m.Values()
		assert.Exactlyf(t, []string{"two value", "one value"}, values, msgWithIndex(i))

		entries := m.Entries()
		assert.Exactlyf(t, []Entry[comparableTypeKey, string]{
			{comparableTypeKey{someInt: 2, someString: "two"}, "two value"},
			{comparableTypeKey{someInt: 1, someString: "one"}, "one value"},
		}, entries, msgWithIndex(i))
	}
}

func getTestable[K comparable, V any]() []LinkedMap[K, V] {
	return []LinkedMap[K, V]{
		NewLinkedMap[K, V](),
		NewLinkedMapWithSize[K, V](10),
		NewSyncLinkedMap[K, V](),
		NewSyncLinkedMapWithSize[K, V](10),
	}
}

func msgWithIndex(i int) string {
	return fmt.Sprintf("testable index: %d", i)
}

func ptr[T any](t T) *T {
	return &t
}

type comparableTypeKey struct {
	someInt    int
	someString string
}
