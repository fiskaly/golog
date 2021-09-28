package golog

import (
	"io"
	"testing"
)

func BenchmarkLogger(b *testing.B) {
	l := NewLogger(io.Discard, nil)

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			l.Debug("important info", nil, nil)
		}
	})
}
