//go:build stats
// +build stats

package golog

import (
	"sync"
	"time"

	"github.com/fiskaly/gostats"
)

var (
	stats = &struct {
		sync.Mutex
		counter uint64

		lastCounter uint64
		accumulator *gostats.Accumulator
	}{}
)

func init() {
	stats.accumulator = gostats.NewAccumulator(60, func() float64 {
		stats.Lock()
		c := stats.counter
		stats.Unlock()

		v := c - stats.lastCounter
		stats.lastCounter = c

		return float64(v)
	})
	stats.accumulator.Run(time.Second)
}

func countLog() {
	stats.Lock()
	stats.counter++
	stats.Unlock()
}

// Stats returns the number of written log lines for the last second and an average over the last minute.
func Stats() (second, minute float64) {
	a := stats.accumulator

	return a.Last(), a.Average()
}
