package utils

import (
	"runtime"

	"github.com/sourcegraph/conc/pool"
)

const parallelProcessingThreshold = 100

// ParallelForEach iterates over a slice and executes the action function for each element in parallel.
// It automatically switches to sequential processing if the number of items is small (below 100).
// The concurrency is limited to runtime.NumCPU().
func ParallelForEach[T any](items []T, action func(index int, item T)) {
	if len(items) < parallelProcessingThreshold {
		for i, item := range items {
			action(i, item)
		}
		return
	}

	p := pool.New().WithMaxGoroutines(runtime.NumCPU())
	for i, item := range items {
		i, item := i, item
		p.Go(func() {
			action(i, item)
		})
	}
	p.Wait()
}
