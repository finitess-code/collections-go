// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package heaps

import "golang.org/x/exp/constraints"

type Heap[K constraints.Ordered] interface {
	Push(K)
	Pop() (K, error)
	IsEmpty() bool
}
