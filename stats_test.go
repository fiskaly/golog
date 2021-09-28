//go:build stats
// +build stats

package golog

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestStats(t *testing.T) {
	start := time.Now()

	ticker := time.Tick(10 * time.Millisecond)
	for range ticker {
		Debug(context.Background(), "hi", nil, nil)

		if time.Since(start) > 10*time.Second {
			break
		}
	}

	second, minute := Stats()
	require.Greater(t, second, 90.0)
	require.Greater(t, minute, 15.0)
}
