// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package queues

type Queue[T any] interface {
	Push(T)
	Pop() (T, error)
	Size() int
	IsEmpty() bool
}
