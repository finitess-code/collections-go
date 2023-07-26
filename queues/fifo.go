// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package queues

import "errors"

type FIFOQueue[K comparable] struct {
	delegate []K
}

func NewFIFOQueue[K comparable]() *FIFOQueue[K] {
	return &FIFOQueue[K]{}
}

func (q *FIFOQueue[K]) Push(elem K) {
	q.delegate = append(q.delegate, elem)
}

func (q *FIFOQueue[K]) Pop() (K, error) {
	if q.IsEmpty() {
		return *new(K), errors.New("the queue is empty, nothing to pop")
	}

	res := q.delegate[0]
	q.delegate = q.delegate[1:]
	return res, nil
}

func (q *FIFOQueue[K]) Size() int {
	return len(q.delegate)
}

func (q *FIFOQueue[K]) IsEmpty() bool {
	return q.Size() == 0
}
