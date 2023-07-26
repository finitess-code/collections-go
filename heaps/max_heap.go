// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package heaps

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type MaxHeap[K constraints.Ordered] struct {
	delegate []K
}

func NewMaxHeap[K constraints.Ordered]() *MaxHeap[K] {
	return &MaxHeap[K]{}
}

func (h *MaxHeap[K]) Push(elem K) {
	h.delegate = append(h.delegate, elem)

	// heapify up
	idx := len(h.delegate) - 1
	for h.delegate[parentIndex(idx)] < h.delegate[idx] {
		h.swap(parentIndex(idx), idx)
		idx = parentIndex(idx)
	}
}

func (h *MaxHeap[K]) Pop() (K, error) {
	if len(h.delegate) == 0 {
		return *new(K), errors.New("the heap is empty, nothing to pop")
	}

	res := h.delegate[0]
	h.delegate[0] = h.delegate[len(h.delegate)-1]
	h.delegate = h.delegate[:len(h.delegate)-1]

	// heapify down
	parentIdx := 0
	leftIdx, rightIdx := leftChildIndex(parentIdx), rightChildIndex(parentIdx)
	for leftIdx < len(h.delegate) {
		var candidateIdx int
		if rightIdx >= len(h.delegate) || h.delegate[leftIdx] > h.delegate[rightIdx] {
			candidateIdx = leftIdx
		} else {
			candidateIdx = rightIdx
		}

		if h.delegate[parentIdx] >= h.delegate[candidateIdx] {
			return res, nil
		}

		h.swap(parentIdx, candidateIdx)
		parentIdx = candidateIdx
		leftIdx, rightIdx = leftChildIndex(parentIdx), rightChildIndex(parentIdx)
	}

	return res, nil
}

func (h *MaxHeap[K]) IsEmpty() bool {
	return len(h.delegate) == 0
}

func (h *MaxHeap[K]) swap(i, j int) {
	h.delegate[i], h.delegate[j] = h.delegate[j], h.delegate[i]
}

func parentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func leftChildIndex(parentIndex int) int {
	return parentIndex*2 + 1
}

func rightChildIndex(parentIndex int) int {
	return parentIndex*2 + 2
}
