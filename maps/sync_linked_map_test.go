// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package maps

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// other tests are parameterized in linked_map_test.go
func TestSyncLinkedMapConcurrently(t *testing.T) {
	m := NewSyncLinkedMap[int, string]()
	m.Put(1, "one")

	testSize := 100
	resChan := make(chan testTimeAndResult, testSize)
	var wg sync.WaitGroup
	wg.Add(testSize)

	var timeOfChange time.Time
	go func() {
		time.Sleep(2 * time.Second)
		timeOfChange = timed(func() {
			m.Put(1, "two")
		})
	}()

	for i := 0; i < testSize; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(3_000)) * time.Millisecond)

			var res string
			stopTime := timed(func() {
				res = *m.GetNillable(1)
			})
			resChan <- testTimeAndResult{stopTime: stopTime, result: res}
		}()
	}

	wg.Wait()
	close(resChan)

	fmt.Printf("put finished at: %v\n", timeOfChange)
	for res := range resChan {
		fmt.Printf("result: %v, get finished at: %v\n", res.result, res.stopTime)
		msg := fmt.Sprintf("get finished at: %v, put finished at: %v", res.stopTime, timeOfChange)
		if res.stopTime.Before(timeOfChange) {
			assert.Equalf(t, "one", res.result, msg)
		} else {
			assert.Equalf(t, "two", res.result, msg)
		}
	}
}

func timed(f func()) time.Time {
	startTime := time.Now()
	f()
	return startTime.Add(time.Since(startTime))
}

type testTimeAndResult struct {
	stopTime time.Time
	result   string
}
