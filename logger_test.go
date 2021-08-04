package golog

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimple(t *testing.T) {
	var logBuffer bytes.Buffer
	logger := NewLogger(&logBuffer, nil)

	type test struct {
		Execute         func(*Logger)
		ExpectedLevel   LogLevel
		ExpectedMessage string
	}

	tests := []test{
		{
			Execute:         func(l *Logger) { l.Warning("Warning", nil, nil) },
			ExpectedLevel:   LevelWarning,
			ExpectedMessage: "Warning",
		},
	}

	for _, test := range tests {
		test.Execute(logger)
		logBytes, err := io.ReadAll(&logBuffer)
		require.NoError(t, err)
		require.NotNil(t, logBytes)
	}
}

func ExampleSimple() {
	l := NewLogger(os.Stdout, nil)

	l.Warning("hello world", nil, nil)

	//Output:
	// {"logging.googleapis.com/sourceLocation":{"file":"/home/christoph/Dev/go/github.com/fiskaly/golog/logger_test.go","line":11},"message":"hello world","severity":"WARNING"}
}

func ExampleAdvanced() {
	l := NewLogger(os.Stdout, Fields{
		"request_id": "12345",
		"dynamic":    func() string { return "interesting" },
	})

	l.Critical("hello world", nil, nil)

	//Output:
	// {"dynamic":"interesting","logging.googleapis.com/sourceLocation":{"file":"/home/christoph/Dev/go/github.com/fiskaly/golog/logger_test.go","line":23},"message":"hello world","request_id":"12345","severity":"CRITICAL"}
}

func ExampleContext() {
	ctx := WithLogger(context.Background(), os.Stdout, nil)

	Debug(ctx, "hello world", nil, nil)

	//Output:
	// {"logging.googleapis.com/sourceLocation":{"file":"/home/christoph/Dev/go/github.com/fiskaly/golog/logger_test.go","line":32},"message":"hello world","severity":"DEBUG"}
}
