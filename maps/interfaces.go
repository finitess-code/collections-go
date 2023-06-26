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
