// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package queues

import "errors"

type FIFOQueue[T any] struct {
	delegate []T
}

func NewFIFOQueue[T any]() *FIFOQueue[T] {
	return &FIFOQueue[T]{}
}

func (q *FIFOQueue[T]) Push(elem T) {
	q.delegate = append(q.delegate, elem)
}

func (q *FIFOQueue[T]) Pop() (T, error) {
	if q.IsEmpty() {
		return *new(T), errors.New("the queue is empty, nothing to pop")
	}

	res := q.delegate[0]
	q.delegate = q.delegate[1:]
	return res, nil
}

func (q *FIFOQueue[T]) Size() int {
	return len(q.delegate)
}

func (q *FIFOQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}
