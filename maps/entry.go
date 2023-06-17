// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package maps

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}
