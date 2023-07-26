// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package heaps

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func TestMaxHeapWithRandomInts(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	expected, initial := copySlice[int](rand.Perm(10_000))
	sort.Slice(expected, func(i, j int) bool {
		return expected[i] > expected[j]
	})

	h := NewMaxHeap[int]()
	for i := range initial {
		h.Push(initial[i])
	}

	actual, err := drainHeap[int](h)
	assert.Nil(t, err)
	assert.Exactly(t, expected, actual)
}

func copySlice[K constraints.Ordered](slice []K) ([]K, []K) {
	res := make([]K, len(slice))
	copy(res, slice)
	return slice, res
}

func drainHeap[K constraints.Ordered](h Heap[K]) ([]K, error) {
	res := make([]K, 0)
	for !h.IsEmpty() {
		elem, err := h.Pop()
		if err != nil {
			return make([]K, 0), err
		}
		res = append(res, elem)
	}
	return res, nil
}
