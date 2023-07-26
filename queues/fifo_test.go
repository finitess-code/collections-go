// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFIFOQueueWithInts(t *testing.T) {
	// assert empty
	queue := NewFIFOQueue[int]()
	assertQueue[int](t, queue, 0, true)

	// assert after push
	queue.Push(2)
	queue.Push(1)
	assertQueue[int](t, queue, 2, false)

	// assert after pop
	elem, err := queue.Pop()
	assert.Equal(t, 2, elem)
	assert.Nil(t, err)
	assertQueue[int](t, queue, 1, false)

	elem, err = queue.Pop()
	assert.Equal(t, 1, elem)
	assert.Nil(t, err)
	assertQueue[int](t, queue, 0, true)

	elem, err = queue.Pop()
	assert.Equal(t, "the queue is empty, nothing to pop", err.Error())
	assertQueue[int](t, queue, 0, true)
}

func TestFIFOQueueWithStrings(t *testing.T) {
	// assert empty
	queue := NewFIFOQueue[string]()
	assertQueue[string](t, queue, 0, true)

	// assert after push
	queue.Push("b")
	queue.Push("a")
	assertQueue[string](t, queue, 2, false)

	// assert after pop
	elem, err := queue.Pop()
	assert.Equal(t, "b", elem)
	assert.Nil(t, err)
	assertQueue[string](t, queue, 1, false)

	elem, err = queue.Pop()
	assert.Equal(t, "a", elem)
	assert.Nil(t, err)
	assertQueue[string](t, queue, 0, true)

	elem, err = queue.Pop()
	assert.Equal(t, "the queue is empty, nothing to pop", err.Error())
	assertQueue[string](t, queue, 0, true)
}

func assertQueue[T any](t *testing.T, queue Queue[T], size int, isEmpty bool) {
	assert.Equal(t, size, queue.Size())
	assert.Equal(t, isEmpty, queue.IsEmpty())
}
