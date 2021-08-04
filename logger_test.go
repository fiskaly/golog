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

	l.Debug("hello world", nil, nil)

	//Output:
	// {"logging.googleapis.com/sourceLocation":{"file":"logger_test.go","line":11},"message":"hello world","severity":"DEBUG"}
}

func ExampleAdvanced() {
	l := NewLogger(os.Stdout, Fields{
		"request_id": "12345",
		"dynamic":    func() string { return "interesting" },
	})

	l.Info("hello world", nil, nil)

	//Output:
	// {"dynamic":"interesting","logging.googleapis.com/sourceLocation":{"file":"logger_test.go","line":23},"message":"hello world","request_id":"12345","severity":"INFO"}
}

func ExampleContext() {
	ctx := WithLogger(context.Background(), os.Stdout, Fields{
		"test": 12345,
	})

	Notice(ctx, "hello world", nil, nil)

	//Output:
	// {"logging.googleapis.com/sourceLocation":{"file":"logger_test.go","line":34},"message":"hello world","severity":"NOTICE","test":12345}
}

func ExampleDefault() {
	// empty context
	Warning(context.Background(), "empty context", nil, nil)

	// it works even with a nil context
	Error(nil, "nil context", nil, nil)

	//Output:
	// {"logging.googleapis.com/sourceLocation":{"file":"logger_test.go","line":42},"message":"empty context","severity":"WARNING"}
	// {"logging.googleapis.com/sourceLocation":{"file":"logger_test.go","line":45},"message":"nil context","severity":"ERROR"}
}
